<?php
declare(strict_types=1);

// BudPayments SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/DebugFeature.php';
require_once __DIR__ . '/feature/IdempotencyFeature.php';
require_once __DIR__ . '/feature/MetricsFeature.php';
require_once __DIR__ . '/feature/PagingFeature.php';
require_once __DIR__ . '/feature/RatelimitFeature.php';
require_once __DIR__ . '/feature/RetryFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';
require_once __DIR__ . '/feature/TimeoutFeature.php';


class BudPaymentsFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new BudPaymentsBaseFeature();
            case "debug":
                return new BudPaymentsDebugFeature();
            case "idempotency":
                return new BudPaymentsIdempotencyFeature();
            case "metrics":
                return new BudPaymentsMetricsFeature();
            case "paging":
                return new BudPaymentsPagingFeature();
            case "ratelimit":
                return new BudPaymentsRatelimitFeature();
            case "retry":
                return new BudPaymentsRetryFeature();
            case "test":
                return new BudPaymentsTestFeature();
            case "timeout":
                return new BudPaymentsTimeoutFeature();
            default:
                return new BudPaymentsBaseFeature();
        }
    }

    /**
     * Does a generated feature class back this name? False for a name only
     * an options extend instance can supply (the station adopt path) - the
     * constructor uses this to skip make_feature for such names instead of
     * adding a stray BaseFeature.
     */
    public static function has_feature(string $name): bool
    {
        switch ($name) {
            case "base":
            case "debug":
            case "idempotency":
            case "metrics":
            case "paging":
            case "ratelimit":
            case "retry":
            case "test":
            case "timeout":
                return true;
            default:
                return false;
        }
    }
}
