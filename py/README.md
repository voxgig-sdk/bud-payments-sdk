# BudPayments Python SDK



The Python SDK for the BudPayments API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.InitiatePaymentBudLicense()` — each
carrying a small, uniform set of operations (`list`, `load`, `create`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/bud-payments-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from budpayments_sdk import BudPaymentsSDK

client = BudPaymentsSDK({
    "apikey": os.environ.get("BUD_PAYMENTS_APIKEY"),
})
```

### 2. List initiatepaymentbudlicense records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    initiatepaymentbudlicenses = client.InitiatePaymentBudLicense().list()
    for initiatepaymentbudlicense in initiatepaymentbudlicenses:
        print(initiatepaymentbudlicense)
except Exception as err:
    print(f"list failed: {err}")
```

### 4. Create, update, and remove

```python
# Create — returns the ENTITY (call data_get() for the record)
created = client.InitiatePaymentBudLicense().create({"bud_pay_url": "example_bud_pay_url", "country_code": "example_country_code", "display_name": "example_display_name", "icon": "example_icon", "implementation_type": "example_implementation_type", "maintenance_status": "example_maintenance_status", "maintenance_window": "example_maintenance_window", "payment_details": {}, "payment_id": "example_payment_id", "provider": "example_provider", "redirect_url": "example_redirect_url", "required_actions": "example_required_actions", "scheduled_payment_details": {}, "services": [], "standing_order_details": {}, "supported_currencies": []})

```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    initiatepaymentbudlicenses = client.InitiatePaymentBudLicense().list()
    print(initiatepaymentbudlicenses)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = BudPaymentsSDK.test()

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
initiatepaymentbudlicense = client.InitiatePaymentBudLicense().list()
# initiatepaymentbudlicense contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = BudPaymentsSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### BudPaymentsSDK

```python
from budpayments_sdk import BudPaymentsSDK

client = BudPaymentsSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = BudPaymentsSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### BudPaymentsSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
| `InitiatePaymentBudLicense` | `(data) -> InitiatePaymentBudLicenseEntity` | Create an InitiatePaymentBudLicense entity instance. |
| `InitiatePaymentClientLicense` | `(data) -> InitiatePaymentClientLicenseEntity` | Create an InitiatePaymentClientLicense entity instance. |
| `ManagePayment` | `(data) -> ManagePaymentEntity` | Create a ManagePayment entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `initiate_payment_bud_license = client.InitiatePaymentBudLicense()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bud_pay_url` | `str` | A Bud Pay URL, allowing Customers to complete (authenticate) a relevant payment resource with the relevant payment service provider. |
| `country_code` | `str` | Provider country code following the ISO 3166 alpha-3 code format |
| `display_name` | `str` | The name of the payment provider to be used for display purposes (e.g. |
| `icon` | `str` | A URL for the icon of the corresponding provider |
| `implementation_type` | `str` | The implementation standard adopted by the provider |
| `maintenance_status` | `str` | The status of the maintenance window associated with the provider |
| `maintenance_window` | `Any` |  |
| `payment_details` | `dict` | Payment details required to make a Single Payment |
| `payment_id` | `str` | Unique identifier for the payment |
| `provider` | `str` | The name (identifier) of the payment provider |
| `provider_types` | `list` | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `str` | URL where the user will be redirected to once they have completed the TPP Connection flow. |
| `required_actions` | `Any` | An array of enum fields to indicate if any additional actions are required in order to complete a given payment flow for the given provider. |
| `scheduled_payment_details` | `dict` | Details of the scheduled payment |
| `services` | `list` | A payment service supported by a payment provider |
| `standing_order_details` | `dict` | Details of the standing order |
| `supported_currencies` | `list` | The currencies supported by the payment provider |

#### Example: List

```python
initiate_payment_bud_licenses = client.InitiatePaymentBudLicense().list()
```

#### Example: Create

```python
initiate_payment_bud_license = client.InitiatePaymentBudLicense().create({
    "bud_pay_url": "example_bud_pay_url",  # str
    "country_code": "example_country_code",  # str
    "display_name": "example_display_name",  # str
    "icon": "example_icon",  # str
    "implementation_type": "example_implementation_type",  # str
    "maintenance_status": "example_maintenance_status",  # str
    "maintenance_window": "example_maintenance_window",  # Any
    "payment_details": {},  # dict
    "payment_id": "example_payment_id",  # str
    "provider": "example_provider",  # str
    "redirect_url": "example_redirect_url",  # str
    "required_actions": "example_required_actions",  # Any
    "scheduled_payment_details": {},  # dict
    "services": [],  # list
    "standing_order_details": {},  # dict
    "supported_currencies": [],  # list
})
```


### InitiatePaymentClientLicense

Create an instance: `initiate_payment_client_license = client.InitiatePaymentClientLicense()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_url` | `str` | The authorisation URL for a relevant payment provider. |
| `code` | `str` | Code parameter returned in redirect parameters |
| `payment_details` | `dict` | Payment details required to make a Single Payment |
| `payment_id` | `str` | Payment Identifier |
| `payment_type` | `str` | The type of payment iniated through Bud's Payments Service |
| `provider` | `str` | The name (identifier) of the payment provider |
| `provider_redirect_url` | `str` | This is the url that your Customer will be redirected to once they have authorised with the relevant provider. |
| `provider_types` | `list` | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `str` | This is the url that your Customer will be redirected to when using the Bud user interface. |
| `required_action` | `str` | If present, this fild will indicate the next required action (if any) to be taken by the client in order to complete the payment flow | Value | Description | |:-------|:-------------| | confirm_scheduled_payment | Indicates that the client… |
| `result` | `str` |  |
| `scheduled_payment_details` | `dict` | Details of the scheduled payment |
| `standing_order_details` | `dict` | Details of the standing order |
| `state` | `str` | The task ID |

#### Example: Create

```python
initiate_payment_client_license = client.InitiatePaymentClientLicense().create({
    "authorisation_url": "example_authorisation_url",  # str
    "payment_details": {},  # dict
    "provider": "example_provider",  # str
    "scheduled_payment_details": {},  # dict
    "standing_order_details": {},  # dict
    "state": "example_state",  # str
})
```


### ManagePayment

Create an instance: `manage_payment = client.ManagePayment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `dict` | The monetary amount. |
| `client_id` | `str` | Bud's Client/App Identifier |
| `created_at` | `str` | The created at timestamp for the payment |
| `data` | `dict` | Data structure for the Scheduled Payment object |
| `errors` | `list` | A list of any errors that may be associated with the payment |
| `first_payment_date` | `str` | Date-time (ISO 8601) the first payment should be made. |
| `frequency` | `str` | The payment frequency for a Standing Order |
| `known_charges` | `dict` | A list detailing information surrounding any known charges that may be associated with the payment |
| `last_payment_date` | `str` | Date-time (ISO 8601) after which payments should stop. |
| `metadata` | `dict` |  |
| `method` | `str` | The payment method |
| `operation_id` | `str` | A unique identifier/reference associated with a given endpoint/operation |
| `organisation_id` | `str` | Bud's Organisation Identifier |
| `payment_id` | `str` | Payment Identifier |
| `recipient` | `dict` | Details of the recipient of the payment |
| `recurring_amount` | `dict` | The monetary amount. |
| `reference` | `str` | Reference to use for this payment |
| `requested_execution_date` | `str` | Date-time (ISO 8601) the Scheduled Payment should be made. |
| `required_action` | `str` | This field will indicate that there is an action required to be taken by the client in order to complete the payment. |
| `sender` | `dict` | Details of the payment sender |
| `state` | `dict` | Details of the payment state and history |
| `supplementary_status` | `dict` | A list of supplementary states associated with the payment as described by the specific provider |

#### Example: Load

```python
manage_payment = client.ManagePayment().load({"payment_id": "payment_id"})
```

#### Example: List

```python
manage_payments = client.ManagePayment().list()
```

#### Example: Create

```python
manage_payment = client.ManagePayment().create({
    "scheduled_id": "example_scheduled_id",  # str
    "amount": {},  # dict
    "client_id": "example_client_id",  # str
    "created_at": "example_created_at",  # str
    "data": {},  # dict
    "errors": [],  # list
    "first_payment_date": "example_first_payment_date",  # str
    "frequency": "example_frequency",  # str
    "known_charges": {},  # dict
    "method": "example_method",  # str
    "operation_id": "example_operation_id",  # str
    "organisation_id": "example_organisation_id",  # str
    "payment_id": "example_payment_id",  # str
    "recipient": {},  # dict
    "recurring_amount": {},  # dict
    "reference": "example_reference",  # str
    "requested_execution_date": "example_requested_execution_date",  # str
    "required_action": "example_required_action",  # str
    "sender": {},  # dict
    "state": {},  # dict
    "supplementary_status": {},  # dict
})
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

Features are the extension mechanism. A feature is a Python class
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

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── budpayments_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── schema.py                    -- Generated option + entity specs
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`budpayments_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
initiatepaymentbudlicense = client.InitiatePaymentBudLicense()
initiatepaymentbudlicense.list()

# initiatepaymentbudlicense.data_get() now returns the initiatepaymentbudlicense data from the last list
# initiatepaymentbudlicense.match_get() returns the last match criteria
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
