package sdktest

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	sdk "github.com/voxgig-sdk/bud-payments-sdk/go"
	"github.com/voxgig-sdk/bud-payments-sdk/go/core"
)

func TestInitiatePaymentBudLicenseDirect(t *testing.T) {
	t.Run("direct-list-initiate_payment_bud_license", func(t *testing.T) {
		setup := initiate_payment_bud_licenseDirectSetup([]any{
			map[string]any{"id": "direct01"},
			map[string]any{"id": "direct02"},
		})
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		if _shouldSkip, _reason := isControlSkipped("direct", "direct-list-initiate_payment_bud_license", _mode); _shouldSkip {
			if _reason == "" {
				_reason = "skipped via sdk-test-control.json"
			}
			t.Skip(_reason)
			return
		}
		client := setup.client


		result, err := client.Direct(map[string]any{
			"path":   "v1/payments/providers",
			"method": "GET",
			"params": map[string]any{},
		})
		if setup.live {
			// Live-mode leniency is a model decision
			// (main.kit.test.live.strict): synthetic IDs 4xx constantly
			// against an arbitrary public API, so the default SKIPS here.
			// A project that owns its test server sets strict and FAILS.
			if err != nil {
				t.Fatalf("list call failed (likely synthetic IDs against live API): %v", err)
			}
			if result["ok"] != true {
				t.Fatalf("list call not ok (likely synthetic IDs against live API): %v", result)
			}
			status := core.ToInt(result["status"])
			if status < 200 || status >= 300 {
				t.Fatalf("expected 2xx status, got %v", result["status"])
			}
		} else {
			if err != nil {
				t.Fatalf("direct failed: %v", err)
			}
			if result["ok"] != true {
				t.Fatalf("expected ok to be true, got %v", result["ok"])
			}
			if core.ToInt(result["status"]) != 200 {
				t.Fatalf("expected status 200, got %v", result["status"])
			}
		}

		if !setup.live {
			if dataList, ok := result["data"].([]any); ok {
				if len(dataList) != 2 {
					t.Fatalf("expected 2 items, got %d", len(dataList))
				}
			} else {
				t.Fatalf("expected data to be an array, got %T", result["data"])
			}

			if len(*setup.calls) != 1 {
				t.Fatalf("expected 1 call, got %d", len(*setup.calls))
			}
		}
	})

}

type initiate_payment_bud_licenseDirectSetupResult struct {
	client *sdk.BudPaymentsSDK
	calls  *[]map[string]any
	live   bool
	idmap  map[string]any
}

func initiate_payment_bud_licenseDirectSetup(mockres any) *initiate_payment_bud_licenseDirectSetupResult {
	loadEnvLocal()

	calls := &[]map[string]any{}

	env := envOverride(map[string]any{
		"BUD_PAYMENTS_TEST_INITIATE_PAYMENT_BUD_LICENSE_ENTID": map[string]any{},
		"BUD_PAYMENTS_TEST_LIVE":    "FALSE",
		"BUD_PAYMENTS_APIKEY":       "",
	})

	live := env["BUD_PAYMENTS_TEST_LIVE"] == "TRUE"

	if live {
		// sdk-test-control.json's test.client.options seeds the live
		// client; the generated fields below overwrite anything they name.
		mergedOpts := map[string]any{}
		for k, v := range liveClientOptions() {
			mergedOpts[k] = v
		}
		for k, v := range map[string]any{
			"apikey": env["BUD_PAYMENTS_APIKEY"],
		} {
			mergedOpts[k] = v
		}
		client := sdk.NewBudPaymentsSDK(mergedOpts)

		idmap := map[string]any{}
		if entidRaw, ok := env["BUD_PAYMENTS_TEST_INITIATE_PAYMENT_BUD_LICENSE_ENTID"]; ok {
			if entidStr, ok := entidRaw.(string); ok && strings.HasPrefix(entidStr, "{") {
				json.Unmarshal([]byte(entidStr), &idmap)
			} else if entidMap, ok := entidRaw.(map[string]any); ok {
				idmap = entidMap
			}
		}

		return &initiate_payment_bud_licenseDirectSetupResult{client: client, calls: calls, live: true, idmap: idmap}
	}

	mockFetch := func(url string, init map[string]any) (map[string]any, error) {
		*calls = append(*calls, map[string]any{"url": url, "init": init})
		return map[string]any{
			"status":     200,
			"statusText": "OK",
			"headers":    map[string]any{},
			"json": (func() any)(func() any {
				if mockres != nil {
					return mockres
				}
				return map[string]any{"id": "direct01"}
			}),
		}, nil
	}

	client := sdk.NewBudPaymentsSDK(map[string]any{
		"base": "http://localhost:8080",
		"system": map[string]any{
			"fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
		},
	})

	return &initiate_payment_bud_licenseDirectSetupResult{client: client, calls: calls, live: false, idmap: map[string]any{}}
}

var _ = os.Getenv
var _ = json.Unmarshal
