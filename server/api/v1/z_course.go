
package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateCourse(c *gin.Context) {
	var resCourse model.ResCourse
	_ = c.ShouldBindJSON(&resCourse)
	if err := service.CreateCourse(resCourse); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func DeleteCourse(c *gin.Context) {
	var resCourse model.ResCourse
	_ = c.ShouldBindJSON(&resCourse)
	if err := service.DeleteCourse(resCourse); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func UpdateCourse(c *gin.Context) {
	var resCourse model.ResCourse
	_ = c.ShouldBindJSON(&resCourse)
	if err := service.UpdateCourse(resCourse); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func FindCourse(c *gin.Context) {
	var resCourse model.ResCourse
	_ = c.ShouldBindQuery(&resCourse)
	if err, recourse := service.GetCourse(resCourse.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败: " + err.Error(), c)
	} else {
		response.OkWithData(gin.H{"recourse": recourse}, c)
	}
}

func GetCourseList(c *gin.Context) {
	var pageInfo request.CourseSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetCourseInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败: " + err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func GetClassListFromCourse(c *gin.Context) {
	if err, list := service.GetClassListFromCourse(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败: " + err.Error(), c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

