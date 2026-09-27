<?php
declare(strict_types=1);

// Typed models for the BudPayments SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** InitiatePaymentBudLicense entity data model. */
class InitiatePaymentBudLicense
{
    public string $bud_pay_url;
    public string $country_code;
    public string $display_name;
    public string $icon;
    public string $implementation_type;
    public string $maintenance_status;
    public mixed $maintenance_window;
    public array $payment_details;
    public string $payment_id;
    public string $provider;
    public ?array $provider_types = null;
    public string $redirect_url;
    public mixed $required_actions;
    public array $scheduled_payment_details;
    public array $services;
    public array $standing_order_details;
    public array $supported_currencies;
}

/** Request payload for InitiatePaymentBudLicense#list. */
class InitiatePaymentBudLicenseListMatch
{
    public ?string $country = null;
    public ?array $product = null;
    public ?string $type = null;
}

/** Request payload for InitiatePaymentBudLicense#create. */
class InitiatePaymentBudLicenseCreateData
{
    public string $bud_pay_url;
    public string $country_code;
    public string $display_name;
    public string $icon;
    public string $implementation_type;
    public string $maintenance_status;
    public mixed $maintenance_window;
    public array $payment_details;
    public string $payment_id;
    public string $provider;
    public ?array $provider_types = null;
    public string $redirect_url;
    public mixed $required_actions;
    public array $scheduled_payment_details;
    public array $services;
    public array $standing_order_details;
    public array $supported_currencies;
}

/** InitiatePaymentClientLicense entity data model. */
class InitiatePaymentClientLicense
{
    public string $authorisation_url;
    public ?string $code = null;
    public array $payment_details;
    public ?string $payment_id = null;
    public ?string $payment_type = null;
    public string $provider;
    public ?string $provider_redirect_url = null;
    public ?array $provider_types = null;
    public ?string $redirect_url = null;
    public ?string $required_action = null;
    public ?string $result = null;
    public array $scheduled_payment_details;
    public array $standing_order_details;
    public string $state;
}

/** Request payload for InitiatePaymentClientLicense#create. */
class InitiatePaymentClientLicenseCreateData
{
    public string $authorisation_url;
    public ?string $code = null;
    public array $payment_details;
    public ?string $payment_id = null;
    public ?string $payment_type = null;
    public string $provider;
    public ?string $provider_redirect_url = null;
    public ?array $provider_types = null;
    public ?string $redirect_url = null;
    public ?string $required_action = null;
    public ?string $result = null;
    public array $scheduled_payment_details;
    public array $standing_order_details;
    public string $state;
}

/** ManagePayment entity data model. */
class ManagePayment
{
    public array $amount;
    public string $client_id;
    public string $created_at;
    public array $data;
    public array $errors;
    public string $first_payment_date;
    public string $frequency;
    public array $known_charges;
    public ?string $last_payment_date = null;
    public ?array $metadata = null;
    public string $method;
    public string $operation_id;
    public string $organisation_id;
    public string $payment_id;
    public array $recipient;
    public array $recurring_amount;
    public string $reference;
    public string $requested_execution_date;
    public string $required_action;
    public array $sender;
    public array $state;
    public array $supplementary_status;
}

/** Request payload for ManagePayment#load. */
class ManagePaymentLoadMatch
{
    public string $payment_id;
}

/** Request payload for ManagePayment#list. */
class ManagePaymentListMatch
{
    public ?array $amount = null;
    public ?string $client_id = null;
    public ?string $created_at = null;
    public ?array $data = null;
    public ?array $errors = null;
    public ?string $first_payment_date = null;
    public ?string $frequency = null;
    public ?array $known_charges = null;
    public ?string $last_payment_date = null;
    public ?array $metadata = null;
    public ?string $method = null;
    public ?string $operation_id = null;
    public ?string $organisation_id = null;
    public ?string $payment_id = null;
    public ?array $recipient = null;
    public ?array $recurring_amount = null;
    public ?string $reference = null;
    public ?string $requested_execution_date = null;
    public ?string $required_action = null;
    public ?array $sender = null;
    public ?array $state = null;
    public ?array $supplementary_status = null;
}

/** Request payload for ManagePayment#create. */
class ManagePaymentCreateData
{
    public string $scheduled_id;
    public array $amount;
    public string $client_id;
    public string $created_at;
    public array $data;
    public array $errors;
    public string $first_payment_date;
    public string $frequency;
    public array $known_charges;
    public ?string $last_payment_date = null;
    public ?array $metadata = null;
    public string $method;
    public string $operation_id;
    public string $organisation_id;
    public string $payment_id;
    public array $recipient;
    public array $recurring_amount;
    public string $reference;
    public string $requested_execution_date;
    public string $required_action;
    public array $sender;
    public array $state;
    public array $supplementary_status;
}

