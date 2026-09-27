<?php
declare(strict_types=1);

// BudPayments SDK exists test

require_once __DIR__ . '/../budpayments_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = BudPaymentsSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
