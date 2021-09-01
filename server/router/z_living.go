
package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitLivingBillRouter 初始化 LivingBill 路由信息
func InitLivingRouter(Router *gin.RouterGroup) {
	LivingRouter := Router.Group("living").Use(middleware.OperationRecord())
	{
		LivingRouter.POST("createLivingBill", v1.CreateLivingBill)   // 新建LivingBill
		LivingRouter.DELETE("deleteLivingBill", v1.DeleteLivingBill) // 删除LivingBill
		LivingRouter.DELETE("deleteLivingBillByIds", v1.DeleteLivingBillByIds) // 批量删除LivingBill
		LivingRouter.PUT("updateLivingBill", v1.UpdateLivingBill)    // 更新LivingBill
		LivingRouter.GET("findLivingBill", v1.FindLivingBill)        // 根据ID获取LivingBill
		LivingRouter.GET("getLivingBillList", v1.GetLivingBillList)  // 获取LivingBill列表

		LivingRouter.POST("createLivingCost", v1.CreateLivingCost)   // 新建LivingCost
		LivingRouter.DELETE("deleteLivingCost", v1.DeleteLivingCost) // 删除LivingCost
		LivingRouter.DELETE("deleteLivingCostByIds", v1.DeleteLivingCostByIds) // 批量删除LivingCost
		LivingRouter.PUT("updateLivingCost", v1.UpdateLivingCost)    // 更新LivingCost
		LivingRouter.GET("findLivingCost", v1.FindLivingCost)        // 根据ID获取LivingCost
		LivingRouter.GET("getLivingCostList", v1.GetLivingCostList)  // 获取LivingCost列表
	}
}
