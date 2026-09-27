# BudPayments Golang SDK



The Golang SDK for the BudPayments API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.InitiatePaymentBudLicense(nil)` — each with the same small set of operations (`List`, `Load`, `Create`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `go-cli`, `go-mcp`, `js`, `lua`, `php`, `py`, `ts` — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/bud-payments-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/bud-payments-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/bud-payments-sdk/go=../bud-payments-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    "os"
    sdk "github.com/voxgig-sdk/bud-payments-sdk/go"
)

func main() {
    client := sdk.NewBudPaymentsSDK(map[string]any{
        "apikey": os.Getenv("BUD_PAYMENTS_APIKEY"),
    })

    // List initiatePaymentBudLicense records — the value is the array of records itself.
    initiatePaymentBudLicenses, err := client.InitiatePaymentBudLicense(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range initiatePaymentBudLicenses.([]any) {
        fmt.Println(item)
    }

    // Create a initiatePaymentBudLicense.
    created, err := client.InitiatePaymentBudLicense(nil).Create(map[string]any{"bud_pay_url": "example_bud_pay_url", "country_code": "example_country_code", "display_name": "example_display_name", "icon": "example_icon", "implementation_type": "example_implementation_type", "maintenance_status": "example_maintenance_status", "maintenance_window": "example_maintenance_window", "payment_details": map[string]any{}, "payment_id": "example_payment_id", "provider": "example_provider", "redirect_url": "example_redirect_url", "required_actions": "example_required_actions", "scheduled_payment_details": map[string]any{}, "services": []any{}, "standing_order_details": map[string]any{}, "supported_currencies": []any{}}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(created)
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
initiatepaymentbudlicenses, err := client.InitiatePaymentBudLicense(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = initiatepaymentbudlicenses
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

initiatePaymentBudLicense, err := client.InitiatePaymentBudLicense(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(initiatePaymentBudLicense) // the returned mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewBudPaymentsSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
BUD_PAYMENTS_TEST_LIVE=TRUE
BUD_PAYMENTS_APIKEY=<your-key>
```

Then run:

```bash
cd go && go test ./test/...
```


## Reference

### NewBudPaymentsSDK

```go
func NewBudPaymentsSDK(options map[string]any) *BudPaymentsSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *BudPaymentsSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### BudPaymentsSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `InitiatePaymentBudLicense` | `(data map[string]any) BudPaymentsEntity` | Create an InitiatePaymentBudLicense entity instance. |
| `InitiatePaymentClientLicense` | `(data map[string]any) BudPaymentsEntity` | Create an InitiatePaymentClientLicense entity instance. |
| `ManagePayment` | `(data map[string]any) BudPaymentsEntity` | Create a ManagePayment entity instance. |

### Entity interface (BudPaymentsEntity)

All entities implement the `BudPaymentsEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    initiatePaymentBudLicense, err := client.InitiatePaymentBudLicense(nil).List(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // initiatePaymentBudLicense is the returned record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### InitiatePaymentBudLicense

| Field | Description |
| --- | --- |
| `"bud_pay_url"` | A Bud Pay URL, allowing Customers to complete (authenticate) a relevant payment resource with the relevant payment service provider. |
| `"country_code"` | Provider country code following the ISO 3166 alpha-3 code format |
| `"display_name"` | The name of the payment provider to be used for display purposes (e.g. |
| `"icon"` | A URL for the icon of the corresponding provider |
| `"implementation_type"` | The implementation standard adopted by the provider |
| `"maintenance_status"` | The status of the maintenance window associated with the provider |
| `"maintenance_window"` |  |
| `"payment_details"` | Payment details required to make a Single Payment |
| `"payment_id"` | Unique identifier for the payment |
| `"provider"` | The name (identifier) of the payment provider |
| `"provider_types"` | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `"redirect_url"` | URL where the user will be redirected to once they have completed the TPP Connection flow. |
| `"required_actions"` | An array of enum fields to indicate if any additional actions are required in order to complete a given payment flow for the given provider. |
| `"scheduled_payment_details"` | Details of the scheduled payment |
| `"services"` | A payment service supported by a payment provider |
| `"standing_order_details"` | Details of the standing order |
| `"supported_currencies"` | The currencies supported by the payment provider |

Operations: Create, List.

API path: `/v1/payments/scheduled/bud-pay-url`

#### InitiatePaymentClientLicense

| Field | Description |
| --- | --- |
| `"authorisation_url"` | The authorisation URL for a relevant payment provider. |
| `"code"` | Code parameter returned in redirect parameters |
| `"payment_details"` | Payment details required to make a Single Payment |
| `"payment_id"` | Payment Identifier |
| `"payment_type"` | The type of payment iniated through Bud's Payments Service |
| `"provider"` | The name (identifier) of the payment provider |
| `"provider_redirect_url"` | This is the url that your Customer will be redirected to once they have authorised with the relevant provider. |
| `"provider_types"` | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `"redirect_url"` | This is the url that your Customer will be redirected to when using the Bud user interface. |
| `"required_action"` | If present, this fild will indicate the next required action (if any) to be taken by the client in order to complete the payment flow | Value | Description | |:-------|:-------------| | confirm_scheduled_payment | Indicates that the client… |
| `"result"` |  |
| `"scheduled_payment_details"` | Details of the scheduled payment |
| `"standing_order_details"` | Details of the standing order |
| `"state"` | The task ID |

Operations: Create.

API path: `/v1/payments/authorisation-codes`

#### ManagePayment

| Field | Description |
| --- | --- |
| `"amount"` | The monetary amount. |
| `"client_id"` | Bud's Client/App Identifier |
| `"created_at"` | The created at timestamp for the payment |
| `"data"` | Data structure for the Scheduled Payment object |
| `"errors"` | A list of any errors that may be associated with the payment |
| `"first_payment_date"` | Date-time (ISO 8601) the first payment should be made. |
| `"frequency"` | The payment frequency for a Standing Order |
| `"known_charges"` | A list detailing information surrounding any known charges that may be associated with the payment |
| `"last_payment_date"` | Date-time (ISO 8601) after which payments should stop. |
| `"metadata"` |  |
| `"method"` | The payment method |
| `"operation_id"` | A unique identifier/reference associated with a given endpoint/operation |
| `"organisation_id"` | Bud's Organisation Identifier |
| `"payment_id"` | Payment Identifier |
| `"recipient"` | Details of the recipient of the payment |
| `"recurring_amount"` | The monetary amount. |
| `"reference"` | Reference to use for this payment |
| `"requested_execution_date"` | Date-time (ISO 8601) the Scheduled Payment should be made. |
| `"required_action"` | This field will indicate that there is an action required to be taken by the client in order to complete the payment. |
| `"sender"` | Details of the payment sender |
| `"state"` | Details of the payment state and history |
| `"supplementary_status"` | A list of supplementary states associated with the payment as described by the specific provider |

Operations: Create, List, Load.

API path: `/v1/payments/scheduled/{payment_id}/confirm`



## Entities


### InitiatePaymentBudLicense

Create an instance: `initiatePaymentBudLicense := client.InitiatePaymentBudLicense(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bud_pay_url` | `string` | A Bud Pay URL, allowing Customers to complete (authenticate) a relevant payment resource with the relevant payment service provider. |
| `country_code` | `string` | Provider country code following the ISO 3166 alpha-3 code format |
| `display_name` | `string` | The name of the payment provider to be used for display purposes (e.g. |
| `icon` | `string` | A URL for the icon of the corresponding provider |
| `implementation_type` | `string` | The implementation standard adopted by the provider |
| `maintenance_status` | `string` | The status of the maintenance window associated with the provider |
| `maintenance_window` | `any` |  |
| `payment_details` | `map[string]any` | Payment details required to make a Single Payment |
| `payment_id` | `string` | Unique identifier for the payment |
| `provider` | `string` | The name (identifier) of the payment provider |
| `provider_types` | `[]any` | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `string` | URL where the user will be redirected to once they have completed the TPP Connection flow. |
| `required_actions` | `any` | An array of enum fields to indicate if any additional actions are required in order to complete a given payment flow for the given provider. |
| `scheduled_payment_details` | `map[string]any` | Details of the scheduled payment |
| `services` | `[]any` | A payment service supported by a payment provider |
| `standing_order_details` | `map[string]any` | Details of the standing order |
| `supported_currencies` | `[]any` | The currencies supported by the payment provider |

#### Example: List

```go
initiatePaymentBudLicenses, err := client.InitiatePaymentBudLicense(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(initiatePaymentBudLicenses) // the array of records
```

#### Example: Create

```go
result, err := client.InitiatePaymentBudLicense(nil).Create(map[string]any{
    "bud_pay_url": "example_bud_pay_url",
    "country_code": "example_country_code",
    "display_name": "example_display_name",
    "icon": "example_icon",
    "implementation_type": "example_implementation_type",
    "maintenance_status": "example_maintenance_status",
    "maintenance_window": "example_maintenance_window",
    "payment_details": map[string]any{},
    "payment_id": "example_payment_id",
    "provider": "example_provider",
    "redirect_url": "example_redirect_url",
    "required_actions": "example_required_actions",
    "scheduled_payment_details": map[string]any{},
    "services": []any{},
    "standing_order_details": map[string]any{},
    "supported_currencies": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### InitiatePaymentClientLicense

Create an instance: `initiatePaymentClientLicense := client.InitiatePaymentClientLicense(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_url` | `string` | The authorisation URL for a relevant payment provider. |
| `code` | `string` | Code parameter returned in redirect parameters |
| `payment_details` | `map[string]any` | Payment details required to make a Single Payment |
| `payment_id` | `string` | Payment Identifier |
| `payment_type` | `string` | The type of payment iniated through Bud's Payments Service |
| `provider` | `string` | The name (identifier) of the payment provider |
| `provider_redirect_url` | `string` | This is the url that your Customer will be redirected to once they have authorised with the relevant provider. |
| `provider_types` | `[]any` | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `string` | This is the url that your Customer will be redirected to when using the Bud user interface. |
| `required_action` | `string` | If present, this fild will indicate the next required action (if any) to be taken by the client in order to complete the payment flow | Value | Description | |:-------|:-------------| | confirm_scheduled_payment | Indicates that the client… |
| `result` | `string` |  |
| `scheduled_payment_details` | `map[string]any` | Details of the scheduled payment |
| `standing_order_details` | `map[string]any` | Details of the standing order |
| `state` | `string` | The task ID |

#### Example: Create

```go
result, err := client.InitiatePaymentClientLicense(nil).Create(map[string]any{
    "authorisation_url": "example_authorisation_url",
    "payment_details": map[string]any{},
    "provider": "example_provider",
    "scheduled_payment_details": map[string]any{},
    "standing_order_details": map[string]any{},
    "state": "example_state",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ManagePayment

Create an instance: `managePayment := client.ManagePayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `map[string]any` | The monetary amount. |
| `client_id` | `string` | Bud's Client/App Identifier |
| `created_at` | `string` | The created at timestamp for the payment |
| `data` | `map[string]any` | Data structure for the Scheduled Payment object |
| `errors` | `[]any` | A list of any errors that may be associated with the payment |
| `first_payment_date` | `string` | Date-time (ISO 8601) the first payment should be made. |
| `frequency` | `string` | The payment frequency for a Standing Order |
| `known_charges` | `map[string]any` | A list detailing information surrounding any known charges that may be associated with the payment |
| `last_payment_date` | `string` | Date-time (ISO 8601) after which payments should stop. |
| `metadata` | `map[string]any` |  |
| `method` | `string` | The payment method |
| `operation_id` | `string` | A unique identifier/reference associated with a given endpoint/operation |
| `organisation_id` | `string` | Bud's Organisation Identifier |
| `payment_id` | `string` | Payment Identifier |
| `recipient` | `map[string]any` | Details of the recipient of the payment |
| `recurring_amount` | `map[string]any` | The monetary amount. |
| `reference` | `string` | Reference to use for this payment |
| `requested_execution_date` | `string` | Date-time (ISO 8601) the Scheduled Payment should be made. |
| `required_action` | `string` | This field will indicate that there is an action required to be taken by the client in order to complete the payment. |
| `sender` | `map[string]any` | Details of the payment sender |
| `state` | `map[string]any` | Details of the payment state and history |
| `supplementary_status` | `map[string]any` | A list of supplementary states associated with the payment as described by the specific provider |

#### Example: Load

```go
managePayment, err := client.ManagePayment(nil).Load(map[string]any{"payment_id": "payment_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(managePayment) // the loaded record
```

#### Example: List

```go
managePayments, err := client.ManagePayment(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(managePayments) // the array of records
```

#### Example: Create

```go
result, err := client.ManagePayment(nil).Create(map[string]any{
    "scheduled_id": "example_scheduled_id",
    "amount": map[string]any{},
    "client_id": "example_client_id",
    "created_at": "example_created_at",
    "data": map[string]any{},
    "errors": []any{},
    "first_payment_date": "example_first_payment_date",
    "frequency": "example_frequency",
    "known_charges": map[string]any{},
    "method": "example_method",
    "operation_id": "example_operation_id",
    "organisation_id": "example_organisation_id",
    "payment_id": "example_payment_id",
    "recipient": map[string]any{},
    "recurring_amount": map[string]any{},
    "reference": "example_reference",
    "requested_execution_date": "example_requested_execution_date",
    "required_action": "example_required_action",
    "sender": map[string]any{},
    "state": map[string]any{},
    "supplementary_status": map[string]any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

## Features

This SDK ships 8 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`debug`](#debug) | Debug capture |
| [`idempotency`](#idempotency) | Idempotency |
| [`metrics`](#metrics) | Metrics |
| [`paging`](#paging) | Paging |
| [`ratelimit`](#ratelimit) | Rate limiting |
| [`retry`](#retry) | Retry |
| [`test`](#test) | Test transport |
| [`timeout`](#timeout) | Timeout |

> **Order matters for `ratelimit`, `retry`, `timeout`.** These wrap the
> transport, so each one wraps whatever is already installed: the order you
> activate them in IS the nesting order. Activating them as an ordered list
> rather than a map is what fixes that order.

### debug

Debug capture.

| Option | Default |
|---|---|
| `active` | `false` |
| `max` | `100` |
| `redact` | `['authorization', 'cookie', 'set-cookie', 'api-key', 'apikey', 'x-api-key', 'idempotency-key']` |

Set `feature.debug.active` to enable it, then override any of the options above.

### idempotency

Idempotency.

| Option | Default |
|---|---|
| `active` | `false` |
| `header` | `'Idempotency-Key'` |
| `methods` | `['POST', 'PUT', 'PATCH', 'DELETE']` |
| `ops` | `['create', 'update', 'remove']` |

Set `feature.idempotency.active` to enable it, then override any of the options above.

### metrics

Metrics.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.metrics.active` to enable it, then override any of the options above.

### paging

Paging.

| Option | Default |
|---|---|
| `active` | `false` |
| `afterVar` | `'after'` |
| `cursorParam` | `'cursor'` |
| `firstVar` | `'first'` |
| `limitParam` | `'limit'` |
| `pageParam` | `'page'` |
| `startPage` | `1` |

Set `feature.paging.active` to enable it, then override any of the options above.

### ratelimit

Rate limiting.

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

Set `feature.ratelimit.active` to enable it, then override any of the options above.

`ratelimit` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### retry

Retry.

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

Set `feature.retry.active` to enable it, then override any of the options above.

`retry` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### test

Test transport.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.

### timeout

Timeout.

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

Set `feature.timeout.active` to enable it, then override any of the options above.

`timeout` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **DebugFeature**: Debug capture
- **IdempotencyFeature**: Idempotency
- **MetricsFeature**: Metrics
- **PagingFeature**: Paging
- **RatelimitFeature**: Rate limiting
- **RetryFeature**: Retry
- **TestFeature**: Test transport
- **TimeoutFeature**: Timeout

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/bud-payments-sdk/go/
├── bud-payments.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/bud-payments-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `List`, the entity
stores the returned data and match criteria internally.

```go
initiatepaymentbudlicense := client.InitiatePaymentBudLicense(nil)
initiatepaymentbudlicense.List(nil, nil)

// initiatepaymentbudlicense.Data() now returns the initiatepaymentbudlicense data from the last list
// initiatepaymentbudlicense.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
