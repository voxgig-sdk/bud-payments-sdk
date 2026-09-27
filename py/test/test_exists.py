# BudPayments SDK exists test

import pytest
from budpayments_sdk import BudPaymentsSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = BudPaymentsSDK.test(None, None)
        assert testsdk is not None
