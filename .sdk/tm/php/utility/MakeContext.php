<?php
declare(strict_types=1);

// BudPayments SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class BudPaymentsMakeContext
{
    public static function call(array $ctxmap, ?BudPaymentsContext $basectx): BudPaymentsContext
    {
        return new BudPaymentsContext($ctxmap, $basectx);
    }
}
