package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewDebugFeatureFunc func() Feature

var NewIdempotencyFeatureFunc func() Feature

var NewMetricsFeatureFunc func() Feature

var NewPagingFeatureFunc func() Feature

var NewRatelimitFeatureFunc func() Feature

var NewRetryFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewTimeoutFeatureFunc func() Feature

var NewInitiatePaymentBudLicenseEntityFunc func(client *BudPaymentsSDK, entopts map[string]any) BudPaymentsEntity

var NewInitiatePaymentClientLicenseEntityFunc func(client *BudPaymentsSDK, entopts map[string]any) BudPaymentsEntity

var NewManagePaymentEntityFunc func(client *BudPaymentsSDK, entopts map[string]any) BudPaymentsEntity

