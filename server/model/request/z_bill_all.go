package request

import "gin-vue-admin/model"

type BillSearch struct{
	model.Bill
	PageInfo
}

type PaymentSearch struct{
	model.Payment
	PageInfo
}

type IncomeSearch struct{
	model.Income
	PageInfo
}

type OutcomeSearch struct{
	model.Outcome
	PageInfo
}

type WageSearch struct{
	model.Wage
	PageInfo
}