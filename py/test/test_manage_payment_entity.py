# ManagePayment entity test

import json
import os
import time

import pytest

from budpayments_sdk.utility.voxgig_struct import voxgig_struct as vs
from budpayments_sdk import BudPaymentsSDK
from budpayments_sdk.core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestManagePaymentEntity:

    def test_should_create_instance(self):
        testsdk = BudPaymentsSDK.test(None, None)
        ent = testsdk.ManagePayment(None)
        assert ent is not None

    def test_should_stream(self):
        # Feature #4: the entity stream(action, ...) method runs the op
        # pipeline and yields result items. With the streaming feature active
        # it yields the feature's incremental output; otherwise it falls back
        # to the materialised list so stream always yields.
        seed = {
            "entity": {
                "manage_payment": {
                    "s1": {"id": "s1"},
                    "s2": {"id": "s2"},
                    "s3": {"id": "s3"},
                }
            }
        }

        # Fallback: streaming inactive -> yields the materialised list items.
        base = BudPaymentsSDK.test(seed, None)
        seen = list(base.ManagePayment(None).stream("list", None, None))
        assert len(seen) == 3

        # Inbound: streaming active -> yields each item from the feature.
        from budpayments_sdk.config import shared_config
        cfg = shared_config()
        if isinstance(cfg.get("feature"), dict) and "streaming" in cfg["feature"]:
            sdk = BudPaymentsSDK.test(
                seed, {"feature": {"streaming": {"active": True}}})
            got = []
            for item in sdk.ManagePayment(None).stream("list", None, None):
                if isinstance(item, list):
                    got.extend(item)
                else:
                    got.append(item)
            assert len(got) == 3

    def test_should_run_basic_flow(self):
        setup = _manage_payment_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create", "list", "load"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "manage_payment." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set BUD_PAYMENTS_TEST_MANAGE_PAYMENT_ENTID JSON to run live")
        client = setup["client"]

        # CREATE
        manage_payment_ref01_ent = client.ManagePayment(None)
        manage_payment_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.manage_payment"), "manage_payment_ref01"))
        manage_payment_ref01_data["standing_order_id"] = setup["idmap"]["standing_order01"]

        manage_payment_ref01_data = helpers.to_map(runner.entity_data(manage_payment_ref01_ent.create(manage_payment_ref01_data, None)))
        assert manage_payment_ref01_data is not None

        # LIST
        manage_payment_ref01_match = {}

        manage_payment_ref01_list_result = manage_payment_ref01_ent.list(manage_payment_ref01_match, None)
        assert isinstance(manage_payment_ref01_list_result, list)

        # LOAD
        manage_payment_ref01_match_dt0 = {}
        manage_payment_ref01_data_dt0_loaded = manage_payment_ref01_ent.load(manage_payment_ref01_match_dt0, None)
        assert manage_payment_ref01_data_dt0_loaded is not None



def _manage_payment_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/manage_payment/ManagePaymentTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = BudPaymentsSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["manage_payment01", "manage_payment02", "manage_payment03", "standing_order01"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Detect ENTID env override before envOverride consumes it. When live
    # mode is on without a real override, the basic test runs against synthetic
    # IDs from the fixture and 4xx's. We surface this so the test can skip.
    _entid_env_raw = os.environ.get(
        "BUD_PAYMENTS_TEST_MANAGE_PAYMENT_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "BUD_PAYMENTS_TEST_MANAGE_PAYMENT_ENTID": idmap,
        "BUD_PAYMENTS_TEST_LIVE": "FALSE",
        "BUD_PAYMENTS_TEST_EXPLAIN": "FALSE",
        "BUD_PAYMENTS_APIKEY": "",
    })

    idmap_resolved = helpers.to_map(
        env.get("BUD_PAYMENTS_TEST_MANAGE_PAYMENT_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

    if env.get("BUD_PAYMENTS_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            # FIRST, so the generated fields below win: sdk-test-control.json's
            # test.client.options adds to the live client, it does not
            # redirect it.
            runner.live_client_options(),
            {
                "apikey": env.get("BUD_PAYMENTS_APIKEY"),
            },
            extra or {},
        ])
        client = BudPaymentsSDK(helpers.to_map(merged_opts))

    _live = env.get("BUD_PAYMENTS_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("BUD_PAYMENTS_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
