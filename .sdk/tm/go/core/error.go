package core

type BudPaymentsError struct {
	IsBudPaymentsError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewBudPaymentsError(code string, msg string, ctx *Context) *BudPaymentsError {
	return &BudPaymentsError{
		IsBudPaymentsError: true,
		Sdk:              "BudPayments",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *BudPaymentsError) Error() string {
	return e.Msg
}
