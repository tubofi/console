package router

import (
"gin-vue-admin/api/v1"
"gin-vue-admin/middleware"
"github.com/gin-gonic/gin"
)

// InitUploadRouter 初始化 Upload 路由信息
func InitUploadRouter(Router *gin.RouterGroup) {
	UploadRouter := Router.Group("upload").Use(middleware.OperationRecord())
	{
		UploadRouter.POST("createUpload", v1.CreateUpload)   // 新建Upload
		UploadRouter.DELETE("deleteUpload", v1.DeleteUpload) // 删除Upload
		UploadRouter.DELETE("deleteUploadByIds", v1.DeleteUploadByIds) // 批量删除Upload
		UploadRouter.PUT("updateUpload", v1.UpdateUpload)    // 更新Upload
		UploadRouter.GET("findUpload", v1.FindUpload)        // 根据ID获取Upload
		UploadRouter.GET("getUploadList", v1.GetUploadList)  // 获取Upload列表

		UploadRouter.GET("getTmpSecret", v1.GetTmpSecret)  				// 获取访问cos临时密钥
		UploadRouter.GET("getImageTmpSecret", v1.GetImageTmpSecret)  		// 获取访问cos临时密钥-images这个bucket
	}
}