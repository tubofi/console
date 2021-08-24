package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserDefinedRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("trial").Use(middleware.OperationRecord())
	{
		UserRouter.GET("getCredential", v1.GetCredential)  		// 获取访问cos临时密钥
	}
}
