# Payments API

These endpoints allow Bud&#39;s Clients to initiate (that is make) payments through the use of Open Banking Payment APIs. Bud manages the connections to each of the payment providers, and serves these connections under a single set of API endpoints, enabling Bud&#39;s clients (and in turn their customers) to initiate a payment from one of these providers to a chosen recipient. Bud currently supports two methods for initiating payments with available providers: 1. If you are not regulated to provide payment initiation services, or if you are regulated but do not wish to use your own permissions, then use Bud&#39;s [TPP Initiation](#tag/TPP-Initiation) endpoints. 2. If on the other hand, your are regulated as an Payment Initiation Service Provider (PISP), then you are able to use Bud as a Technical Service Provider (see [TSP Initiation](#tag/TSP-Initiation)).

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 3 entities and 17 HTTP routes. There are 6 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### [InitiatePaymentBudLicense](docs/api/initiate_payment_bud_license.html)

Results: Successful request.

SDK operations: `create`, `list`.

Key fields to recognise:

- `bud_pay_url`: A Bud Pay URL, allowing Customers to complete (authenticate) a relevant payment resource with the relevant payment service provider.
- `country_code`: Provider country code following the ISO 3166 alpha-3 code format
- `display_name`: The name of the payment provider to be used for display purposes (for example within a frontend application)
- `icon`: A URL for the icon of the corresponding provider
- `implementation_type`: The implementation standard adopted by the provider

### [InitiatePaymentClientLicense](docs/api/initiate_payment_client_license.html)

Results: Successful request.; Successful request.

SDK operations: `create`.

Key fields to recognise:

- `authorisation_url`: The authorisation URL for a relevant payment provider. Redirect your Customer to this url to allow them to authorise the scheduled instruction
- `code`: Code parameter returned in redirect parameters
- `payment_details`: Payment details required to make a Single Payment
- `payment_id`: Payment Identifier
- `payment_type`: The type of payment iniated through Bud&#39;s Payments Service

### [ManagePayment](docs/api/manage_payment.html)

Results: Request to create scheduled payment succeeded; Payment order created successfully; Standing order created successfully; Successful request.

SDK operations: `create`, `list`, `load`.

Key fields to recognise:

- `amount`: The monetary amount.
- `client_id`: Bud&#39;s Client/App Identifier
- `created_at`: Date-time (ISO 8601) of when the Scheduled Payment was created
- `data`: Array of Scheduled Payment Status Objects
- `errors`: Details surrounding any errors when attempting to initiate a Scheduled Payment

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| [InitiatePaymentBudLicense](docs/api/initiate_payment_bud_license.html) | `create` | `POST /v1/payments/scheduled/bud-pay-url` | Required |
| [InitiatePaymentBudLicense](docs/api/initiate_payment_bud_license.html) | `create` | `POST /v1/payments/single/bud-pay-url` | Required |
| [InitiatePaymentBudLicense](docs/api/initiate_payment_bud_license.html) | `create` | `POST /v1/payments/standing-order/bud-pay-url` | Required |
| [InitiatePaymentBudLicense](docs/api/initiate_payment_bud_license.html) | `list` | `GET /v1/payments/providers` | Required |
| [InitiatePaymentClientLicense](docs/api/initiate_payment_client_license.html) | `create` | `POST /v1/payments/authorisation-codes` | Required |
| [InitiatePaymentClientLicense](docs/api/initiate_payment_client_license.html) | `create` | `POST /v1/payments/scheduled` | Required |
| [InitiatePaymentClientLicense](docs/api/initiate_payment_client_license.html) | `create` | `POST /v1/payments/single` | Required |
| [InitiatePaymentClientLicense](docs/api/initiate_payment_client_license.html) | `create` | `POST /v1/payments/standing-order` | Required |
| [ManagePayment](docs/api/manage_payment.html) | `create` | `POST /v1/payments/scheduled/{payment_id}/confirm` | Required |
| [ManagePayment](docs/api/manage_payment.html) | `create` | `POST /v1/payments/single/{payment_id}/confirm` | Required |
| [ManagePayment](docs/api/manage_payment.html) | `create` | `POST /v1/payments/standing-order/{payment_id}/confirm` | Required |
| [ManagePayment](docs/api/manage_payment.html) | `list` | `GET /v1/payments/scheduled` | Required |
| [ManagePayment](docs/api/manage_payment.html) | `list` | `GET /v1/payments/single` | Required |
| [ManagePayment](docs/api/manage_payment.html) | `list` | `GET /v1/payments/standing-order` | Required |
| [ManagePayment](docs/api/manage_payment.html) | `load` | `GET /v1/payments/scheduled/{payment_id}` | Required |
| [ManagePayment](docs/api/manage_payment.html) | `load` | `GET /v1/payments/single/{payment_id}` | Required |
| [ManagePayment](docs/api/manage_payment.html) | `load` | `GET /v1/payments/standing-order/{payment_id}` | Required |

## Connect to the API

- Bud sandbox: `https://api-sandbox.thisisbud.com`

The default credential is sent in the `Authorization` header with the `Bearer` prefix.

Authentication flow: 1. Perform OAuth2 Client Credentials authentication using API Credentials (`client_id`,`client_secret`) to obtain an `access_token` against `/v1/oauth/token` endpoint, 2. Use `access_token` as Bearer Authorisation for every other API request, 3. Include `X-Client-Id` (=client_id) within the header of every API request, 4. Note that some of the requests may also require `X-Customer-Id` to be provided within the request header. ### Examples Obtain OAuth2 `access_token` and `refresh_token` using `grant_type=client_credentials` and HTTP Basic auth header ``` curl --basic --user &#123;&#123;client_id&#125;&#125;:&#123;&#123;client_secret&#125;&#125; \ -X POST https://api-sandbox.thisisbud.com/v1/oauth/token \ -H &#39;Content-Type: application/x-www-form-urlencoded&#39; \ -d grant_type=client_credentials ``` Successful response: ``` &#123; &quot;operation_id&quot;: &quot;oauth_token_post&quot;, &quot;data&quot;: &#123; &quot;access_token&quot;: &quot;dd0c17e3fd6d2ce94aa091257a3ea393b4f9b5cf3d3e998f07dc9826da86ff15&quot;, &quot;token_type&quot;: &quot;bearer&quot;, &quot;expires_in&quot;: 3600, &quot;refresh_token&quot;: &quot;fac32cca7559d9f6e8f1dfe9a99c71fa1dcfeb482bedf287d7934d2667ae54b3&quot; &#125; &#125; ``` Refresh `access_token` token using `refresh_token` against `/v1/oauth/token` endpoint with `grant_type=refresh_token` ``` curl -X POST \ https://api-sandbox.thisisbud.com/v1/oauth/token \ -H &#39;Content-Type: application/x-www-form-urlencoded&#39; \ -H &#39;X-Client-Id: &#123;&#123;client_id&#125;&#125;&#39; \ -d &#39;grant_type=refresh_token&amp;refresh_token=&#123;&#123;refresh_token&#125;&#125;&#39; ``` Successful response: ``` &#123; &quot;operation_id&quot;: &quot;oauth_token_post&quot;, &quot;data&quot;: &#123; &quot;access_token&quot;: &quot;cc0c17e3fd6d2ce94aa091257a3ea393b4f9b5cf3d3e998f07dc9826da86ff94&quot;, &quot;token_type&quot;: &quot;bearer&quot;, &quot;expires_in&quot;: 3600, &quot;refresh_token&quot;: &quot;ffc30cca7559d9f6e8f1dfe9a99c71fa1dcfeb482bedf287d7934d2667ae54b3&quot; &#125; &#125; ```

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| [Golang](docs/sdks/go.html) | `go/` | Build from source |
| [JavaScript](docs/sdks/js.html) | `js/` | Build from source |
| [Lua](docs/sdks/lua.html) | `lua/` | Build from source |
| [PHP](docs/sdks/php.html) | `php/` | Build from source |
| [Python](docs/sdks/py.html) | `py/` | Build from source |
| [TypeScript](docs/sdks/ts.html) | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### [Go CLI](docs/tools/go-cli.html)

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### [Go MCP server](docs/tools/go-mcp.html)

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `bud-payments_list`: List records for an entity. Supported entities: `initiate_payment_bud_license`, `manage_payment`.
- `bud-payments_load`: Load one record for an entity. Supported entities: `manage_payment`.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- [`debug`](docs/features/debug.html): Request/response capture ring buffer for debugging
- [`idempotency`](docs/features/idempotency.html): Idempotency keys for safe retries of mutating operations
- [`metrics`](docs/features/metrics.html): Statistics capture: per-operation counters and latency
- [`paging`](docs/features/paging.html): Pagination signals for list operations
- [`ratelimit`](docs/features/ratelimit.html): Client-side rate limiting via a token bucket
- [`retry`](docs/features/retry.html): Automatic retry of transient failures with exponential backoff
- [`test`](docs/features/test.html): In-memory mock transport for testing without a live server
- [`timeout`](docs/features/timeout.html): Per-request timeout with transport abort

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the [first-call guide](docs/guides/first-call.html) for the setup sequence.
- Read the [authentication guide](docs/guides/authentication.html) before using protected routes.
- Use the [API reference](docs/api/index.html) for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

