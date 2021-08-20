
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

// CreateTeacher 创建Teacher
// @Tags Teacher
// @Summary 创建Teacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Teacher true "创建Teacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teacher/createTeacher [post]
func CreateTeacher(c *gin.Context) {
	var teacher model.Teacher
	_ = c.ShouldBindJSON(&teacher)
	if err := service.CreateTeacher(teacher); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeacher 删除Teacher
// @Tags Teacher
// @Summary 删除Teacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Teacher true "删除Teacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teacher/deleteTeacher [delete]
func DeleteTeacher(c *gin.Context) {
	var teacher model.Teacher
	_ = c.ShouldBindJSON(&teacher)
	if err := service.DeleteTeacher(teacher); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeacherByIds 批量删除Teacher
// @Tags Teacher
// @Summary 批量删除Teacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Teacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teacher/deleteTeacherByIds [delete]
func DeleteTeacherByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTeacherByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeacher 更新Teacher
// @Tags Teacher
// @Summary 更新Teacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Teacher true "更新Teacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teacher/updateTeacher [put]
func UpdateTeacher(c *gin.Context) {
	var teacher model.Teacher
	_ = c.ShouldBindJSON(&teacher)
	if err := service.UpdateTeacher(teacher); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeacher 用id查询Teacher
// @Tags Teacher
// @Summary 用id查询Teacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Teacher true "用id查询Teacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teacher/findTeacher [get]
func FindTeacher(c *gin.Context) {
	var teacher model.Teacher
	_ = c.ShouldBindQuery(&teacher)
	if err, reteacher := service.GetTeacher(teacher.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteacher": reteacher}, c)
	}
}

// GetTeacherList 分页获取Teacher列表
// @Tags Teacher
// @Summary 分页获取Teacher列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TeacherSearch true "分页获取Teacher列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teacher/getTeacherList [get]
func GetTeacherList(c *gin.Context) {
	var pageInfo request.TeacherSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTeacherInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func GetAllTeachers(c *gin.Context) {
	if err, teachers := service.GetAllTeachers(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败: " + err.Error(), c)
	} else {
		response.OkWithDetailed(teachers, "获取成功!", c)
	}
}
