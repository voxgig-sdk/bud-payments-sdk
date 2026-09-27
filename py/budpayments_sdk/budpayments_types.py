# Typed models for the BudPayments SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class InitiatePaymentBudLicenseRequired(TypedDict):
    bud_pay_url: str
    country_code: str
    display_name: str
    icon: str
    implementation_type: str
    maintenance_status: str
    maintenance_window: Any
    payment_details: dict
    payment_id: str
    provider: str
    redirect_url: str
    required_actions: Any
    scheduled_payment_details: dict
    services: list
    standing_order_details: dict
    supported_currencies: list


class InitiatePaymentBudLicense(InitiatePaymentBudLicenseRequired, total=False):
    provider_types: list


class InitiatePaymentBudLicenseListMatch(TypedDict, total=False):
    country: str
    product: list
    type: str


class InitiatePaymentBudLicenseCreateDataRequired(TypedDict):
    bud_pay_url: str
    country_code: str
    display_name: str
    icon: str
    implementation_type: str
    maintenance_status: str
    maintenance_window: Any
    payment_details: dict
    payment_id: str
    provider: str
    redirect_url: str
    required_actions: Any
    scheduled_payment_details: dict
    services: list
    standing_order_details: dict
    supported_currencies: list


class InitiatePaymentBudLicenseCreateData(InitiatePaymentBudLicenseCreateDataRequired, total=False):
    provider_types: list


class InitiatePaymentClientLicenseRequired(TypedDict):
    authorisation_url: str
    payment_details: dict
    provider: str
    scheduled_payment_details: dict
    standing_order_details: dict
    state: str


class InitiatePaymentClientLicense(InitiatePaymentClientLicenseRequired, total=False):
    code: str
    payment_id: str
    payment_type: str
    provider_redirect_url: str
    provider_types: list
    redirect_url: str
    required_action: str
    result: str


class InitiatePaymentClientLicenseCreateDataRequired(TypedDict):
    authorisation_url: str
    payment_details: dict
    provider: str
    scheduled_payment_details: dict
    standing_order_details: dict
    state: str


class InitiatePaymentClientLicenseCreateData(InitiatePaymentClientLicenseCreateDataRequired, total=False):
    code: str
    payment_id: str
    payment_type: str
    provider_redirect_url: str
    provider_types: list
    redirect_url: str
    required_action: str
    result: str


class ManagePaymentRequired(TypedDict):
    amount: dict
    client_id: str
    created_at: str
    data: dict
    errors: list
    first_payment_date: str
    frequency: str
    known_charges: dict
    method: str
    operation_id: str
    organisation_id: str
    payment_id: str
    recipient: dict
    recurring_amount: dict
    reference: str
    requested_execution_date: str
    required_action: str
    sender: dict
    state: dict
    supplementary_status: dict


class ManagePayment(ManagePaymentRequired, total=False):
    last_payment_date: str
    metadata: dict


class ManagePaymentLoadMatch(TypedDict):
    payment_id: str


class ManagePaymentListMatch(TypedDict, total=False):
    amount: dict
    client_id: str
    created_at: str
    data: dict
    errors: list
    first_payment_date: str
    frequency: str
    known_charges: dict
    last_payment_date: str
    metadata: dict
    method: str
    operation_id: str
    organisation_id: str
    payment_id: str
    recipient: dict
    recurring_amount: dict
    reference: str
    requested_execution_date: str
    required_action: str
    sender: dict
    state: dict
    supplementary_status: dict


class ManagePaymentCreateDataRequired(TypedDict):
    scheduled_id: str
    amount: dict
    client_id: str
    created_at: str
    data: dict
    errors: list
    first_payment_date: str
    frequency: str
    known_charges: dict
    method: str
    operation_id: str
    organisation_id: str
    payment_id: str
    recipient: dict
    recurring_amount: dict
    reference: str
    requested_execution_date: str
    required_action: str
    sender: dict
    state: dict
    supplementary_status: dict


class ManagePaymentCreateData(ManagePaymentCreateDataRequired, total=False):
    last_payment_date: str
    metadata: dict
