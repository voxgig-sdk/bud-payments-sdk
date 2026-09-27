# BudPayments SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "BudPayments",
            "slug": "bud-payments",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "debug": {
        "options": {
          "active": False,
          "max": 100,
          "redact": [
            "authorization",
            "cookie",
            "set-cookie",
            "api-key",
            "apikey",
            "x-api-key",
            "idempotency-key",
          ],
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "onEntry": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "idempotency": {
        "options": {
          "active": False,
          "header": "Idempotency-Key",
          "methods": [
            "POST",
            "PUT",
            "PATCH",
            "DELETE",
          ],
          "ops": [
            "create",
            "update",
            "remove",
          ],
        },
        "optspec": {
          "keygen": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "metrics": {
        "options": {
          "active": False,
        },
        "optspec": {
          "now": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "paging": {
        "options": {
          "active": False,
          "afterVar": "after",
          "cursorParam": "cursor",
          "firstVar": "first",
          "limitParam": "limit",
          "pageParam": "page",
          "startPage": 1,
        },
        "optspec": {
          "limit": "`$NUMBER`",
          "ops": "`$LIST`",
        },
        "strict": False,
        "transport": "none",
      },
            "ratelimit": {
        "options": {
          "active": False,
          "burst": 5,
          "rate": 5,
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "retry": {
        "options": {
          "active": False,
          "factor": 2,
          "maxDelay": 2000,
          "minDelay": 50,
          "retries": 2,
          "statuses": [
            408,
            425,
            429,
            500,
            502,
            503,
            504,
          ],
        },
        "optspec": {
          "jitter": "`$BOOLEAN`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "test": {
        "options": {
          "active": False,
        },
        "optspec": {
          "entity": "`$MAP`",
          "net": "`$MAP`",
        },
        "strict": False,
        "transport": "base",
      },
            "timeout": {
        "options": {
          "active": False,
          "ms": 30000,
        },
        "optspec": {
          "clearTimer": "`$FUNCTION`",
          "setTimer": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
        },
        "options": {
            "base": "https://api-sandbox.thisisbud.com",
            "auth": {
                "prefix": "Bearer",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "initiate_payment_bud_license": {},
                "initiate_payment_client_license": {},
                "manage_payment": {},
            },
        },
        "entity": {
      "initiate_payment_bud_license": {
        "fields": [
          {
            "name": "bud_pay_url",
            "title": "Bud Pay Url",
            "type": "`$STRING`",
            "req": True,
            "short": "A Bud Pay URL, allowing Customers to complete (authenticate) a relevant payment resource with the relevant payment service provider.",
          },
          {
            "name": "country_code",
            "title": "Country Code",
            "type": "`$STRING`",
            "req": True,
            "short": "Provider country code following the ISO 3166 alpha-3 code format",
          },
          {
            "name": "display_name",
            "title": "Display Name",
            "type": "`$STRING`",
            "req": True,
            "short": "The name of the payment provider to be used for display purposes (e.g.",
          },
          {
            "name": "icon",
            "title": "Icon",
            "type": "`$STRING`",
            "req": True,
            "short": "A URL for the icon of the corresponding provider",
          },
          {
            "name": "implementation_type",
            "title": "Implementation Type",
            "type": "`$STRING`",
            "req": True,
            "short": "The implementation standard adopted by the provider",
          },
          {
            "name": "maintenance_status",
            "title": "Maintenance Status",
            "type": "`$STRING`",
            "req": True,
            "short": "The status of the maintenance window associated with the provider",
          },
          {
            "name": "maintenance_window",
            "title": "Maintenance Window",
            "type": "`$ANY`",
            "req": True,
          },
          {
            "name": "payment_details",
            "title": "Payment Details",
            "type": "`$OBJECT`",
            "req": True,
            "short": "Payment details required to make a Single Payment",
          },
          {
            "name": "payment_id",
            "title": "Payment Id",
            "type": "`$STRING`",
            "req": True,
            "short": "Unique identifier for the payment",
          },
          {
            "name": "provider",
            "title": "Provider",
            "type": "`$STRING`",
            "req": True,
            "short": "The name (identifier) of the payment provider",
          },
          {
            "name": "provider_types",
            "title": "Provider Types",
            "type": "`$ARRAY`",
            "short": "One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow.",
          },
          {
            "name": "redirect_url",
            "title": "Redirect Url",
            "type": "`$STRING`",
            "req": True,
            "short": "URL where the user will be redirected to once they have completed the TPP Connection flow.",
          },
          {
            "name": "required_actions",
            "title": "Required Actions",
            "type": "`$ANY`",
            "req": True,
            "short": "An array of enum fields to indicate if any additional actions are required in order to complete a given payment flow for the given provider.",
          },
          {
            "name": "scheduled_payment_details",
            "title": "Scheduled Payment Details",
            "type": "`$OBJECT`",
            "req": True,
            "short": "Details of the scheduled payment",
          },
          {
            "name": "services",
            "title": "Services",
            "type": "`$ARRAY`",
            "req": True,
            "short": "A payment service supported by a payment provider",
          },
          {
            "name": "standing_order_details",
            "title": "Standing Order Details",
            "type": "`$OBJECT`",
            "req": True,
            "short": "Details of the standing order",
          },
          {
            "name": "supported_currencies",
            "title": "Supported Currencies",
            "type": "`$ARRAY`",
            "req": True,
            "short": "The currencies supported by the payment provider",
          },
        ],
        "name": "initiate_payment_bud_license",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "kind": "http",
                "method": "POST",
                "orig": "/v1/payments/scheduled/bud-pay-url",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "scheduled",
                  },
                  {
                    "lit": "bud-pay-url",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "scheduled",
                  "bud-pay-url",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "POST",
                "orig": "/v1/payments/single/bud-pay-url",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "single",
                  },
                  {
                    "lit": "bud-pay-url",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "single",
                  "bud-pay-url",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "POST",
                "orig": "/v1/payments/standing-order/bud-pay-url",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "standing-order",
                  },
                  {
                    "lit": "bud-pay-url",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "standing-order",
                  "bud-pay-url",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                  ],
                },
              },
            ],
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v1/payments/providers",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "providers",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "providers",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                  "query": [
                    {
                      "name": "country",
                      "orig": "country",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "GBR,DEU",
                    },
                    {
                      "name": "product",
                      "orig": "product",
                      "type": "`$ARRAY`",
                      "kind": "query",
                    },
                    {
                      "name": "type",
                      "orig": "type",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "country",
                    "product",
                    "type",
                    "x_client_id",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "initiate_payment_client_license": {
        "fields": [
          {
            "name": "authorisation_url",
            "title": "Authorisation Url",
            "type": "`$STRING`",
            "req": True,
            "short": "The authorisation URL for a relevant payment provider.",
          },
          {
            "name": "code",
            "title": "Code",
            "type": "`$STRING`",
            "short": "Code parameter returned in redirect parameters",
          },
          {
            "name": "payment_details",
            "title": "Payment Details",
            "type": "`$OBJECT`",
            "req": True,
            "short": "Payment details required to make a Single Payment",
          },
          {
            "name": "payment_id",
            "title": "Payment Id",
            "type": "`$STRING`",
            "op": {
              "create": {
                "req": True,
                "type": "`$STRING`",
              },
            },
            "short": "Payment Identifier",
          },
          {
            "name": "payment_type",
            "title": "Payment Type",
            "type": "`$STRING`",
            "short": "The type of payment iniated through Bud's Payments Service",
          },
          {
            "name": "provider",
            "title": "Provider",
            "type": "`$STRING`",
            "req": True,
            "short": "The name (identifier) of the payment provider",
          },
          {
            "name": "provider_redirect_url",
            "title": "Provider Redirect Url",
            "type": "`$STRING`",
            "short": "This is the url that your Customer will be redirected to once they have authorised with the relevant provider.",
          },
          {
            "name": "provider_types",
            "title": "Provider Types",
            "type": "`$ARRAY`",
            "short": "One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow.",
          },
          {
            "name": "redirect_url",
            "title": "Redirect Url",
            "type": "`$STRING`",
            "short": "This is the url that your Customer will be redirected to when using the Bud user interface.",
          },
          {
            "name": "required_action",
            "title": "Required Action",
            "type": "`$STRING`",
            "short": "If present, this fild will indicate the next required action (if any) to be taken by the client in order to complete the payment flow | Value | Description | |:-------|:-------------| | confirm_scheduled_payment | Indicates that the client…",
          },
          {
            "name": "result",
            "title": "Result",
            "type": "`$STRING`",
          },
          {
            "name": "scheduled_payment_details",
            "title": "Scheduled Payment Details",
            "type": "`$OBJECT`",
            "req": True,
            "short": "Details of the scheduled payment",
          },
          {
            "name": "standing_order_details",
            "title": "Standing Order Details",
            "type": "`$OBJECT`",
            "req": True,
            "short": "Details of the standing order",
          },
          {
            "name": "state",
            "title": "State",
            "type": "`$STRING`",
            "req": True,
            "short": "The task ID",
          },
        ],
        "name": "initiate_payment_client_license",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "kind": "http",
                "method": "POST",
                "orig": "/v1/payments/authorisation-codes",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "authorisation-codes",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "authorisation-codes",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "POST",
                "orig": "/v1/payments/scheduled",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "scheduled",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "scheduled",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "POST",
                "orig": "/v1/payments/single",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "single",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "single",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "POST",
                "orig": "/v1/payments/standing-order",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "standing-order",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "standing-order",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_client_id",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "manage_payment": {
        "fields": [
          {
            "name": "amount",
            "title": "Amount",
            "type": "`$OBJECT`",
            "req": True,
            "short": "The monetary amount.",
          },
          {
            "name": "client_id",
            "title": "Client Id",
            "type": "`$STRING`",
            "req": True,
            "short": "Bud's Client/App Identifier",
          },
          {
            "name": "created_at",
            "title": "Created At",
            "type": "`$STRING`",
            "req": True,
            "short": "The created at timestamp for the payment",
            "format": "date-time",
          },
          {
            "name": "data",
            "title": "Data",
            "type": "`$OBJECT`",
            "req": True,
            "short": "Data structure for the Scheduled Payment object",
          },
          {
            "name": "errors",
            "title": "Errors",
            "type": "`$ARRAY`",
            "req": True,
            "short": "A list of any errors that may be associated with the payment",
          },
          {
            "name": "first_payment_date",
            "title": "First Payment Date",
            "type": "`$STRING`",
            "req": True,
            "short": "Date-time (ISO 8601) the first payment should be made.",
            "format": "date-time",
          },
          {
            "name": "frequency",
            "title": "Frequency",
            "type": "`$STRING`",
            "req": True,
            "short": "The payment frequency for a Standing Order",
          },
          {
            "name": "known_charges",
            "title": "Known Charges",
            "type": "`$OBJECT`",
            "req": True,
            "short": "A list detailing information surrounding any known charges that may be associated with the payment",
          },
          {
            "name": "last_payment_date",
            "title": "Last Payment Date",
            "type": "`$STRING`",
            "short": "Date-time (ISO 8601) after which payments should stop.",
            "format": "date-time",
          },
          {
            "name": "metadata",
            "title": "Metadata",
            "type": "`$OBJECT`",
          },
          {
            "name": "method",
            "title": "Method",
            "type": "`$STRING`",
            "req": True,
            "short": "The payment method",
          },
          {
            "name": "operation_id",
            "title": "Operation Id",
            "type": "`$STRING`",
            "req": True,
            "short": "A unique identifier/reference associated with a given endpoint/operation",
          },
          {
            "name": "organisation_id",
            "title": "Organisation Id",
            "type": "`$STRING`",
            "req": True,
            "short": "Bud's Organisation Identifier",
          },
          {
            "name": "payment_id",
            "title": "Payment Id",
            "type": "`$STRING`",
            "req": True,
            "short": "Payment Identifier",
          },
          {
            "name": "recipient",
            "title": "Recipient",
            "type": "`$OBJECT`",
            "req": True,
            "short": "Details of the recipient of the payment",
          },
          {
            "name": "recurring_amount",
            "title": "Recurring Amount",
            "type": "`$OBJECT`",
            "req": True,
            "short": "The monetary amount.",
          },
          {
            "name": "reference",
            "title": "Reference",
            "type": "`$STRING`",
            "req": True,
            "short": "Reference to use for this payment",
          },
          {
            "name": "requested_execution_date",
            "title": "Requested Execution Date",
            "type": "`$STRING`",
            "req": True,
            "short": "Date-time (ISO 8601) the Scheduled Payment should be made.",
            "format": "date-time",
          },
          {
            "name": "required_action",
            "title": "Required Action",
            "type": "`$STRING`",
            "req": True,
            "op": {
              "list": {
                "type": "`$STRING`",
              },
            },
            "short": "This field will indicate that there is an action required to be taken by the client in order to complete the payment.",
          },
          {
            "name": "sender",
            "title": "Sender",
            "type": "`$OBJECT`",
            "req": True,
            "short": "Details of the payment sender",
          },
          {
            "name": "state",
            "title": "State",
            "type": "`$OBJECT`",
            "req": True,
            "short": "Details of the payment state and history",
          },
          {
            "name": "supplementary_status",
            "title": "Supplementary Status",
            "type": "`$OBJECT`",
            "req": True,
            "short": "A list of supplementary states associated with the payment as described by the specific provider",
          },
        ],
        "name": "manage_payment",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "kind": "http",
                "method": "POST",
                "orig": "/v1/payments/scheduled/{payment_id}/confirm",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "scheduled",
                  },
                  {
                    "var": "scheduled_id",
                  },
                  {
                    "lit": "confirm",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "scheduled",
                  "{scheduled_id}",
                  "confirm",
                ],
                "rename": {
                  "param": {
                    "payment_id": "scheduled_id",
                  },
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                  "params": [
                    {
                      "name": "scheduled_id",
                      "orig": "payment_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "confirm",
                  "exist": [
                    "scheduled_id",
                    "x_client_id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "POST",
                "orig": "/v1/payments/single/{payment_id}/confirm",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "single",
                  },
                  {
                    "var": "single_id",
                  },
                  {
                    "lit": "confirm",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "single",
                  "{single_id}",
                  "confirm",
                ],
                "rename": {
                  "param": {
                    "payment_id": "single_id",
                  },
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                  "params": [
                    {
                      "name": "single_id",
                      "orig": "payment_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "confirm",
                  "exist": [
                    "single_id",
                    "x_client_id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "POST",
                "orig": "/v1/payments/standing-order/{payment_id}/confirm",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "standing-order",
                  },
                  {
                    "var": "standing_order_id",
                  },
                  {
                    "lit": "confirm",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "standing-order",
                  "{standing_order_id}",
                  "confirm",
                ],
                "rename": {
                  "param": {
                    "payment_id": "standing_order_id",
                  },
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                  "params": [
                    {
                      "name": "standing_order_id",
                      "orig": "payment_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "confirm",
                  "exist": [
                    "standing_order_id",
                    "x_client_id",
                  ],
                },
              },
            ],
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v1/payments/scheduled",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "scheduled",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "scheduled",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_amount_from",
                      "orig": "x_amount_from",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "10.00",
                    },
                    {
                      "name": "x_amount_to",
                      "orig": "x_amount_to",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "10000.00",
                    },
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_date_from",
                      "orig": "x_date_from",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "2020-06-01T15:00:00Z",
                    },
                    {
                      "name": "x_date_to",
                      "orig": "x_date_to",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "2020-06-01T15:00:00Z",
                    },
                    {
                      "name": "x_page_number",
                      "orig": "x_page_number",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "1",
                    },
                    {
                      "name": "x_payment_id",
                      "orig": "x_payment_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "84470aa9-71a2-4da2-8fa4-660adbf2c9b2",
                    },
                    {
                      "name": "x_sender_provider",
                      "orig": "x_sender_provider",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "Lloyds,Nationwide",
                    },
                    {
                      "name": "x_status_code",
                      "orig": "x_status_code",
                      "type": "`$ARRAY`",
                      "kind": "header",
                      "example": [
                        "CREATED",
                        "SETTLED",
                      ],
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_amount_from",
                    "x_amount_to",
                    "x_client_id",
                    "x_date_from",
                    "x_date_to",
                    "x_page_number",
                    "x_payment_id",
                    "x_sender_provider",
                    "x_status_code",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v1/payments/single",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "single",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "single",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_amount_from",
                      "orig": "x_amount_from",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "10.00",
                    },
                    {
                      "name": "x_amount_to",
                      "orig": "x_amount_to",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "10000.00",
                    },
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_date_from",
                      "orig": "x_date_from",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "2020-06-01T15:00:00Z",
                    },
                    {
                      "name": "x_date_to",
                      "orig": "x_date_to",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "2020-06-01T15:00:00Z",
                    },
                    {
                      "name": "x_page_number",
                      "orig": "x_page_number",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "1",
                    },
                    {
                      "name": "x_payment_id",
                      "orig": "x_payment_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "84470aa9-71a2-4da2-8fa4-660adbf2c9b2",
                    },
                    {
                      "name": "x_sender_provider",
                      "orig": "x_sender_provider",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "Lloyds,Nationwide",
                    },
                    {
                      "name": "x_status_code",
                      "orig": "x_status_code",
                      "type": "`$ARRAY`",
                      "kind": "header",
                      "example": [
                        "CREATED",
                        "SETTLED",
                      ],
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_amount_from",
                    "x_amount_to",
                    "x_client_id",
                    "x_date_from",
                    "x_date_to",
                    "x_page_number",
                    "x_payment_id",
                    "x_sender_provider",
                    "x_status_code",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v1/payments/standing-order",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "standing-order",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "standing-order",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_amount_from",
                      "orig": "x_amount_from",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "10.00",
                    },
                    {
                      "name": "x_amount_to",
                      "orig": "x_amount_to",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "10000.00",
                    },
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                    {
                      "name": "x_date_from",
                      "orig": "x_date_from",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "2020-06-01T15:00:00Z",
                    },
                    {
                      "name": "x_date_to",
                      "orig": "x_date_to",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "2020-06-01T15:00:00Z",
                    },
                    {
                      "name": "x_page_number",
                      "orig": "x_page_number",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "1",
                    },
                    {
                      "name": "x_payment_id",
                      "orig": "x_payment_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "84470aa9-71a2-4da2-8fa4-660adbf2c9b2",
                    },
                    {
                      "name": "x_sender_provider",
                      "orig": "x_sender_provider",
                      "type": "`$STRING`",
                      "kind": "header",
                      "example": "Lloyds,Nationwide",
                    },
                    {
                      "name": "x_status_code",
                      "orig": "x_status_code",
                      "type": "`$ARRAY`",
                      "kind": "header",
                      "example": [
                        "CREATED",
                        "SETTLED",
                      ],
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "x_amount_from",
                    "x_amount_to",
                    "x_client_id",
                    "x_date_from",
                    "x_date_to",
                    "x_page_number",
                    "x_payment_id",
                    "x_sender_provider",
                    "x_status_code",
                  ],
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v1/payments/scheduled/{payment_id}",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "scheduled",
                  },
                  {
                    "var": "payment_id",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "scheduled",
                  "{payment_id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                  "params": [
                    {
                      "name": "payment_id",
                      "orig": "payment_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "payment_id",
                    "x_client_id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v1/payments/single/{payment_id}",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "single",
                  },
                  {
                    "var": "payment_id",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "single",
                  "{payment_id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                  "params": [
                    {
                      "name": "payment_id",
                      "orig": "payment_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "payment_id",
                    "x_client_id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v1/payments/standing-order/{payment_id}",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "payments",
                  },
                  {
                    "lit": "standing-order",
                  },
                  {
                    "var": "payment_id",
                  },
                ],
                "parts": [
                  "v1",
                  "payments",
                  "standing-order",
                  "{payment_id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.data`",
                },
                "args": {
                  "header": [
                    {
                      "name": "x_client_id",
                      "orig": "x_client_id",
                      "type": "`$STRING`",
                      "kind": "header",
                      "reqd": True,
                      "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
                    },
                  ],
                  "params": [
                    {
                      "name": "payment_id",
                      "orig": "payment_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "payment_id",
                    "x_client_id",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
