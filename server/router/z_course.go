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
	}
}
