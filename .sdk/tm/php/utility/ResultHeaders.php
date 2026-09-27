<?php
declare(strict_types=1);

// BudPayments SDK utility: result_headers

class BudPaymentsResultHeaders
{
    public static function call(BudPaymentsContext $ctx): ?BudPaymentsResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
