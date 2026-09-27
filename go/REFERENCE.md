# BudPayments Golang SDK Reference

Complete API reference for the BudPayments Golang SDK.


## BudPaymentsSDK

### Constructor

```go
func NewBudPaymentsSDK(options map[string]any) *BudPaymentsSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *BudPaymentsSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *BudPaymentsSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `InitiatePaymentBudLicense(data map[string]any) BudPaymentsEntity`

Create a new `InitiatePaymentBudLicense` entity instance. Pass `nil` for no initial data.

#### `InitiatePaymentClientLicense(data map[string]any) BudPaymentsEntity`

Create a new `InitiatePaymentClientLicense` entity instance. Pass `nil` for no initial data.

#### `ManagePayment(data map[string]any) BudPaymentsEntity`

Create a new `ManagePayment` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## InitiatePaymentBudLicenseEntity

```go
initiatePaymentBudLicense := client.InitiatePaymentBudLicense(nil)
fmt.Println(initiatePaymentBudLicense.GetName()) // "initiate_payment_bud_license"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bud_pay_url` | `string` | Yes | A Bud Pay URL, allowing Customers to complete (authenticate) a relevant payment resource with the relevant payment service provider. |
| `country_code` | `string` | Yes | Provider country code following the ISO 3166 alpha-3 code format |
| `display_name` | `string` | Yes | The name of the payment provider to be used for display purposes (e.g. |
| `icon` | `string` | Yes | A URL for the icon of the corresponding provider |
| `implementation_type` | `string` | Yes | The implementation standard adopted by the provider |
| `maintenance_status` | `string` | Yes | The status of the maintenance window associated with the provider |
| `maintenance_window` | `any` | Yes |  |
| `payment_details` | `map[string]any` | Yes | Payment details required to make a Single Payment |
| `payment_id` | `string` | Yes | Unique identifier for the payment |
| `provider` | `string` | Yes | The name (identifier) of the payment provider |
| `provider_types` | `[]any` | No | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `string` | Yes | URL where the user will be redirected to once they have completed the TPP Connection flow. |
| `required_actions` | `any` | Yes | An array of enum fields to indicate if any additional actions are required in order to complete a given payment flow for the given provider. |
| `scheduled_payment_details` | `map[string]any` | Yes | Details of the scheduled payment |
| `services` | `[]any` | Yes | A payment service supported by a payment provider |
| `standing_order_details` | `map[string]any` | Yes | Details of the standing order |
| `supported_currencies` | `[]any` | Yes | The currencies supported by the payment provider |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.InitiatePaymentBudLicense(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

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

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `InitiatePaymentBudLicenseEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## InitiatePaymentClientLicenseEntity

```go
initiatePaymentClientLicense := client.InitiatePaymentClientLicense(nil)
fmt.Println(initiatePaymentClientLicense.GetName()) // "initiate_payment_client_license"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_url` | `string` | Yes | The authorisation URL for a relevant payment provider. |
| `code` | `string` | No | Code parameter returned in redirect parameters |
| `payment_details` | `map[string]any` | Yes | Payment details required to make a Single Payment |
| `payment_id` | `string` | No | Payment Identifier |
| `payment_type` | `string` | No | The type of payment iniated through Bud's Payments Service |
| `provider` | `string` | Yes | The name (identifier) of the payment provider |
| `provider_redirect_url` | `string` | No | This is the url that your Customer will be redirected to once they have authorised with the relevant provider. |
| `provider_types` | `[]any` | No | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `string` | No | This is the url that your Customer will be redirected to when using the Bud user interface. |
| `required_action` | `string` | No | If present, this fild will indicate the next required action (if any) to be taken by the client in order to complete the payment flow | Value | Description | |:-------|:-------------| | confirm_scheduled_payment | Indicates that the client… |
| `result` | `string` | No |  |
| `scheduled_payment_details` | `map[string]any` | Yes | Details of the scheduled payment |
| `standing_order_details` | `map[string]any` | Yes | Details of the standing order |
| `state` | `string` | Yes | The task ID |

### Field Usage by Operation

| Field | create |
| --- | --- |
| `authorisation_url` | - |
| `code` | - |
| `payment_details` | - |
| `payment_id` | Yes |
| `payment_type` | - |
| `provider` | - |
| `provider_redirect_url` | - |
| `provider_types` | - |
| `redirect_url` | - |
| `required_action` | - |
| `result` | - |
| `scheduled_payment_details` | - |
| `standing_order_details` | - |
| `state` | - |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

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

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `InitiatePaymentClientLicenseEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ManagePaymentEntity

```go
managePayment := client.ManagePayment(nil)
fmt.Println(managePayment.GetName()) // "manage_payment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `map[string]any` | Yes | The monetary amount. |
| `client_id` | `string` | Yes | Bud's Client/App Identifier |
| `created_at` | `string` | Yes | The created at timestamp for the payment |
| `data` | `map[string]any` | Yes | Data structure for the Scheduled Payment object |
| `errors` | `[]any` | Yes | A list of any errors that may be associated with the payment |
| `first_payment_date` | `string` | Yes | Date-time (ISO 8601) the first payment should be made. |
| `frequency` | `string` | Yes | The payment frequency for a Standing Order |
| `known_charges` | `map[string]any` | Yes | A list detailing information surrounding any known charges that may be associated with the payment |
| `last_payment_date` | `string` | No | Date-time (ISO 8601) after which payments should stop. |
| `metadata` | `map[string]any` | No |  |
| `method` | `string` | Yes | The payment method |
| `operation_id` | `string` | Yes | A unique identifier/reference associated with a given endpoint/operation |
| `organisation_id` | `string` | Yes | Bud's Organisation Identifier |
| `payment_id` | `string` | Yes | Payment Identifier |
| `recipient` | `map[string]any` | Yes | Details of the recipient of the payment |
| `recurring_amount` | `map[string]any` | Yes | The monetary amount. |
| `reference` | `string` | Yes | Reference to use for this payment |
| `requested_execution_date` | `string` | Yes | Date-time (ISO 8601) the Scheduled Payment should be made. |
| `required_action` | `string` | Yes | This field will indicate that there is an action required to be taken by the client in order to complete the payment. |
| `sender` | `map[string]any` | Yes | Details of the payment sender |
| `state` | `map[string]any` | Yes | Details of the payment state and history |
| `supplementary_status` | `map[string]any` | Yes | A list of supplementary states associated with the payment as described by the specific provider |

### Field Usage by Operation

| Field | load | list | create |
| --- | --- | --- | --- |
| `amount` | - | - | - |
| `client_id` | - | - | - |
| `created_at` | - | - | - |
| `data` | - | - | - |
| `errors` | - | - | - |
| `first_payment_date` | - | - | - |
| `frequency` | - | - | - |
| `known_charges` | - | - | - |
| `last_payment_date` | - | - | - |
| `metadata` | - | - | - |
| `method` | - | - | - |
| `operation_id` | - | - | - |
| `organisation_id` | - | - | - |
| `payment_id` | - | - | - |
| `recipient` | - | - | - |
| `recurring_amount` | - | - | - |
| `reference` | - | - | - |
| `requested_execution_date` | - | - | - |
| `required_action` | - | Yes | - |
| `sender` | - | - | - |
| `state` | - | - | - |
| `supplementary_status` | - | - | - |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.ManagePayment(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.ManagePayment(nil).Load(map[string]any{"payment_id": "payment_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

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

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ManagePaymentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `debug` | 0.0.1 | Debug capture |
| `idempotency` | 0.0.1 | Idempotency |
| `metrics` | 0.0.1 | Metrics |
| `paging` | 0.0.1 | Paging |
| `ratelimit` | 0.0.1 | Rate limiting |
| `retry` | 0.0.1 | Retry |
| `test` | 0.0.1 | Test transport |
| `timeout` | 0.0.1 | Timeout |


Features are activated via the `feature` option:

```go
client := sdk.NewBudPaymentsSDK(map[string]any{
    "feature": map[string]any{
        "debug": map[string]any{"active": true},
        "idempotency": map[string]any{"active": true},
        "metrics": map[string]any{"active": true},
        "paging": map[string]any{"active": true},
        "ratelimit": map[string]any{"active": true},
        "retry": map[string]any{"active": true},
        "test": map[string]any{"active": true},
        "timeout": map[string]any{"active": true},
    },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### Ordering

`ratelimit`, `retry`, `timeout` wrap the transport. Each
wraps whatever is already installed, so **activation order is nesting order**:
a feature activated later sits OUTSIDE one activated earlier, and sees the call
first.

That decides behaviour, not just sequence: a feature that short-circuits the
call, such as a cache serving a hit, stops every feature nested inside it from
ever seeing that call.

`debug`, `idempotency`, `metrics`, `paging`, `test` attach to pipeline hooks
rather than the transport, so their order does not affect what they observe.

#### `debug`

Debug capture.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `max` | `100` |
| `redact` | `['authorization', 'cookie', 'set-cookie', 'api-key', 'apikey', 'x-api-key', 'idempotency-key']` |

| Option | Type |
|---|---|
| `now` | function |
| `onEntry` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.debug.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `idempotency`

Idempotency.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `header` | `'Idempotency-Key'` |
| `methods` | `['POST', 'PUT', 'PATCH', 'DELETE']` |
| `ops` | `['create', 'update', 'remove']` |

| Option | Type |
|---|---|
| `keygen` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.idempotency.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `metrics`

Metrics.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `now` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.metrics.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `paging`

Paging.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `afterVar` | `'after'` |
| `cursorParam` | `'cursor'` |
| `firstVar` | `'first'` |
| `limitParam` | `'limit'` |
| `pageParam` | `'page'` |
| `startPage` | `1` |

| Option | Type |
|---|---|
| `limit` | number |
| `ops` | list |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.paging.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `ratelimit`

Rate limiting.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

| Option | Type |
|---|---|
| `now` | function |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.ratelimit.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `retry`

Retry.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

| Option | Type |
|---|---|
| `jitter` | boolean |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.retry.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `test`

Test transport.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

#### `timeout`

Timeout.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

| Option | Type |
|---|---|
| `clearTimer` | function |
| `setTimer` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.timeout.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

