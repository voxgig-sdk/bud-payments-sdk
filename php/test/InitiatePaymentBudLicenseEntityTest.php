<?php
declare(strict_types=1);

// InitiatePaymentBudLicense entity test

require_once __DIR__ . '/../budpayments_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class InitiatePaymentBudLicenseEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = BudPaymentsSDK::test(null, null);
        $ent = $testsdk->InitiatePaymentBudLicense(null);
        $this->assertNotNull($ent);
    }

    // Feature #4: the entity stream(action, ...) method runs the op pipeline
    // and yields result items. With the streaming feature active it yields the
    // feature's incremental output; otherwise it falls back to the materialised
    // list so stream always yields.
    public function test_stream(): void
    {
        $seed = [
            "entity" => [
                "initiate_payment_bud_license" => [
                    "s1" => ["id" => "s1"],
                    "s2" => ["id" => "s2"],
                    "s3" => ["id" => "s3"],
                ],
            ],
        ];

        // Fallback: streaming inactive -> yields the materialised list items.
        $base = BudPaymentsSDK::test($seed, null);
        $seen = iterator_to_array($base->InitiatePaymentBudLicense(null)->stream("list", null, null), false);
        $this->assertCount(3, $seen);

        // Inbound: streaming active -> yields each item from the feature.
        $cfg = BudPaymentsConfig::shared_config();
        if (isset($cfg["feature"]) && is_array($cfg["feature"]) && isset($cfg["feature"]["streaming"])) {
            $sdk = BudPaymentsSDK::test($seed, ["feature" => ["streaming" => ["active" => true]]]);
            $got = [];
            foreach ($sdk->InitiatePaymentBudLicense(null)->stream("list", null, null) as $item) {
                if (is_array($item) && array_is_list($item)) {
                    foreach ($item as $sub) {
                        $got[] = $sub;
                    }
                } else {
                    $got[] = $item;
                }
            }
            $this->assertCount(3, $got);
        }
    }

    public function test_basic_flow(): void
    {
        $setup = initiate_payment_bud_license_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "list"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "initiate_payment_bud_license." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set BUD_PAYMENTS_TEST_INITIATE_PAYMENT_BUD_LICENSE_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $initiate_payment_bud_license_ref01_ent = $client->InitiatePaymentBudLicense(null);
        $initiate_payment_bud_license_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.initiate_payment_bud_license"), "initiate_payment_bud_license_ref01"));

        $initiate_payment_bud_license_ref01_data_result = $initiate_payment_bud_license_ref01_ent->create($initiate_payment_bud_license_ref01_data, null);
        $initiate_payment_bud_license_ref01_data = Helpers::to_map(is_object($initiate_payment_bud_license_ref01_data_result) && method_exists($initiate_payment_bud_license_ref01_data_result, 'data_get') ? $initiate_payment_bud_license_ref01_data_result->data_get() : $initiate_payment_bud_license_ref01_data_result);
        $this->assertNotNull($initiate_payment_bud_license_ref01_data);

        // LIST
        $initiate_payment_bud_license_ref01_match = [];

        $initiate_payment_bud_license_ref01_list_result = $initiate_payment_bud_license_ref01_ent->list($initiate_payment_bud_license_ref01_match, null);
        $this->assertIsArray($initiate_payment_bud_license_ref01_list_result);

    }
}

function initiate_payment_bud_license_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/initiate_payment_bud_license/InitiatePaymentBudLicenseTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = BudPaymentsSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["initiate_payment_bud_license01", "initiate_payment_bud_license02", "initiate_payment_bud_license03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("BUD_PAYMENTS_TEST_INITIATE_PAYMENT_BUD_LICENSE_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "BUD_PAYMENTS_TEST_INITIATE_PAYMENT_BUD_LICENSE_ENTID" => $idmap,
        "BUD_PAYMENTS_TEST_LIVE" => "FALSE",
        "BUD_PAYMENTS_TEST_EXPLAIN" => "FALSE",
        "BUD_PAYMENTS_APIKEY" => "",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["BUD_PAYMENTS_TEST_INITIATE_PAYMENT_BUD_LICENSE_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["BUD_PAYMENTS_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            Runner::live_client_options(),
            [
                "apikey" => $env["BUD_PAYMENTS_APIKEY"],
            ],
            // ismap, not a plain "?? []" default: an empty PHP array is a
            // LIST, and a non-map later entry REPLACES the accumulated map in
            // merge - so the no-extras call discarded live_client_options()
            // and the apikey/server map above it.
            Vs::ismap($extra) ? $extra : new \stdClass(),
        ]);
        // "?? []" because merge legitimately answers with a stdClass when every
        // contributing entry is an EMPTY map - an SDK with no apikey and no
        // server variables generates an empty middle entry, so that is the
        // common case, not the edge one. to_map returns null for a non-array by
        // design, and the constructor takes a non-nullable array, so without the
        // fallback every such SDK died on "must be of type array, null given"
        // the moment live mode was switched on. Offline mode never reaches this
        // branch, which is why the offline suite stayed green.
        $client = new BudPaymentsSDK(Helpers::to_map($merged_opts) ?? []);
    }

    $live = $env["BUD_PAYMENTS_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["BUD_PAYMENTS_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
