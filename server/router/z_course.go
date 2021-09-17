package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitCourseRouter 初始化 Course 路由信息
func InitCourseRouter(Router *gin.RouterGroup) {
	CourseRouter := Router.Group("course").Use(middleware.OperationRecord())
	{
		CourseRouter.POST("createCourse", v1.CreateCourse)   // 新建Course
		CourseRouter.DELETE("deleteCourse", v1.DeleteCourse) // 删除Course
		CourseRouter.PUT("updateCourse", v1.UpdateCourse)    // 更新Course
		CourseRouter.GET("findCourse", v1.FindCourse)        // 根据ID获取Course
		CourseRouter.GET("getCourseList", v1.GetCourseList)  // 获取Course列表
		CourseRouter.GET("getClassListFromCourse", v1.GetClassListFromCourse)  // 获取所有班级列表

		CourseRouter.POST("createCourseChange", v1.CreateCourseChange)   // 新建CourseChange
		CourseRouter.DELETE("deleteCourseChange", v1.DeleteCourseChange) // 删除CourseChange
		CourseRouter.DELETE("deleteCourseChangeByIds", v1.DeleteCourseChangeByIds) // 批量删除CourseChange
		CourseRouter.PUT("updateCourseChange", v1.UpdateCourseChange)    // 更新CourseChange
		CourseRouter.GET("findCourseChange", v1.FindCourseChange)        // 根据ID获取CourseChange
		CourseRouter.GET("getCourseChangeList", v1.GetCourseChangeList)  // 获取CourseChange列表
	}
}


