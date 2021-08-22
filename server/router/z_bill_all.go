package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitBillRouter 初始化 Bill 路由信息
func InitBillRouter(Router *gin.RouterGroup) {
	BillRouter := Router.Group("bill").Use(middleware.OperationRecord())
	{
		//bill
		BillRouter.POST("createBill", v1.CreateBill)   // 新建Bill
		BillRouter.DELETE("deleteBill", v1.DeleteBill) // 删除Bill
		BillRouter.DELETE("deleteBillByIds", v1.DeleteBillByIds) // 批量删除Bill
		BillRouter.PUT("updateBill", v1.UpdateBill)    // 更新Bill
		BillRouter.GET("findBill", v1.FindBill)        // 根据ID获取Bill
		BillRouter.GET("getBillList", v1.GetBillList)  // 获取Bill列表
		//payment
		BillRouter.POST("createPayment", v1.CreatePayment)   // 新建Payment
		BillRouter.DELETE("deletePayment", v1.DeletePayment) // 删除Payment
		BillRouter.DELETE("deletePaymentByIds", v1.DeletePaymentByIds) // 批量删除Payment
		BillRouter.PUT("updatePayment", v1.UpdatePayment)    // 更新Payment
		BillRouter.GET("findPayment", v1.FindPayment)        // 根据ID获取Payment
		BillRouter.GET("getPaymentList", v1.GetPaymentList)  // 获取Payment列表
		//income
		BillRouter.POST("createIncome", v1.CreateIncome)   // 新建Income
		BillRouter.DELETE("deleteIncome", v1.DeleteIncome) // 删除Income
		BillRouter.DELETE("deleteIncomeByIds", v1.DeleteIncomeByIds) // 批量删除Income
		BillRouter.PUT("updateIncome", v1.UpdateIncome)    // 更新Income
		BillRouter.GET("findIncome", v1.FindIncome)        // 根据ID获取Income
		BillRouter.GET("getIncomeList", v1.GetIncomeList)  // 获取Income列表
		//outcome
		BillRouter.POST("createOutcome", v1.CreateOutcome)   // 新建Outcome
		BillRouter.DELETE("deleteOutcome", v1.DeleteOutcome) // 删除Outcome
		BillRouter.DELETE("deleteOutcomeByIds", v1.DeleteOutcomeByIds) // 批量删除Outcome
		BillRouter.PUT("updateOutcome", v1.UpdateOutcome)    // 更新Outcome
		BillRouter.GET("findOutcome", v1.FindOutcome)        // 根据ID获取Outcome
		BillRouter.GET("getOutcomeList", v1.GetOutcomeList)  // 获取Outcome列表
		//wage
		BillRouter.POST("createWage", v1.CreateWage)   // 新建Wage
		BillRouter.DELETE("deleteWage", v1.DeleteWage) // 删除Wage
		BillRouter.DELETE("deleteWageByIds", v1.DeleteWageByIds) // 批量删除Wage
		BillRouter.PUT("updateWage", v1.UpdateWage)    // 更新Wage
		BillRouter.GET("findWage", v1.FindWage)        // 根据ID获取Wage
		BillRouter.GET("getWageList", v1.GetWageList)  // 获取Wage列表

		BillRouter.GET("getAllStudents", v1.GetAllStudents)  	// 获取所有学生
	}
}


