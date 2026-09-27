<?php
declare(strict_types=1);

// BudPayments SDK utility: result_body

class BudPaymentsResultBody
{
    public static function call(BudPaymentsContext $ctx): ?BudPaymentsResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
