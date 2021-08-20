package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitTeacherRouter 初始化 Teacher 路由信息
func InitTeacherRouter(Router *gin.RouterGroup) {
	TeacherRouter := Router.Group("teacher").Use(middleware.OperationRecord())
	{
		TeacherRouter.POST("createTeacher", v1.CreateTeacher)   // 新建Teacher
		TeacherRouter.DELETE("deleteTeacher", v1.DeleteTeacher) // 删除Teacher
		TeacherRouter.DELETE("deleteTeacherByIds", v1.DeleteTeacherByIds) // 批量删除Teacher
		TeacherRouter.PUT("updateTeacher", v1.UpdateTeacher)    // 更新Teacher
		TeacherRouter.GET("findTeacher", v1.FindTeacher)        // 根据ID获取Teacher
		TeacherRouter.GET("getTeacherList", v1.GetTeacherList)  // 获取Teacher列表
		TeacherRouter.GET("getAllTeachers", v1.GetAllTeachers)  // 获取所有Teacher列表
	}
}