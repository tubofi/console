package request

import "gin-vue-admin/model"

type LivingBillSearch struct{
	model.LivingBill
	PageInfo
}

type LivingCostSearch struct{
	model.LivingCost
	PageInfo
}