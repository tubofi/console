package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitPotentialRouter 初始化 Potential 路由信息
func InitPotentialRouter(Router *gin.RouterGroup) {
	PotentialRouter := Router.Group("potential").Use(middleware.OperationRecord())
	{
		PotentialRouter.POST("createPotential", v1.CreatePotential)   // 新建Potential
		PotentialRouter.DELETE("deletePotential", v1.DeletePotential) // 删除Potential
		PotentialRouter.PUT("updatePotential", v1.UpdatePotential)    // 更新Potential
		PotentialRouter.GET("findPotential", v1.FindPotential)        // 根据ID获取Potential
		PotentialRouter.GET("getPotentialList", v1.GetPotentialList)  // 获取Potential列表
		PotentialRouter.PUT("potentialToTrial", v1.PotentialToTrial)  // 意向学员转为试听学员
		PotentialRouter.PUT("upgradePotentialLevel", v1.UpgradePotentialLevel)  // 提升状态等级
		PotentialRouter.PUT("degradePotentialLevel", v1.DegradePotentialLevel)  // 降低状态等级
	}
}
