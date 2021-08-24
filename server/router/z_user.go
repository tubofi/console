package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCosRouter(Router *gin.RouterGroup) {
	CosRouter := Router.Group("cos").Use(middleware.OperationRecord())
	{
		CosRouter.GET("getTmpSecret", v1.GetTmpSecret)  		// 获取访问cos临时密钥
	}
}
