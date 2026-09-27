// Typed models for the BudPayments SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/bud-payments-sdk/go/core"
)

// InitiatePaymentBudLicense is the typed data model for the initiate_payment_bud_license entity.
type InitiatePaymentBudLicense struct {
}

// InitiatePaymentBudLicenseListMatch is the typed request payload for InitiatePaymentBudLicense.ListTyped.
type InitiatePaymentBudLicenseListMatch struct {
	Country *string `json:"country,omitempty"`
	Product *[]any `json:"product,omitempty"`
	Type *string `json:"type,omitempty"`
}

// InitiatePaymentBudLicenseCreateData is the typed request payload for InitiatePaymentBudLicense.CreateTyped.
type InitiatePaymentBudLicenseCreateData struct {
	BudPayUrl string `json:"bud_pay_url"`
	CountryCode string `json:"country_code"`
	DisplayName string `json:"display_name"`
	Icon string `json:"icon"`
	ImplementationType string `json:"implementation_type"`
	MaintenanceStatus string `json:"maintenance_status"`
	MaintenanceWindow any `json:"maintenance_window"`
	PaymentDetails map[string]any `json:"payment_details"`
	PaymentId string `json:"payment_id"`
	Provider string `json:"provider"`
	ProviderTypes *[]any `json:"provider_types,omitempty"`
	RedirectUrl string `json:"redirect_url"`
	RequiredActions any `json:"required_actions"`
	ScheduledPaymentDetails map[string]any `json:"scheduled_payment_details"`
	Services []any `json:"services"`
	StandingOrderDetails map[string]any `json:"standing_order_details"`
	SupportedCurrencies []any `json:"supported_currencies"`
}

// InitiatePaymentClientLicense is the typed data model for the initiate_payment_client_license entity.
type InitiatePaymentClientLicense struct {
}

// InitiatePaymentClientLicenseCreateData is the typed request payload for InitiatePaymentClientLicense.CreateTyped.
type InitiatePaymentClientLicenseCreateData struct {
	AuthorisationUrl string `json:"authorisation_url"`
	Code *string `json:"code,omitempty"`
	PaymentDetails map[string]any `json:"payment_details"`
	PaymentId *string `json:"payment_id,omitempty"`
	PaymentType *string `json:"payment_type,omitempty"`
	Provider string `json:"provider"`
	ProviderRedirectUrl *string `json:"provider_redirect_url,omitempty"`
	ProviderTypes *[]any `json:"provider_types,omitempty"`
	RedirectUrl *string `json:"redirect_url,omitempty"`
	RequiredAction *string `json:"required_action,omitempty"`
	Result *string `json:"result,omitempty"`
	ScheduledPaymentDetails map[string]any `json:"scheduled_payment_details"`
	StandingOrderDetails map[string]any `json:"standing_order_details"`
	State string `json:"state"`
}

// ManagePayment is the typed data model for the manage_payment entity.
type ManagePayment struct {
}

// ManagePaymentLoadMatch is the typed request payload for ManagePayment.LoadTyped.
type ManagePaymentLoadMatch struct {
	PaymentId string `json:"payment_id"`
}

// ManagePaymentListMatch is the typed request payload for ManagePayment.ListTyped.
type ManagePaymentListMatch struct {
	Amount *map[string]any `json:"amount,omitempty"`
	ClientId *string `json:"client_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Data *map[string]any `json:"data,omitempty"`
	Errors *[]any `json:"errors,omitempty"`
	FirstPaymentDate *string `json:"first_payment_date,omitempty"`
	Frequency *string `json:"frequency,omitempty"`
	KnownCharges *map[string]any `json:"known_charges,omitempty"`
	LastPaymentDate *string `json:"last_payment_date,omitempty"`
	Metadata *map[string]any `json:"metadata,omitempty"`
	Method *string `json:"method,omitempty"`
	OperationId *string `json:"operation_id,omitempty"`
	OrganisationId *string `json:"organisation_id,omitempty"`
	PaymentId *string `json:"payment_id,omitempty"`
	Recipient *map[string]any `json:"recipient,omitempty"`
	RecurringAmount *map[string]any `json:"recurring_amount,omitempty"`
	Reference *string `json:"reference,omitempty"`
	RequestedExecutionDate *string `json:"requested_execution_date,omitempty"`
	RequiredAction *string `json:"required_action,omitempty"`
	Sender *map[string]any `json:"sender,omitempty"`
	State *map[string]any `json:"state,omitempty"`
	SupplementaryStatus *map[string]any `json:"supplementary_status,omitempty"`
}

// ManagePaymentCreateData is the typed request payload for ManagePayment.CreateTyped.
type ManagePaymentCreateData struct {
	ScheduledId string `json:"scheduled_id"`
	Amount map[string]any `json:"amount"`
	ClientId string `json:"client_id"`
	CreatedAt string `json:"created_at"`
	Data map[string]any `json:"data"`
	Errors []any `json:"errors"`
	FirstPaymentDate string `json:"first_payment_date"`
	Frequency string `json:"frequency"`
	KnownCharges map[string]any `json:"known_charges"`
	LastPaymentDate *string `json:"last_payment_date,omitempty"`
	Metadata *map[string]any `json:"metadata,omitempty"`
	Method string `json:"method"`
	OperationId string `json:"operation_id"`
	OrganisationId string `json:"organisation_id"`
	PaymentId string `json:"payment_id"`
	Recipient map[string]any `json:"recipient"`
	RecurringAmount map[string]any `json:"recurring_amount"`
	Reference string `json:"reference"`
	RequestedExecutionDate string `json:"requested_execution_date"`
	RequiredAction string `json:"required_action"`
	Sender map[string]any `json:"sender"`
	State map[string]any `json:"state"`
	SupplementaryStatus map[string]any `json:"supplementary_status"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
