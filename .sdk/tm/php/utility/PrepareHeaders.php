<?php
declare(strict_types=1);

// BudPayments SDK utility: prepare_headers

class BudPaymentsPrepareHeaders
{
    public static function call(BudPaymentsContext $ctx): array
    {
        $options = $ctx->client->options_map();
        $headers = \Voxgig\Struct\Struct::getprop($options, 'headers');
        if (!$headers) {
            return [];
        }
        $out = \Voxgig\Struct\Struct::clone($headers);
        return is_array($out) ? $out : [];
    }
}
