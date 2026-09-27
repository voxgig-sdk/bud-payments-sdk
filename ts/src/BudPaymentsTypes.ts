// Typed models for the BudPayments SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface InitiatePaymentBudLicense {
  bud_pay_url: string
  country_code: string
  display_name: string
  icon: string
  implementation_type: string
  maintenance_status: string
  maintenance_window: any
  payment_details: Record<string, any>
  payment_id: string
  provider: string
  provider_types?: any[]
  redirect_url: string
  required_actions: any
  scheduled_payment_details: Record<string, any>
  services: any[]
  standing_order_details: Record<string, any>
  supported_currencies: any[]
}

export interface InitiatePaymentBudLicenseListMatch {
  country?: string
  product?: any[]
  type?: string
}

export interface InitiatePaymentBudLicenseCreateData {
  bud_pay_url: string
  country_code: string
  display_name: string
  icon: string
  implementation_type: string
  maintenance_status: string
  maintenance_window: any
  payment_details: Record<string, any>
  payment_id: string
  provider: string
  provider_types?: any[]
  redirect_url: string
  required_actions: any
  scheduled_payment_details: Record<string, any>
  services: any[]
  standing_order_details: Record<string, any>
  supported_currencies: any[]
}

export interface InitiatePaymentClientLicense {
  authorisation_url: string
  code?: string
  payment_details: Record<string, any>
  payment_id?: string
  payment_type?: string
  provider: string
  provider_redirect_url?: string
  provider_types?: any[]
  redirect_url?: string
  required_action?: string
  result?: string
  scheduled_payment_details: Record<string, any>
  standing_order_details: Record<string, any>
  state: string
}

export interface InitiatePaymentClientLicenseCreateData {
  authorisation_url: string
  code?: string
  payment_details: Record<string, any>
  payment_id?: string
  payment_type?: string
  provider: string
  provider_redirect_url?: string
  provider_types?: any[]
  redirect_url?: string
  required_action?: string
  result?: string
  scheduled_payment_details: Record<string, any>
  standing_order_details: Record<string, any>
  state: string
}

export interface ManagePayment {
  amount: Record<string, any>
  client_id: string
  created_at: string
  data: Record<string, any>
  errors: any[]
  first_payment_date: string
  frequency: string
  known_charges: Record<string, any>
  last_payment_date?: string
  metadata?: Record<string, any>
  method: string
  operation_id: string
  organisation_id: string
  payment_id: string
  recipient: Record<string, any>
  recurring_amount: Record<string, any>
  reference: string
  requested_execution_date: string
  required_action: string
  sender: Record<string, any>
  state: Record<string, any>
  supplementary_status: Record<string, any>
}

export interface ManagePaymentLoadMatch {
  payment_id: string
}

export interface ManagePaymentListMatch {
  amount?: Record<string, any>
  client_id?: string
  created_at?: string
  data?: Record<string, any>
  errors?: any[]
  first_payment_date?: string
  frequency?: string
  known_charges?: Record<string, any>
  last_payment_date?: string
  metadata?: Record<string, any>
  method?: string
  operation_id?: string
  organisation_id?: string
  payment_id?: string
  recipient?: Record<string, any>
  recurring_amount?: Record<string, any>
  reference?: string
  requested_execution_date?: string
  required_action?: string
  sender?: Record<string, any>
  state?: Record<string, any>
  supplementary_status?: Record<string, any>
}

export interface ManagePaymentCreateData {
  scheduled_id: string
  amount: Record<string, any>
  client_id: string
  created_at: string
  data: Record<string, any>
  errors: any[]
  first_payment_date: string
  frequency: string
  known_charges: Record<string, any>
  last_payment_date?: string
  metadata?: Record<string, any>
  method: string
  operation_id: string
  organisation_id: string
  payment_id: string
  recipient: Record<string, any>
  recurring_amount: Record<string, any>
  reference: string
  requested_execution_date: string
  required_action: string
  sender: Record<string, any>
  state: Record<string, any>
  supplementary_status: Record<string, any>

  // Selects a custom action instead of the plain create:
  //   'confirm' | 'confirm' | 'confirm'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

