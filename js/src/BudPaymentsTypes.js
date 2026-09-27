// Typed models for the BudPayments SDK (JSDoc typedefs).
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
// edit by hand.

/**
 * @typedef {Object} InitiatePaymentBudLicense
 * @property {string} bud_pay_url
 * @property {string} country_code
 * @property {string} display_name
 * @property {string} icon
 * @property {string} implementation_type
 * @property {string} maintenance_status
 * @property {*} maintenance_window
 * @property {Object} payment_details
 * @property {string} payment_id
 * @property {string} provider
 * @property {Array} [provider_types]
 * @property {string} redirect_url
 * @property {*} required_actions
 * @property {Object} scheduled_payment_details
 * @property {Array} services
 * @property {Object} standing_order_details
 * @property {Array} supported_currencies
 */

/**
 * @typedef {Object} InitiatePaymentBudLicenseListMatch
 * @property {string} [country]
 * @property {Array} [product]
 * @property {string} [type]
 */

/**
 * @typedef {Object} InitiatePaymentBudLicenseCreateData
 * @property {string} bud_pay_url
 * @property {string} country_code
 * @property {string} display_name
 * @property {string} icon
 * @property {string} implementation_type
 * @property {string} maintenance_status
 * @property {*} maintenance_window
 * @property {Object} payment_details
 * @property {string} payment_id
 * @property {string} provider
 * @property {Array} [provider_types]
 * @property {string} redirect_url
 * @property {*} required_actions
 * @property {Object} scheduled_payment_details
 * @property {Array} services
 * @property {Object} standing_order_details
 * @property {Array} supported_currencies
 */

/**
 * @typedef {Object} InitiatePaymentClientLicense
 * @property {string} authorisation_url
 * @property {string} [code]
 * @property {Object} payment_details
 * @property {string} [payment_id]
 * @property {string} [payment_type]
 * @property {string} provider
 * @property {string} [provider_redirect_url]
 * @property {Array} [provider_types]
 * @property {string} [redirect_url]
 * @property {string} [required_action]
 * @property {string} [result]
 * @property {Object} scheduled_payment_details
 * @property {Object} standing_order_details
 * @property {string} state
 */

/**
 * @typedef {Object} InitiatePaymentClientLicenseCreateData
 * @property {string} authorisation_url
 * @property {string} [code]
 * @property {Object} payment_details
 * @property {string} [payment_id]
 * @property {string} [payment_type]
 * @property {string} provider
 * @property {string} [provider_redirect_url]
 * @property {Array} [provider_types]
 * @property {string} [redirect_url]
 * @property {string} [required_action]
 * @property {string} [result]
 * @property {Object} scheduled_payment_details
 * @property {Object} standing_order_details
 * @property {string} state
 */

/**
 * @typedef {Object} ManagePayment
 * @property {Object} amount
 * @property {string} client_id
 * @property {string} created_at
 * @property {Object} data
 * @property {Array} errors
 * @property {string} first_payment_date
 * @property {string} frequency
 * @property {Object} known_charges
 * @property {string} [last_payment_date]
 * @property {Object} [metadata]
 * @property {string} method
 * @property {string} operation_id
 * @property {string} organisation_id
 * @property {string} payment_id
 * @property {Object} recipient
 * @property {Object} recurring_amount
 * @property {string} reference
 * @property {string} requested_execution_date
 * @property {string} required_action
 * @property {Object} sender
 * @property {Object} state
 * @property {Object} supplementary_status
 */

/**
 * @typedef {Object} ManagePaymentLoadMatch
 * @property {string} payment_id
 */

/**
 * @typedef {Object} ManagePaymentListMatch
 * @property {Object} [amount]
 * @property {string} [client_id]
 * @property {string} [created_at]
 * @property {Object} [data]
 * @property {Array} [errors]
 * @property {string} [first_payment_date]
 * @property {string} [frequency]
 * @property {Object} [known_charges]
 * @property {string} [last_payment_date]
 * @property {Object} [metadata]
 * @property {string} [method]
 * @property {string} [operation_id]
 * @property {string} [organisation_id]
 * @property {string} [payment_id]
 * @property {Object} [recipient]
 * @property {Object} [recurring_amount]
 * @property {string} [reference]
 * @property {string} [requested_execution_date]
 * @property {string} [required_action]
 * @property {Object} [sender]
 * @property {Object} [state]
 * @property {Object} [supplementary_status]
 */

/**
 * @typedef {Object} ManagePaymentCreateData
 * @property {string} scheduled_id
 * @property {Object} amount
 * @property {string} client_id
 * @property {string} created_at
 * @property {Object} data
 * @property {Array} errors
 * @property {string} first_payment_date
 * @property {string} frequency
 * @property {Object} known_charges
 * @property {string} [last_payment_date]
 * @property {Object} [metadata]
 * @property {string} method
 * @property {string} operation_id
 * @property {string} organisation_id
 * @property {string} payment_id
 * @property {Object} recipient
 * @property {Object} recurring_amount
 * @property {string} reference
 * @property {string} requested_execution_date
 * @property {string} required_action
 * @property {Object} sender
 * @property {Object} state
 * @property {Object} supplementary_status
 */

