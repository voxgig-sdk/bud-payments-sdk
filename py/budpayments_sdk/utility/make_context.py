# BudPayments SDK utility: make_context

from budpayments_sdk.core.context import BudPaymentsContext


def make_context_util(ctxmap, basectx):
    return BudPaymentsContext(ctxmap, basectx)
