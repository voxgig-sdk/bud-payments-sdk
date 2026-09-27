# BudPayments JavaScript SDK



The JavaScript SDK for the BudPayments API — an entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.InitiatePaymentBudLicense()` — each with a small set of operations (`list`, `load`, `create`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```js
npm install bud-payments
```
## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.


### Create a Client

```js
const { BudPaymentsSDK } = require('@voxgig-sdk/bud-payments-sdk-js')

const client = new BudPaymentsSDK({
  apikey: process.env.BUD_PAYMENTS_APIKEY,
})
```

### List InitiatePaymentBudLicense Records

```js
const initiate_payment_bud_licenses = await client.InitiatePaymentBudLicense().list()
for (const initiate_payment_bud_license of initiate_payment_bud_licenses) {
  console.log(initiate_payment_bud_license)
}
```

### Create a InitiatePaymentBudLicense

```js
const created = await client.InitiatePaymentBudLicense().create({
  bud_pay_url: 'example_bud_pay_url',
  country_code: 'example_country_code',
  display_name: 'example_display_name',
  icon: 'example_icon',
  implementation_type: 'example_implementation_type',
  maintenance_status: 'example_maintenance_status',
  maintenance_window: 'example_maintenance_window',
  payment_details: {},
  payment_id: 'example_payment_id',
  provider: 'example_provider',
  redirect_url: 'example_redirect_url',
  required_actions: 'example_required_actions',
  scheduled_payment_details: {},
  services: [],
  standing_order_details: {},
  supported_currencies: [],
})
console.log(created)
```

### Direct API Access

Use `client.direct()` to call any API endpoint directly:

```js
const result = await client.direct({
  path: '/custom/endpoint/{id}',
  method: 'GET',
  params: { id: 'abc123' },
})

if (result.ok) {
  console.log(result.data)
}
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const initiatepaymentbudlicenses = await client.InitiatePaymentBudLicense().list()
  console.log(initiatepaymentbudlicenses)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```js
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```js
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```js
const client = BudPaymentsSDK.test()

const initiatepaymentbudlicense = await client.InitiatePaymentBudLicense().list()
// initiatepaymentbudlicense is the entity, populated with mock response data
// — call initiatepaymentbudlicense.data() for the record itself
console.log(initiatepaymentbudlicense)
```

You can also use the instance method:

```js
const client = new BudPaymentsSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```js
const entity = client.InitiatePaymentBudLicense()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data)
```

### Add custom middleware

Pass features via the `extend` option:

```js
const logger = {
  hooks: {
    PreRequest: (ctx) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new BudPaymentsSDK({
  apikey: '...',
  extend: [logger],
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
cd js && npm test
```


## Reference

### BudPaymentsSDK

#### Constructor

```js
new BudPaymentsSDK(options?)
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `InitiatePaymentBudLicense(data?)` | `InitiatePaymentBudLicenseEntity` | Create an InitiatePaymentBudLicense entity instance. |
| `InitiatePaymentClientLicense(data?)` | `InitiatePaymentClientLicenseEntity` | Create an InitiatePaymentClientLicense entity instance. |
| `ManagePayment(data?)` | `ManagePaymentEntity` | Create a ManagePayment entity instance. |
| `tester(testopts?, sdkopts?)` | `BudPaymentsSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `BudPaymentsSDK.test(testopts?, sdkopts?)` | `BudPaymentsSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): BudPaymentsSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load` and `create` resolve to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```js
{
  ok: true,
  status: 200,
  headers: {},
  data: {}
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```js
{
  url: 'string',
  method: 'string',
  headers: {},
  body: undefined
}
```

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

Operations: create, list.

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

Operations: create.

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

Operations: create, list, load.

API path: `/v1/payments/scheduled/{payment_id}/confirm`



## Entities


### InitiatePaymentBudLicense

Create an instance: `const initiate_payment_bud_license = client.InitiatePaymentBudLicense()`

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
| `maintenance_window` | `*` |  |
| `payment_details` | `Object` | Payment details required to make a Single Payment |
| `payment_id` | `string` | Unique identifier for the payment |
| `provider` | `string` | The name (identifier) of the payment provider |
| `provider_types` | `Array` | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `string` | URL where the user will be redirected to once they have completed the TPP Connection flow. |
| `required_actions` | `*` | An array of enum fields to indicate if any additional actions are required in order to complete a given payment flow for the given provider. |
| `scheduled_payment_details` | `Object` | Details of the scheduled payment |
| `services` | `Array` | A payment service supported by a payment provider |
| `standing_order_details` | `Object` | Details of the standing order |
| `supported_currencies` | `Array` | The currencies supported by the payment provider |

#### Example: List

```ts
const initiate_payment_bud_licenses = await client.InitiatePaymentBudLicense().list()
```

#### Example: Create

```ts
const initiate_payment_bud_license = await client.InitiatePaymentBudLicense().create({
  bud_pay_url: 'example_bud_pay_url',
  country_code: 'example_country_code',
  display_name: 'example_display_name',
  icon: 'example_icon',
  implementation_type: 'example_implementation_type',
  maintenance_status: 'example_maintenance_status',
  maintenance_window: 'example_maintenance_window',
  payment_details: {},
  payment_id: 'example_payment_id',
  provider: 'example_provider',
  redirect_url: 'example_redirect_url',
  required_actions: 'example_required_actions',
  scheduled_payment_details: {},
  services: [],
  standing_order_details: {},
  supported_currencies: [],
})
```


### InitiatePaymentClientLicense

Create an instance: `const initiate_payment_client_license = client.InitiatePaymentClientLicense()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_url` | `string` | The authorisation URL for a relevant payment provider. |
| `code` | `string` | Code parameter returned in redirect parameters |
| `payment_details` | `Object` | Payment details required to make a Single Payment |
| `payment_id` | `string` | Payment Identifier |
| `payment_type` | `string` | The type of payment iniated through Bud's Payments Service |
| `provider` | `string` | The name (identifier) of the payment provider |
| `provider_redirect_url` | `string` | This is the url that your Customer will be redirected to once they have authorised with the relevant provider. |
| `provider_types` | `Array` | One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. |
| `redirect_url` | `string` | This is the url that your Customer will be redirected to when using the Bud user interface. |
| `required_action` | `string` | If present, this fild will indicate the next required action (if any) to be taken by the client in order to complete the payment flow | Value | Description | |:-------|:-------------| | confirm_scheduled_payment | Indicates that the client… |
| `result` | `string` |  |
| `scheduled_payment_details` | `Object` | Details of the scheduled payment |
| `standing_order_details` | `Object` | Details of the standing order |
| `state` | `string` | The task ID |

#### Example: Create

```ts
const initiate_payment_client_license = await client.InitiatePaymentClientLicense().create({
  authorisation_url: 'example_authorisation_url',
  payment_details: {},
  provider: 'example_provider',
  scheduled_payment_details: {},
  standing_order_details: {},
  state: 'example_state',
})
```


### ManagePayment

Create an instance: `const manage_payment = client.ManagePayment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `Object` | The monetary amount. |
| `client_id` | `string` | Bud's Client/App Identifier |
| `created_at` | `string` | The created at timestamp for the payment |
| `data` | `Object` | Data structure for the Scheduled Payment object |
| `errors` | `Array` | A list of any errors that may be associated with the payment |
| `first_payment_date` | `string` | Date-time (ISO 8601) the first payment should be made. |
| `frequency` | `string` | The payment frequency for a Standing Order |
| `known_charges` | `Object` | A list detailing information surrounding any known charges that may be associated with the payment |
| `last_payment_date` | `string` | Date-time (ISO 8601) after which payments should stop. |
| `metadata` | `Object` |  |
| `method` | `string` | The payment method |
| `operation_id` | `string` | A unique identifier/reference associated with a given endpoint/operation |
| `organisation_id` | `string` | Bud's Organisation Identifier |
| `payment_id` | `string` | Payment Identifier |
| `recipient` | `Object` | Details of the recipient of the payment |
| `recurring_amount` | `Object` | The monetary amount. |
| `reference` | `string` | Reference to use for this payment |
| `requested_execution_date` | `string` | Date-time (ISO 8601) the Scheduled Payment should be made. |
| `required_action` | `string` | This field will indicate that there is an action required to be taken by the client in order to complete the payment. |
| `sender` | `Object` | Details of the payment sender |
| `state` | `Object` | Details of the payment state and history |
| `supplementary_status` | `Object` | A list of supplementary states associated with the payment as described by the specific provider |

#### Example: Load

```ts
const manage_payment = await client.ManagePayment().load({ payment_id: 'payment_id' })
```

#### Example: List

```ts
const manage_payments = await client.ManagePayment().list()
```

#### Example: Create

```ts
const manage_payment = await client.ManagePayment().create({
  scheduled_id: 'example_scheduled_id',
  amount: {},
  client_id: 'example_client_id',
  created_at: 'example_created_at',
  data: {},
  errors: [],
  first_payment_date: 'example_first_payment_date',
  frequency: 'example_frequency',
  known_charges: {},
  method: 'example_method',
  operation_id: 'example_operation_id',
  organisation_id: 'example_organisation_id',
  payment_id: 'example_payment_id',
  recipient: {},
  recurring_amount: {},
  reference: 'example_reference',
  requested_execution_date: 'example_requested_execution_date',
  required_action: 'example_required_action',
  sender: {},
  state: {},
  supplementary_status: {},
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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

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

### Module structure

```
bud-payments/
├── src/
│   ├── BudPaymentsSDK.js        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
└── test/                   # Test suites
```

Import the SDK from the package root:

```js
const { BudPaymentsSDK } = require('@voxgig-sdk/bud-payments-sdk-js')
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const initiatepaymentbudlicense = client.InitiatePaymentBudLicense()
await initiatepaymentbudlicense.list()

// initiatepaymentbudlicense.data() now returns the initiatepaymentbudlicense data from the last `list`
// initiatepaymentbudlicense.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
