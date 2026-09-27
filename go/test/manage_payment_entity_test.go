package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/bud-payments-sdk/go"
	"github.com/voxgig-sdk/bud-payments-sdk/go/core"

	vs "github.com/voxgig-sdk/bud-payments-sdk/go/utility/struct"
)

func TestManagePaymentEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.ManagePayment(nil)
		if ent == nil {
			t.Fatal("expected non-nil ManagePaymentEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"manage_payment": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.ManagePayment(nil).Stream("list", nil, nil) {
			seen = append(seen, item)
		}
		if len(seen) != 3 {
			t.Fatalf("expected 3 streamed items, got %d", len(seen))
		}

		// Inbound: streaming active -> yields each item from the feature iterator.
		hasStreaming := false
		if fm, ok := core.SharedConfig()["feature"].(map[string]any); ok {
			_, hasStreaming = fm["streaming"]
		}
		if hasStreaming {
			streamSdk := sdk.TestSDK(seed, map[string]any{
				"feature": map[string]any{"streaming": map[string]any{"active": true}},
			})
			var got []any
			for item := range streamSdk.ManagePayment(nil).Stream("list", nil, nil) {
				if sub, ok := item.([]any); ok {
					got = append(got, sub...)
				} else {
					got = append(got, item)
				}
			}
			if len(got) != 3 {
				t.Fatalf("expected 3 items via streaming feature, got %d", len(got))
			}
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := manage_paymentBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "manage_payment." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set BUD_PAYMENTS_TEST_MANAGE_PAYMENT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		managePaymentRef01Ent := client.ManagePayment(nil)
		managePaymentRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "manage_payment"}), "manage_payment_ref01"))
		managePaymentRef01Data["standing_order_id"] = setup.idmap["standing_order01"]

		managePaymentRef01DataResult, err := managePaymentRef01Ent.Create(managePaymentRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		managePaymentRef01Data = core.ToMapAny(entityData(managePaymentRef01DataResult))
		if managePaymentRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

		// LIST
		managePaymentRef01Match := map[string]any{}

		managePaymentRef01ListResult, err := managePaymentRef01Ent.List(managePaymentRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, managePaymentRef01ListOk := managePaymentRef01ListResult.([]any)
		if !managePaymentRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", managePaymentRef01ListResult)
		}

		// LOAD
		managePaymentRef01MatchDt0 := map[string]any{}
		managePaymentRef01DataDt0Loaded, err := managePaymentRef01Ent.Load(managePaymentRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if managePaymentRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}

	})
}

func manage_paymentBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "manage_payment", "ManagePaymentTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read manage_payment test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse manage_payment test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"manage_payment01", "manage_payment02", "manage_payment03", "standing_order01"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("BUD_PAYMENTS_TEST_MANAGE_PAYMENT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"BUD_PAYMENTS_TEST_MANAGE_PAYMENT_ENTID": idmap,
		"BUD_PAYMENTS_TEST_LIVE":      "FALSE",
		"BUD_PAYMENTS_TEST_EXPLAIN":   "FALSE",
		"BUD_PAYMENTS_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["BUD_PAYMENTS_TEST_MANAGE_PAYMENT_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["BUD_PAYMENTS_TEST_LIVE"] == "TRUE" {
		// An empty map, not a nil one: Merge returns nil when its last entry
		// is nil, and BasicSetup is normally called with no extras - so a
		// bare nil silently discarded the apikey and server values below.
		extraOpts := extra
		if extraOpts == nil {
			extraOpts = map[string]any{}
		}

		mergedOpts := vs.Merge([]any{
			// liveClientOptions() FIRST, so the generated fields below win:
			// sdk-test-control.json's test.client.options adds to the live
			// client, it does not redirect it.
			liveClientOptions(),
			map[string]any{
				"apikey": env["BUD_PAYMENTS_APIKEY"],
			},
			extraOpts,
		})
		client = sdk.NewBudPaymentsSDK(core.ToMapAny(mergedOpts))
	}

	live := env["BUD_PAYMENTS_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["BUD_PAYMENTS_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
