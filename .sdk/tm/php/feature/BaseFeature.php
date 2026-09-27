<?php
declare(strict_types=1);

// BudPayments SDK base feature

class BudPaymentsBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(BudPaymentsContext $ctx, array $options): void {}
    public function PostConstruct(BudPaymentsContext $ctx): void {}
    public function PostConstructEntity(BudPaymentsContext $ctx): void {}
    public function SetData(BudPaymentsContext $ctx): void {}
    public function GetData(BudPaymentsContext $ctx): void {}
    public function GetMatch(BudPaymentsContext $ctx): void {}
    public function SetMatch(BudPaymentsContext $ctx): void {}
    public function PrePoint(BudPaymentsContext $ctx): void {}
    public function PreSpec(BudPaymentsContext $ctx): void {}
    public function PreRequest(BudPaymentsContext $ctx): void {}
    public function PreResponse(BudPaymentsContext $ctx): void {}
    public function PreResult(BudPaymentsContext $ctx): void {}
    public function PreDone(BudPaymentsContext $ctx): void {}
    public function PreUnexpected(BudPaymentsContext $ctx): void {}
}
