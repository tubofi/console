package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitTrialCourseRecordRouter 初始化 TrialCourseRecord 路由信息
func InitTrialRouter(Router *gin.RouterGroup) {
	TrialRouter := Router.Group("trial").Use(middleware.OperationRecord())
	{
		TrialRouter.POST("createTrialCourseRecord", v1.CreateTrialCourseRecord)   // 新建TrialCourseRecord
		TrialRouter.DELETE("deleteTrialCourseRecord", v1.DeleteTrialCourseRecord) // 删除TrialCourseRecord
		TrialRouter.DELETE("deleteTrialCourseRecordByIds", v1.DeleteTrialCourseRecordByIds) // 批量删除TrialCourseRecord
		TrialRouter.PUT("updateTrialCourseRecord", v1.UpdateTrialCourseRecord)    // 更新TrialCourseRecord
		TrialRouter.GET("findTrialCourseRecord", v1.FindTrialCourseRecord)        // 根据ID获取TrialCourseRecord
		TrialRouter.GET("getTrialCourseRecordList", v1.GetTrialCourseRecordList)  // 获取TrialCourseRecord列表
		TrialRouter.GET("getAllTrialStudents", v1.GetAllTrialStudents)  		// 获取所有体验学生列表
		TrialRouter.PUT("feedbackTrialCourseRecord", v1.FeedbackTrialCourseRecord)  		// 获取所有体验学生列表
	}
}