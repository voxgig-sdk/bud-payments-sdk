package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "BudPayments",
			"slug": "bud-payments",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"debug": map[string]any{
				"options": map[string]any{
					"active": false,
					"max": 100,
					"redact": []any{
						"authorization",
						"cookie",
						"set-cookie",
						"api-key",
						"apikey",
						"x-api-key",
						"idempotency-key",
					},
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"onEntry": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"idempotency": map[string]any{
				"options": map[string]any{
					"active": false,
					"header": "Idempotency-Key",
					"methods": []any{
						"POST",
						"PUT",
						"PATCH",
						"DELETE",
					},
					"ops": []any{
						"create",
						"update",
						"remove",
					},
				},
				"optspec": map[string]any{
					"keygen": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"metrics": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"paging": map[string]any{
				"options": map[string]any{
					"active": false,
					"afterVar": "after",
					"cursorParam": "cursor",
					"firstVar": "first",
					"limitParam": "limit",
					"pageParam": "page",
					"startPage": 1,
				},
				"optspec": map[string]any{
					"limit": "`$NUMBER`",
					"ops": "`$LIST`",
				},
				"strict": false,
				"transport": "none",
			},
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://api-sandbox.thisisbud.com",
			"auth": map[string]any{
				"prefix": "Bearer",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"initiate_payment_bud_license": map[string]any{},
				"initiate_payment_client_license": map[string]any{},
				"manage_payment": map[string]any{},
			},
		},
		"entity": map[string]any{
			"initiate_payment_bud_license": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "bud_pay_url",
						"title": "Bud Pay Url",
						"type": "`$STRING`",
						"req": true,
						"short": "A Bud Pay URL, allowing Customers to complete (authenticate) a relevant payment resource with the relevant payment service provider.",
					},
					map[string]any{
						"name": "country_code",
						"title": "Country Code",
						"type": "`$STRING`",
						"req": true,
						"short": "Provider country code following the ISO 3166 alpha-3 code format",
					},
					map[string]any{
						"name": "display_name",
						"title": "Display Name",
						"type": "`$STRING`",
						"req": true,
						"short": "The name of the payment provider to be used for display purposes (e.g.",
					},
					map[string]any{
						"name": "icon",
						"title": "Icon",
						"type": "`$STRING`",
						"req": true,
						"short": "A URL for the icon of the corresponding provider",
					},
					map[string]any{
						"name": "implementation_type",
						"title": "Implementation Type",
						"type": "`$STRING`",
						"req": true,
						"short": "The implementation standard adopted by the provider",
					},
					map[string]any{
						"name": "maintenance_status",
						"title": "Maintenance Status",
						"type": "`$STRING`",
						"req": true,
						"short": "The status of the maintenance window associated with the provider",
					},
					map[string]any{
						"name": "maintenance_window",
						"title": "Maintenance Window",
						"type": "`$ANY`",
						"req": true,
					},
					map[string]any{
						"name": "payment_details",
						"title": "Payment Details",
						"type": "`$OBJECT`",
						"req": true,
						"short": "Payment details required to make a Single Payment",
					},
					map[string]any{
						"name": "payment_id",
						"title": "Payment Id",
						"type": "`$STRING`",
						"req": true,
						"short": "Unique identifier for the payment",
					},
					map[string]any{
						"name": "provider",
						"title": "Provider",
						"type": "`$STRING`",
						"req": true,
						"short": "The name (identifier) of the payment provider",
					},
					map[string]any{
						"name": "provider_types",
						"title": "Provider Types",
						"type": "`$ARRAY`",
						"short": "One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow.",
					},
					map[string]any{
						"name": "redirect_url",
						"title": "Redirect Url",
						"type": "`$STRING`",
						"req": true,
						"short": "URL where the user will be redirected to once they have completed the TPP Connection flow.",
					},
					map[string]any{
						"name": "required_actions",
						"title": "Required Actions",
						"type": "`$ANY`",
						"req": true,
						"short": "An array of enum fields to indicate if any additional actions are required in order to complete a given payment flow for the given provider.",
					},
					map[string]any{
						"name": "scheduled_payment_details",
						"title": "Scheduled Payment Details",
						"type": "`$OBJECT`",
						"req": true,
						"short": "Details of the scheduled payment",
					},
					map[string]any{
						"name": "services",
						"title": "Services",
						"type": "`$ARRAY`",
						"req": true,
						"short": "A payment service supported by a payment provider",
					},
					map[string]any{
						"name": "standing_order_details",
						"title": "Standing Order Details",
						"type": "`$OBJECT`",
						"req": true,
						"short": "Details of the standing order",
					},
					map[string]any{
						"name": "supported_currencies",
						"title": "Supported Currencies",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The currencies supported by the payment provider",
					},
				},
				"name": "initiate_payment_bud_license",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/payments/scheduled/bud-pay-url",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "scheduled",
									},
									map[string]any{
										"lit": "bud-pay-url",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"scheduled",
									"bud-pay-url",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/payments/single/bud-pay-url",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "single",
									},
									map[string]any{
										"lit": "bud-pay-url",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"single",
									"bud-pay-url",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/payments/standing-order/bud-pay-url",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "standing-order",
									},
									map[string]any{
										"lit": "bud-pay-url",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"standing-order",
									"bud-pay-url",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v1/payments/providers",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "providers",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"providers",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"query": []any{
										map[string]any{
											"name": "country",
											"orig": "country",
											"type": "`$STRING`",
											"kind": "query",
											"example": "GBR,DEU",
										},
										map[string]any{
											"name": "product",
											"orig": "product",
											"type": "`$ARRAY`",
											"kind": "query",
										},
										map[string]any{
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"country",
										"product",
										"type",
										"x_client_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"initiate_payment_client_license": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "authorisation_url",
						"title": "Authorisation Url",
						"type": "`$STRING`",
						"req": true,
						"short": "The authorisation URL for a relevant payment provider.",
					},
					map[string]any{
						"name": "code",
						"title": "Code",
						"type": "`$STRING`",
						"short": "Code parameter returned in redirect parameters",
					},
					map[string]any{
						"name": "payment_details",
						"title": "Payment Details",
						"type": "`$OBJECT`",
						"req": true,
						"short": "Payment details required to make a Single Payment",
					},
					map[string]any{
						"name": "payment_id",
						"title": "Payment Id",
						"type": "`$STRING`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "Payment Identifier",
					},
					map[string]any{
						"name": "payment_type",
						"title": "Payment Type",
						"type": "`$STRING`",
						"short": "The type of payment iniated through Bud's Payments Service",
					},
					map[string]any{
						"name": "provider",
						"title": "Provider",
						"type": "`$STRING`",
						"req": true,
						"short": "The name (identifier) of the payment provider",
					},
					map[string]any{
						"name": "provider_redirect_url",
						"title": "Provider Redirect Url",
						"type": "`$STRING`",
						"short": "This is the url that your Customer will be redirected to once they have authorised with the relevant provider.",
					},
					map[string]any{
						"name": "provider_types",
						"title": "Provider Types",
						"type": "`$ARRAY`",
						"short": "One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow.",
					},
					map[string]any{
						"name": "redirect_url",
						"title": "Redirect Url",
						"type": "`$STRING`",
						"short": "This is the url that your Customer will be redirected to when using the Bud user interface.",
					},
					map[string]any{
						"name": "required_action",
						"title": "Required Action",
						"type": "`$STRING`",
						"short": "If present, this fild will indicate the next required action (if any) to be taken by the client in order to complete the payment flow | Value | Description | |:-------|:-------------| | confirm_scheduled_payment | Indicates that the client…",
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "scheduled_payment_details",
						"title": "Scheduled Payment Details",
						"type": "`$OBJECT`",
						"req": true,
						"short": "Details of the scheduled payment",
					},
					map[string]any{
						"name": "standing_order_details",
						"title": "Standing Order Details",
						"type": "`$OBJECT`",
						"req": true,
						"short": "Details of the standing order",
					},
					map[string]any{
						"name": "state",
						"title": "State",
						"type": "`$STRING`",
						"req": true,
						"short": "The task ID",
					},
				},
				"name": "initiate_payment_client_license",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/payments/authorisation-codes",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "authorisation-codes",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"authorisation-codes",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/payments/scheduled",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "scheduled",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"scheduled",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/payments/single",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "single",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"single",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/payments/standing-order",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "standing-order",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"standing-order",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"manage_payment": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "amount",
						"title": "Amount",
						"type": "`$OBJECT`",
						"req": true,
						"short": "The monetary amount.",
					},
					map[string]any{
						"name": "client_id",
						"title": "Client Id",
						"type": "`$STRING`",
						"req": true,
						"short": "Bud's Client/App Identifier",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"short": "The created at timestamp for the payment",
						"format": "date-time",
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
						"short": "Data structure for the Scheduled Payment object",
					},
					map[string]any{
						"name": "errors",
						"title": "Errors",
						"type": "`$ARRAY`",
						"req": true,
						"short": "A list of any errors that may be associated with the payment",
					},
					map[string]any{
						"name": "first_payment_date",
						"title": "First Payment Date",
						"type": "`$STRING`",
						"req": true,
						"short": "Date-time (ISO 8601) the first payment should be made.",
						"format": "date-time",
					},
					map[string]any{
						"name": "frequency",
						"title": "Frequency",
						"type": "`$STRING`",
						"req": true,
						"short": "The payment frequency for a Standing Order",
					},
					map[string]any{
						"name": "known_charges",
						"title": "Known Charges",
						"type": "`$OBJECT`",
						"req": true,
						"short": "A list detailing information surrounding any known charges that may be associated with the payment",
					},
					map[string]any{
						"name": "last_payment_date",
						"title": "Last Payment Date",
						"type": "`$STRING`",
						"short": "Date-time (ISO 8601) after which payments should stop.",
						"format": "date-time",
					},
					map[string]any{
						"name": "metadata",
						"title": "Metadata",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "method",
						"title": "Method",
						"type": "`$STRING`",
						"req": true,
						"short": "The payment method",
					},
					map[string]any{
						"name": "operation_id",
						"title": "Operation Id",
						"type": "`$STRING`",
						"req": true,
						"short": "A unique identifier/reference associated with a given endpoint/operation",
					},
					map[string]any{
						"name": "organisation_id",
						"title": "Organisation Id",
						"type": "`$STRING`",
						"req": true,
						"short": "Bud's Organisation Identifier",
					},
					map[string]any{
						"name": "payment_id",
						"title": "Payment Id",
						"type": "`$STRING`",
						"req": true,
						"short": "Payment Identifier",
					},
					map[string]any{
						"name": "recipient",
						"title": "Recipient",
						"type": "`$OBJECT`",
						"req": true,
						"short": "Details of the recipient of the payment",
					},
					map[string]any{
						"name": "recurring_amount",
						"title": "Recurring Amount",
						"type": "`$OBJECT`",
						"req": true,
						"short": "The monetary amount.",
					},
					map[string]any{
						"name": "reference",
						"title": "Reference",
						"type": "`$STRING`",
						"req": true,
						"short": "Reference to use for this payment",
					},
					map[string]any{
						"name": "requested_execution_date",
						"title": "Requested Execution Date",
						"type": "`$STRING`",
						"req": true,
						"short": "Date-time (ISO 8601) the Scheduled Payment should be made.",
						"format": "date-time",
					},
					map[string]any{
						"name": "required_action",
						"title": "Required Action",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"list": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "This field will indicate that there is an action required to be taken by the client in order to complete the payment.",
					},
					map[string]any{
						"name": "sender",
						"title": "Sender",
						"type": "`$OBJECT`",
						"req": true,
						"short": "Details of the payment sender",
					},
					map[string]any{
						"name": "state",
						"title": "State",
						"type": "`$OBJECT`",
						"req": true,
						"short": "Details of the payment state and history",
					},
					map[string]any{
						"name": "supplementary_status",
						"title": "Supplementary Status",
						"type": "`$OBJECT`",
						"req": true,
						"short": "A list of supplementary states associated with the payment as described by the specific provider",
					},
				},
				"name": "manage_payment",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/payments/scheduled/{payment_id}/confirm",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "scheduled",
									},
									map[string]any{
										"var": "scheduled_id",
									},
									map[string]any{
										"lit": "confirm",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"scheduled",
									"{scheduled_id}",
									"confirm",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"payment_id": "scheduled_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "scheduled_id",
											"orig": "payment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "confirm",
									"exist": []any{
										"scheduled_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/payments/single/{payment_id}/confirm",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "single",
									},
									map[string]any{
										"var": "single_id",
									},
									map[string]any{
										"lit": "confirm",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"single",
									"{single_id}",
									"confirm",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"payment_id": "single_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "single_id",
											"orig": "payment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "confirm",
									"exist": []any{
										"single_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/payments/standing-order/{payment_id}/confirm",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "standing-order",
									},
									map[string]any{
										"var": "standing_order_id",
									},
									map[string]any{
										"lit": "confirm",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"standing-order",
									"{standing_order_id}",
									"confirm",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"payment_id": "standing_order_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "standing_order_id",
											"orig": "payment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "confirm",
									"exist": []any{
										"standing_order_id",
										"x_client_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v1/payments/scheduled",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "scheduled",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"scheduled",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_amount_from",
											"orig": "x_amount_from",
											"type": "`$STRING`",
											"kind": "header",
											"example": "10.00",
										},
										map[string]any{
											"name": "x_amount_to",
											"orig": "x_amount_to",
											"type": "`$STRING`",
											"kind": "header",
											"example": "10000.00",
										},
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_date_from",
											"orig": "x_date_from",
											"type": "`$STRING`",
											"kind": "header",
											"example": "2020-06-01T15:00:00Z",
										},
										map[string]any{
											"name": "x_date_to",
											"orig": "x_date_to",
											"type": "`$STRING`",
											"kind": "header",
											"example": "2020-06-01T15:00:00Z",
										},
										map[string]any{
											"name": "x_page_number",
											"orig": "x_page_number",
											"type": "`$STRING`",
											"kind": "header",
											"example": "1",
										},
										map[string]any{
											"name": "x_payment_id",
											"orig": "x_payment_id",
											"type": "`$STRING`",
											"kind": "header",
											"example": "84470aa9-71a2-4da2-8fa4-660adbf2c9b2",
										},
										map[string]any{
											"name": "x_sender_provider",
											"orig": "x_sender_provider",
											"type": "`$STRING`",
											"kind": "header",
											"example": "Lloyds,Nationwide",
										},
										map[string]any{
											"name": "x_status_code",
											"orig": "x_status_code",
											"type": "`$ARRAY`",
											"kind": "header",
											"example": []any{
												"CREATED",
												"SETTLED",
											},
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_amount_from",
										"x_amount_to",
										"x_client_id",
										"x_date_from",
										"x_date_to",
										"x_page_number",
										"x_payment_id",
										"x_sender_provider",
										"x_status_code",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v1/payments/single",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "single",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"single",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_amount_from",
											"orig": "x_amount_from",
											"type": "`$STRING`",
											"kind": "header",
											"example": "10.00",
										},
										map[string]any{
											"name": "x_amount_to",
											"orig": "x_amount_to",
											"type": "`$STRING`",
											"kind": "header",
											"example": "10000.00",
										},
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_date_from",
											"orig": "x_date_from",
											"type": "`$STRING`",
											"kind": "header",
											"example": "2020-06-01T15:00:00Z",
										},
										map[string]any{
											"name": "x_date_to",
											"orig": "x_date_to",
											"type": "`$STRING`",
											"kind": "header",
											"example": "2020-06-01T15:00:00Z",
										},
										map[string]any{
											"name": "x_page_number",
											"orig": "x_page_number",
											"type": "`$STRING`",
											"kind": "header",
											"example": "1",
										},
										map[string]any{
											"name": "x_payment_id",
											"orig": "x_payment_id",
											"type": "`$STRING`",
											"kind": "header",
											"example": "84470aa9-71a2-4da2-8fa4-660adbf2c9b2",
										},
										map[string]any{
											"name": "x_sender_provider",
											"orig": "x_sender_provider",
											"type": "`$STRING`",
											"kind": "header",
											"example": "Lloyds,Nationwide",
										},
										map[string]any{
											"name": "x_status_code",
											"orig": "x_status_code",
											"type": "`$ARRAY`",
											"kind": "header",
											"example": []any{
												"CREATED",
												"SETTLED",
											},
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_amount_from",
										"x_amount_to",
										"x_client_id",
										"x_date_from",
										"x_date_to",
										"x_page_number",
										"x_payment_id",
										"x_sender_provider",
										"x_status_code",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v1/payments/standing-order",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "standing-order",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"standing-order",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_amount_from",
											"orig": "x_amount_from",
											"type": "`$STRING`",
											"kind": "header",
											"example": "10.00",
										},
										map[string]any{
											"name": "x_amount_to",
											"orig": "x_amount_to",
											"type": "`$STRING`",
											"kind": "header",
											"example": "10000.00",
										},
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_date_from",
											"orig": "x_date_from",
											"type": "`$STRING`",
											"kind": "header",
											"example": "2020-06-01T15:00:00Z",
										},
										map[string]any{
											"name": "x_date_to",
											"orig": "x_date_to",
											"type": "`$STRING`",
											"kind": "header",
											"example": "2020-06-01T15:00:00Z",
										},
										map[string]any{
											"name": "x_page_number",
											"orig": "x_page_number",
											"type": "`$STRING`",
											"kind": "header",
											"example": "1",
										},
										map[string]any{
											"name": "x_payment_id",
											"orig": "x_payment_id",
											"type": "`$STRING`",
											"kind": "header",
											"example": "84470aa9-71a2-4da2-8fa4-660adbf2c9b2",
										},
										map[string]any{
											"name": "x_sender_provider",
											"orig": "x_sender_provider",
											"type": "`$STRING`",
											"kind": "header",
											"example": "Lloyds,Nationwide",
										},
										map[string]any{
											"name": "x_status_code",
											"orig": "x_status_code",
											"type": "`$ARRAY`",
											"kind": "header",
											"example": []any{
												"CREATED",
												"SETTLED",
											},
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_amount_from",
										"x_amount_to",
										"x_client_id",
										"x_date_from",
										"x_date_to",
										"x_page_number",
										"x_payment_id",
										"x_sender_provider",
										"x_status_code",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v1/payments/scheduled/{payment_id}",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "scheduled",
									},
									map[string]any{
										"var": "payment_id",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"scheduled",
									"{payment_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "payment_id",
											"orig": "payment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"payment_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v1/payments/single/{payment_id}",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "single",
									},
									map[string]any{
										"var": "payment_id",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"single",
									"{payment_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "payment_id",
											"orig": "payment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"payment_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v1/payments/standing-order/{payment_id}",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "payments",
									},
									map[string]any{
										"lit": "standing-order",
									},
									map[string]any{
										"var": "payment_id",
									},
								},
								"parts": []any{
									"v1",
									"payments",
									"standing-order",
									"{payment_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "payment_id",
											"orig": "payment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"payment_id",
										"x_client_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "debug":
		if NewDebugFeatureFunc != nil {
			return NewDebugFeatureFunc()
		}
	case "idempotency":
		if NewIdempotencyFeatureFunc != nil {
			return NewIdempotencyFeatureFunc()
		}
	case "metrics":
		if NewMetricsFeatureFunc != nil {
			return NewMetricsFeatureFunc()
		}
	case "paging":
		if NewPagingFeatureFunc != nil {
			return NewPagingFeatureFunc()
		}
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
