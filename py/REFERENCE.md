# BudPayments Python SDK Reference

Complete API reference for the BudPayments Python SDK.


## BudPaymentsSDK

### Constructor

```python
from budpayments_sdk import BudPaymentsSDK

client = BudPaymentsSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `BudPaymentsSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = BudPaymentsSDK.test()
```


### Instance Methods

#### `InitiatePaymentBudLicense(data=None)`

Create a new `InitiatePaymentBudLicenseEntity` instance. Pass `None` for no initial data.

#### `InitiatePaymentClientLicense(data=None)`

Create a new `InitiatePaymentClientLicenseEntity` instance. Pass `None` for no initial data.

#### `ManagePayment(data=None)`

Create a new `ManagePaymentEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## InitiatePaymentBudLicenseEntity

```python
initiate_payment_bud_license = client.InitiatePaymentBudLicense()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bud_pay_url` | `str` | Yes | A Bud Pay URL, allowing Customers to complete (authenticate) a relevant payment resource with the relevant payment service provider. |
| `country_code` | `str` | Yes | Provider country code following the ISO 3166 alpha-3 code format |
| `display_name` | `str` | Yes | The name of the payment provider to be used for display purposes (e.g. |
| `icon` | `str` | Yes | A URL for the icon of the corresponding provider |
| `implementation_type` | `str` | Yes | The implementation standard adopted by the provider |
| `maintenance_status` | `str` | Yes | The status of the maintenance window associated with the provider |
| `maintenance_window` | `Any` | Yes |  |
| `payment_details` | `dict` | Yes | Payment details required to make a Single Payment |
| `payment_id` | `str` | Yes | Unique identifier for the payment |
| `provider` | `str` | Yes | The name (identifier) of the payment provider |
| `provider_types` | `list` | No | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `str` | Yes | URL where the user will be redirected to once they have completed the TPP Connection flow. |
| `required_actions` | `Any` | Yes | An array of enum fields to indicate if any additional actions are required in order to complete a given payment flow for the given provider. |
| `scheduled_payment_details` | `dict` | Yes | Details of the scheduled payment |
| `services` | `list` | Yes | A payment service supported by a payment provider |
| `standing_order_details` | `dict` | Yes | Details of the standing order |
| `supported_currencies` | `list` | Yes | The currencies supported by the payment provider |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.InitiatePaymentBudLicense().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.InitiatePaymentBudLicense().list()
for initiate_payment_bud_license in results:
    print(initiate_payment_bud_license)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InitiatePaymentBudLicenseEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InitiatePaymentClientLicenseEntity

```python
initiate_payment_client_license = client.InitiatePaymentClientLicense()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_url` | `str` | Yes | The authorisation URL for a relevant payment provider. |
| `code` | `str` | No | Code parameter returned in redirect parameters |
| `payment_details` | `dict` | Yes | Payment details required to make a Single Payment |
| `payment_id` | `str` | No | Payment Identifier |
| `payment_type` | `str` | No | The type of payment iniated through Bud's Payments Service |
| `provider` | `str` | Yes | The name (identifier) of the payment provider |
| `provider_redirect_url` | `str` | No | This is the url that your Customer will be redirected to once they have authorised with the relevant provider. |
| `provider_types` | `list` | No | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `str` | No | This is the url that your Customer will be redirected to when using the Bud user interface. |
| `required_action` | `str` | No | If present, this fild will indicate the next required action (if any) to be taken by the client in order to complete the payment flow | Value | Description | |:-------|:-------------| | confirm_scheduled_payment | Indicates that the client… |
| `result` | `str` | No |  |
| `scheduled_payment_details` | `dict` | Yes | Details of the scheduled payment |
| `standing_order_details` | `dict` | Yes | Details of the standing order |
| `state` | `str` | Yes | The task ID |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.InitiatePaymentClientLicense().create({
    "authorisation_url": "example_authorisation_url",  # str
    "payment_details": {},  # dict
    "provider": "example_provider",  # str
    "scheduled_payment_details": {},  # dict
    "standing_order_details": {},  # dict
    "state": "example_state",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InitiatePaymentClientLicenseEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ManagePaymentEntity

```python
manage_payment = client.ManagePayment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `dict` | Yes | The monetary amount. |
| `client_id` | `str` | Yes | Bud's Client/App Identifier |
| `created_at` | `str` | Yes | The created at timestamp for the payment |
| `data` | `dict` | Yes | Data structure for the Scheduled Payment object |
| `errors` | `list` | Yes | A list of any errors that may be associated with the payment |
| `first_payment_date` | `str` | Yes | Date-time (ISO 8601) the first payment should be made. |
| `frequency` | `str` | Yes | The payment frequency for a Standing Order |
| `known_charges` | `dict` | Yes | A list detailing information surrounding any known charges that may be associated with the payment |
| `last_payment_date` | `str` | No | Date-time (ISO 8601) after which payments should stop. |
| `metadata` | `dict` | No |  |
| `method` | `str` | Yes | The payment method |
| `operation_id` | `str` | Yes | A unique identifier/reference associated with a given endpoint/operation |
| `organisation_id` | `str` | Yes | Bud's Organisation Identifier |
| `payment_id` | `str` | Yes | Payment Identifier |
| `recipient` | `dict` | Yes | Details of the recipient of the payment |
| `recurring_amount` | `dict` | Yes | The monetary amount. |
| `reference` | `str` | Yes | Reference to use for this payment |
| `requested_execution_date` | `str` | Yes | Date-time (ISO 8601) the Scheduled Payment should be made. |
| `required_action` | `str` | Yes | This field will indicate that there is an action required to be taken by the client in order to complete the payment. |
| `sender` | `dict` | Yes | Details of the payment sender |
| `state` | `dict` | Yes | Details of the payment state and history |
| `supplementary_status` | `dict` | Yes | A list of supplementary states associated with the payment as described by the specific provider |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ManagePayment().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.ManagePayment().list()
for manage_payment in results:
    print(manage_payment)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.ManagePayment().load({"payment_id": "payment_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ManagePaymentEntity` instance with the same options.

#### `get_name() -> str`

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

```python
client = BudPaymentsSDK({
    "feature": {
        "debug": {"active": True},
        "idempotency": {"active": True},
        "metrics": {"active": True},
        "paging": {"active": True},
        "ratelimit": {"active": True},
        "retry": {"active": True},
        "test": {"active": True},
        "timeout": {"active": True},
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

