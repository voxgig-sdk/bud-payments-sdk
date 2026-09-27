# BudPayments SDK feature factory

from budpayments_sdk.feature.base_feature import BudPaymentsBaseFeature
from budpayments_sdk.feature.debug_feature import BudPaymentsDebugFeature
from budpayments_sdk.feature.idempotency_feature import BudPaymentsIdempotencyFeature
from budpayments_sdk.feature.metrics_feature import BudPaymentsMetricsFeature
from budpayments_sdk.feature.paging_feature import BudPaymentsPagingFeature
from budpayments_sdk.feature.ratelimit_feature import BudPaymentsRatelimitFeature
from budpayments_sdk.feature.retry_feature import BudPaymentsRetryFeature
from budpayments_sdk.feature.test_feature import BudPaymentsTestFeature
from budpayments_sdk.feature.timeout_feature import BudPaymentsTimeoutFeature


_FEATURES = {
    "base": lambda: BudPaymentsBaseFeature(),
    "debug": lambda: BudPaymentsDebugFeature(),
    "idempotency": lambda: BudPaymentsIdempotencyFeature(),
    "metrics": lambda: BudPaymentsMetricsFeature(),
    "paging": lambda: BudPaymentsPagingFeature(),
    "ratelimit": lambda: BudPaymentsRatelimitFeature(),
    "retry": lambda: BudPaymentsRetryFeature(),
    "test": lambda: BudPaymentsTestFeature(),
    "timeout": lambda: BudPaymentsTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
