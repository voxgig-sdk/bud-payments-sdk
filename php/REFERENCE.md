# BudPayments PHP SDK Reference

Complete API reference for the BudPayments PHP SDK.


## BudPaymentsSDK

### Constructor

```php
require_once __DIR__ . '/budpayments_sdk.php';

$client = new BudPaymentsSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `BudPaymentsSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = BudPaymentsSDK::test();
```


### Instance Methods

#### `InitiatePaymentBudLicense($data = null)`

Create a new `InitiatePaymentBudLicenseEntity` instance. Pass `null` for no initial data.

#### `InitiatePaymentClientLicense($data = null)`

Create a new `InitiatePaymentClientLicenseEntity` instance. Pass `null` for no initial data.

#### `ManagePayment($data = null)`

Create a new `ManagePaymentEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): BudPaymentsUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## InitiatePaymentBudLicenseEntity

```php
$initiate_payment_bud_license = $client->InitiatePaymentBudLicense();
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
| `maintenance_window` | `mixed` | Yes |  |
| `payment_details` | `array` | Yes | Payment details required to make a Single Payment |
| `payment_id` | `string` | Yes | Unique identifier for the payment |
| `provider` | `string` | Yes | The name (identifier) of the payment provider |
| `provider_types` | `array` | No | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `string` | Yes | URL where the user will be redirected to once they have completed the TPP Connection flow. |
| `required_actions` | `mixed` | Yes | An array of enum fields to indicate if any additional actions are required in order to complete a given payment flow for the given provider. |
| `scheduled_payment_details` | `array` | Yes | Details of the scheduled payment |
| `services` | `array` | Yes | A payment service supported by a payment provider |
| `standing_order_details` | `array` | Yes | Details of the standing order |
| `supported_currencies` | `array` | Yes | The currencies supported by the payment provider |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->InitiatePaymentBudLicense()->create([
  "bud_pay_url" => null, // string
  "country_code" => null, // string
  "display_name" => null, // string
  "icon" => null, // string
  "implementation_type" => null, // string
  "maintenance_status" => null, // string
  "maintenance_window" => null, // mixed
  "payment_details" => null, // array
  "payment_id" => null, // string
  "provider" => null, // string
  "redirect_url" => null, // string
  "required_actions" => null, // mixed
  "scheduled_payment_details" => null, // array
  "services" => null, // array
  "standing_order_details" => null, // array
  "supported_currencies" => null, // array
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->InitiatePaymentBudLicense()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): InitiatePaymentBudLicenseEntity`

Create a new `InitiatePaymentBudLicenseEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## InitiatePaymentClientLicenseEntity

```php
$initiate_payment_client_license = $client->InitiatePaymentClientLicense();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_url` | `string` | Yes | The authorisation URL for a relevant payment provider. |
| `code` | `string` | No | Code parameter returned in redirect parameters |
| `payment_details` | `array` | Yes | Payment details required to make a Single Payment |
| `payment_id` | `string` | No | Payment Identifier |
| `payment_type` | `string` | No | The type of payment iniated through Bud's Payments Service |
| `provider` | `string` | Yes | The name (identifier) of the payment provider |
| `provider_redirect_url` | `string` | No | This is the url that your Customer will be redirected to once they have authorised with the relevant provider. |
| `provider_types` | `array` | No | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `string` | No | This is the url that your Customer will be redirected to when using the Bud user interface. |
| `required_action` | `string` | No | If present, this fild will indicate the next required action (if any) to be taken by the client in order to complete the payment flow | Value | Description | |:-------|:-------------| | confirm_scheduled_payment | Indicates that the client… |
| `result` | `string` | No |  |
| `scheduled_payment_details` | `array` | Yes | Details of the scheduled payment |
| `standing_order_details` | `array` | Yes | Details of the standing order |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->InitiatePaymentClientLicense()->create([
  "authorisation_url" => null, // string
  "payment_details" => null, // array
  "provider" => null, // string
  "scheduled_payment_details" => null, // array
  "standing_order_details" => null, // array
  "state" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): InitiatePaymentClientLicenseEntity`

Create a new `InitiatePaymentClientLicenseEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ManagePaymentEntity

```php
$manage_payment = $client->ManagePayment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `array` | Yes | The monetary amount. |
| `client_id` | `string` | Yes | Bud's Client/App Identifier |
| `created_at` | `string` | Yes | The created at timestamp for the payment |
| `data` | `array` | Yes | Data structure for the Scheduled Payment object |
| `errors` | `array` | Yes | A list of any errors that may be associated with the payment |
| `first_payment_date` | `string` | Yes | Date-time (ISO 8601) the first payment should be made. |
| `frequency` | `string` | Yes | The payment frequency for a Standing Order |
| `known_charges` | `array` | Yes | A list detailing information surrounding any known charges that may be associated with the payment |
| `last_payment_date` | `string` | No | Date-time (ISO 8601) after which payments should stop. |
| `metadata` | `array` | No |  |
| `method` | `string` | Yes | The payment method |
| `operation_id` | `string` | Yes | A unique identifier/reference associated with a given endpoint/operation |
| `organisation_id` | `string` | Yes | Bud's Organisation Identifier |
| `payment_id` | `string` | Yes | Payment Identifier |
| `recipient` | `array` | Yes | Details of the recipient of the payment |
| `recurring_amount` | `array` | Yes | The monetary amount. |
| `reference` | `string` | Yes | Reference to use for this payment |
| `requested_execution_date` | `string` | Yes | Date-time (ISO 8601) the Scheduled Payment should be made. |
| `required_action` | `string` | Yes | This field will indicate that there is an action required to be taken by the client in order to complete the payment. |
| `sender` | `array` | Yes | Details of the payment sender |
| `state` | `array` | Yes | Details of the payment state and history |
| `supplementary_status` | `array` | Yes | A list of supplementary states associated with the payment as described by the specific provider |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ManagePayment()->create([
  "scheduled_id" => null, // string
  "amount" => null, // array
  "client_id" => null, // string
  "created_at" => null, // string
  "data" => null, // array
  "errors" => null, // array
  "first_payment_date" => null, // string
  "frequency" => null, // string
  "known_charges" => null, // array
  "method" => null, // string
  "operation_id" => null, // string
  "organisation_id" => null, // string
  "payment_id" => null, // string
  "recipient" => null, // array
  "recurring_amount" => null, // array
  "reference" => null, // string
  "requested_execution_date" => null, // string
  "required_action" => null, // string
  "sender" => null, // array
  "state" => null, // array
  "supplementary_status" => null, // array
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->ManagePayment()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->ManagePayment()->load(["payment_id" => "payment_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ManagePaymentEntity`

Create a new `ManagePaymentEntity` instance with the same client and
options.

#### `get_name(): string`

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

```php
$client = new BudPaymentsSDK([
  "feature" => [
    "debug" => ["active" => true],
    "idempotency" => ["active" => true],
    "metrics" => ["active" => true],
    "paging" => ["active" => true],
    "ratelimit" => ["active" => true],
    "retry" => ["active" => true],
    "test" => ["active" => true],
    "timeout" => ["active" => true],
  ],
]);
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

