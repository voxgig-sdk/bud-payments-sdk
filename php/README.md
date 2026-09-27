# BudPayments PHP SDK



The PHP SDK for the BudPayments API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->InitiatePaymentBudLicense()` — with named operations (`list`/`load`/`create`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/bud-payments-sdk/releases](https://github.com/voxgig-sdk/bud-payments-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'budpayments_sdk.php';

$client = new BudPaymentsSDK([
    "apikey" => getenv("BUD_PAYMENTS_APIKEY"),
]);
```

### 2. List initiatepaymentbudlicense records

```php
try {
    // list() returns entity instances; data_get() reads each record.
    $initiatepaymentbudlicenses = $client->InitiatePaymentBudLicense()->list();
    foreach ($initiatepaymentbudlicenses as $record) {
        $item = $record->data_get();
        echo $item["bud_pay_url"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 4. Create, update, and remove

```php
// create() returns the ENTITY — call data_get() for the created InitiatePaymentBudLicense record.
$created = $client->InitiatePaymentBudLicense()->create(["bud_pay_url" => "example_bud_pay_url", "country_code" => "example_country_code", "display_name" => "example_display_name", "icon" => "example_icon", "implementation_type" => "example_implementation_type", "maintenance_status" => "example_maintenance_status", "maintenance_window" => "example_maintenance_window", "payment_details" => [], "payment_id" => "example_payment_id", "provider" => "example_provider", "redirect_url" => "example_redirect_url", "required_actions" => "example_required_actions", "scheduled_payment_details" => [], "services" => [], "standing_order_details" => [], "supported_currencies" => []]);

```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $initiatepaymentbudlicenses = $client->InitiatePaymentBudLicense()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required:

```php
$client = BudPaymentsSDK::test();

// list() returns entity instances (throws on error);
// call data_get() for the mock record.
$initiatepaymentbudlicense = $client->InitiatePaymentBudLicense()->list();
print_r(array_map(fn($item) => $item->data_get(), $initiatepaymentbudlicense));
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new BudPaymentsSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
BUD_PAYMENTS_TEST_LIVE=TRUE
BUD_PAYMENTS_APIKEY=<your-key>
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### BudPaymentsSDK

```php
require_once 'budpayments_sdk.php';
$client = new BudPaymentsSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = BudPaymentsSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### BudPaymentsSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `InitiatePaymentBudLicense` | `($data): InitiatePaymentBudLicenseEntity` | Create an InitiatePaymentBudLicense entity instance. |
| `InitiatePaymentClientLicense` | `($data): InitiatePaymentClientLicenseEntity` | Create an InitiatePaymentClientLicense entity instance. |
| `ManagePayment` | `($data): ManagePaymentEntity` | Create a ManagePayment entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

### Entities

#### InitiatePaymentBudLicense

| Field | Description |
| --- | --- |
| `bud_pay_url` | A Bud Pay URL, allowing Customers to complete (authenticate) a relevant payment resource with the relevant payment service provider. |
| `country_code` | Provider country code following the ISO 3166 alpha-3 code format |
| `display_name` | The name of the payment provider to be used for display purposes (e.g. |
| `icon` | A URL for the icon of the corresponding provider |
| `implementation_type` | The implementation standard adopted by the provider |
| `maintenance_status` | The status of the maintenance window associated with the provider |
| `maintenance_window` |  |
| `payment_details` | Payment details required to make a Single Payment |
| `payment_id` | Unique identifier for the payment |
| `provider` | The name (identifier) of the payment provider |
| `provider_types` | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | URL where the user will be redirected to once they have completed the TPP Connection flow. |
| `required_actions` | An array of enum fields to indicate if any additional actions are required in order to complete a given payment flow for the given provider. |
| `scheduled_payment_details` | Details of the scheduled payment |
| `services` | A payment service supported by a payment provider |
| `standing_order_details` | Details of the standing order |
| `supported_currencies` | The currencies supported by the payment provider |

Operations: Create, List.

API path: `/v1/payments/scheduled/bud-pay-url`

#### InitiatePaymentClientLicense

| Field | Description |
| --- | --- |
| `authorisation_url` | The authorisation URL for a relevant payment provider. |
| `code` | Code parameter returned in redirect parameters |
| `payment_details` | Payment details required to make a Single Payment |
| `payment_id` | Payment Identifier |
| `payment_type` | The type of payment iniated through Bud's Payments Service |
| `provider` | The name (identifier) of the payment provider |
| `provider_redirect_url` | This is the url that your Customer will be redirected to once they have authorised with the relevant provider. |
| `provider_types` | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | This is the url that your Customer will be redirected to when using the Bud user interface. |
| `required_action` | If present, this fild will indicate the next required action (if any) to be taken by the client in order to complete the payment flow | Value | Description | |:-------|:-------------| | confirm_scheduled_payment | Indicates that the client… |
| `result` |  |
| `scheduled_payment_details` | Details of the scheduled payment |
| `standing_order_details` | Details of the standing order |
| `state` | The task ID |

Operations: Create.

API path: `/v1/payments/authorisation-codes`

#### ManagePayment

| Field | Description |
| --- | --- |
| `amount` | The monetary amount. |
| `client_id` | Bud's Client/App Identifier |
| `created_at` | The created at timestamp for the payment |
| `data` | Data structure for the Scheduled Payment object |
| `errors` | A list of any errors that may be associated with the payment |
| `first_payment_date` | Date-time (ISO 8601) the first payment should be made. |
| `frequency` | The payment frequency for a Standing Order |
| `known_charges` | A list detailing information surrounding any known charges that may be associated with the payment |
| `last_payment_date` | Date-time (ISO 8601) after which payments should stop. |
| `metadata` |  |
| `method` | The payment method |
| `operation_id` | A unique identifier/reference associated with a given endpoint/operation |
| `organisation_id` | Bud's Organisation Identifier |
| `payment_id` | Payment Identifier |
| `recipient` | Details of the recipient of the payment |
| `recurring_amount` | The monetary amount. |
| `reference` | Reference to use for this payment |
| `requested_execution_date` | Date-time (ISO 8601) the Scheduled Payment should be made. |
| `required_action` | This field will indicate that there is an action required to be taken by the client in order to complete the payment. |
| `sender` | Details of the payment sender |
| `state` | Details of the payment state and history |
| `supplementary_status` | A list of supplementary states associated with the payment as described by the specific provider |

Operations: Create, List, Load.

API path: `/v1/payments/scheduled/{payment_id}/confirm`



## Entities


### InitiatePaymentBudLicense

Create an instance: `$initiate_payment_bud_license = $client->InitiatePaymentBudLicense();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bud_pay_url` | `string` | A Bud Pay URL, allowing Customers to complete (authenticate) a relevant payment resource with the relevant payment service provider. |
| `country_code` | `string` | Provider country code following the ISO 3166 alpha-3 code format |
| `display_name` | `string` | The name of the payment provider to be used for display purposes (e.g. |
| `icon` | `string` | A URL for the icon of the corresponding provider |
| `implementation_type` | `string` | The implementation standard adopted by the provider |
| `maintenance_status` | `string` | The status of the maintenance window associated with the provider |
| `maintenance_window` | `mixed` |  |
| `payment_details` | `array` | Payment details required to make a Single Payment |
| `payment_id` | `string` | Unique identifier for the payment |
| `provider` | `string` | The name (identifier) of the payment provider |
| `provider_types` | `array` | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `string` | URL where the user will be redirected to once they have completed the TPP Connection flow. |
| `required_actions` | `mixed` | An array of enum fields to indicate if any additional actions are required in order to complete a given payment flow for the given provider. |
| `scheduled_payment_details` | `array` | Details of the scheduled payment |
| `services` | `array` | A payment service supported by a payment provider |
| `standing_order_details` | `array` | Details of the standing order |
| `supported_currencies` | `array` | The currencies supported by the payment provider |

#### Example: List

```php
// list() returns an array of InitiatePaymentBudLicense records (throws on error).
$initiate_payment_bud_licenses = $client->InitiatePaymentBudLicense()->list();
```

#### Example: Create

```php
$initiate_payment_bud_license = $client->InitiatePaymentBudLicense()->create([
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


### InitiatePaymentClientLicense

Create an instance: `$initiate_payment_client_license = $client->InitiatePaymentClientLicense();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_url` | `string` | The authorisation URL for a relevant payment provider. |
| `code` | `string` | Code parameter returned in redirect parameters |
| `payment_details` | `array` | Payment details required to make a Single Payment |
| `payment_id` | `string` | Payment Identifier |
| `payment_type` | `string` | The type of payment iniated through Bud's Payments Service |
| `provider` | `string` | The name (identifier) of the payment provider |
| `provider_redirect_url` | `string` | This is the url that your Customer will be redirected to once they have authorised with the relevant provider. |
| `provider_types` | `array` | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `string` | This is the url that your Customer will be redirected to when using the Bud user interface. |
| `required_action` | `string` | If present, this fild will indicate the next required action (if any) to be taken by the client in order to complete the payment flow | Value | Description | |:-------|:-------------| | confirm_scheduled_payment | Indicates that the client… |
| `result` | `string` |  |
| `scheduled_payment_details` | `array` | Details of the scheduled payment |
| `standing_order_details` | `array` | Details of the standing order |
| `state` | `string` | The task ID |

#### Example: Create

```php
$initiate_payment_client_license = $client->InitiatePaymentClientLicense()->create([
    "authorisation_url" => null, // string
    "payment_details" => null, // array
    "provider" => null, // string
    "scheduled_payment_details" => null, // array
    "standing_order_details" => null, // array
    "state" => null, // string
]);
```


### ManagePayment

Create an instance: `$manage_payment = $client->ManagePayment();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `array` | The monetary amount. |
| `client_id` | `string` | Bud's Client/App Identifier |
| `created_at` | `string` | The created at timestamp for the payment |
| `data` | `array` | Data structure for the Scheduled Payment object |
| `errors` | `array` | A list of any errors that may be associated with the payment |
| `first_payment_date` | `string` | Date-time (ISO 8601) the first payment should be made. |
| `frequency` | `string` | The payment frequency for a Standing Order |
| `known_charges` | `array` | A list detailing information surrounding any known charges that may be associated with the payment |
| `last_payment_date` | `string` | Date-time (ISO 8601) after which payments should stop. |
| `metadata` | `array` |  |
| `method` | `string` | The payment method |
| `operation_id` | `string` | A unique identifier/reference associated with a given endpoint/operation |
| `organisation_id` | `string` | Bud's Organisation Identifier |
| `payment_id` | `string` | Payment Identifier |
| `recipient` | `array` | Details of the recipient of the payment |
| `recurring_amount` | `array` | The monetary amount. |
| `reference` | `string` | Reference to use for this payment |
| `requested_execution_date` | `string` | Date-time (ISO 8601) the Scheduled Payment should be made. |
| `required_action` | `string` | This field will indicate that there is an action required to be taken by the client in order to complete the payment. |
| `sender` | `array` | Details of the payment sender |
| `state` | `array` | Details of the payment state and history |
| `supplementary_status` | `array` | A list of supplementary states associated with the payment as described by the specific provider |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the ManagePayment record (throws on error).
$manage_payment = $client->ManagePayment()->load(["payment_id" => "payment_id"]);
```

#### Example: List

```php
// list() returns an array of ManagePayment records (throws on error).
$manage_payments = $client->ManagePayment()->list();
```

#### Example: Create

```php
$manage_payment = $client->ManagePayment()->create([
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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

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

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── budpayments_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── schema.php                     -- Generated option + entity specs
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`budpayments_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$initiatepaymentbudlicense = $client->InitiatePaymentBudLicense();
$initiatepaymentbudlicense->list();

// $initiatepaymentbudlicense->data_get() now returns the initiatepaymentbudlicense data from the last list
// $initiatepaymentbudlicense->match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
