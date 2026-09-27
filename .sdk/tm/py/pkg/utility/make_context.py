# BudPayments SDK utility: make_context

from projectname_sdk.core.context import BudPaymentsContext


def make_context_util(ctxmap, basectx):
    return BudPaymentsContext(ctxmap, basectx)
