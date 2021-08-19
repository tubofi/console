package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitClassRouter 初始化 Class 路由信息
func InitClassRouter(Router *gin.RouterGroup) {
	ClassRouter := Router.Group("class").Use(middleware.OperationRecord())
	{
		ClassRouter.POST("createClass", v1.CreateClass)   // 新建Class
		ClassRouter.DELETE("deleteClass", v1.DeleteClass) // 删除Class
		ClassRouter.PUT("updateClass", v1.UpdateClass)    // 更新Class
		ClassRouter.GET("findClass", v1.FindClass)        // 根据ID获取Class
		ClassRouter.GET("getClassList", v1.GetClassList)  // 获取Class列表
		ClassRouter.GET("getIdleStudents", v1.GetIdleStudents)  // 获取没有分配班级学生列表
		ClassRouter.GET("getProfileClass", v1.GetProfileClass)  // 获取概况预览页班级数据

		ClassRouter.POST("createCourseRecord", v1.CreateCourseRecord)   // 新建CourseRecord
		ClassRouter.DELETE("deleteCourseRecord", v1.DeleteCourseRecord) // 删除CourseRecord
		ClassRouter.DELETE("deleteCourseRecordByIds", v1.DeleteCourseRecordByIds) // 批量删除CourseRecord
		ClassRouter.PUT("updateCourseRecord", v1.UpdateCourseRecord)    // 更新CourseRecord
		ClassRouter.GET("findCourseRecord", v1.FindCourseRecord)        // 根据ID获取CourseRecord
		ClassRouter.GET("getCourseRecordList", v1.GetCourseRecordList)  // 获取CourseRecord列表
		ClassRouter.GET("getCramCourseRecordList", v1.GetCramCourseRecordList)  // 获取CourseRecord列表
		ClassRouter.PUT("cramCourseRecord", v1.CramCourseRecord)  				// 获取CourseRecord列表
		ClassRouter.GET("getFeedbackCourseRecordList", v1.GetFeedbackCourseRecordList)  // 获取CourseRecord列表
		ClassRouter.PUT("feedbackCourseRecord", v1.FeedbackCourseRecord)  				// 获取CourseRecord列表
		ClassRouter.PUT("ignoreCramCourseRecord", v1.IgnoreCramCourseRecord)  				// 获取CourseRecord列表
	}
}