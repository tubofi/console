
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

// CreateCourseChange 创建CourseChange
// @Tags CourseChange
// @Summary 创建CourseChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseChange true "创建CourseChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /courseChange/createCourseChange [post]
func CreateCourseChange(c *gin.Context) {
	var courseChange model.CourseChange
	_ = c.ShouldBindJSON(&courseChange)
	if err := service.CreateCourseChange(courseChange); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCourseChange 删除CourseChange
// @Tags CourseChange
// @Summary 删除CourseChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseChange true "删除CourseChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /courseChange/deleteCourseChange [delete]
func DeleteCourseChange(c *gin.Context) {
	var courseChange model.CourseChange
	_ = c.ShouldBindJSON(&courseChange)
	if err := service.DeleteCourseChange(courseChange); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCourseChangeByIds 批量删除CourseChange
// @Tags CourseChange
// @Summary 批量删除CourseChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CourseChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /courseChange/deleteCourseChangeByIds [delete]
func DeleteCourseChangeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteCourseChangeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCourseChange 更新CourseChange
// @Tags CourseChange
// @Summary 更新CourseChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseChange true "更新CourseChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /courseChange/updateCourseChange [put]
func UpdateCourseChange(c *gin.Context) {
	var courseChange model.CourseChange
	_ = c.ShouldBindJSON(&courseChange)
	if err := service.UpdateCourseChange(courseChange); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCourseChange 用id查询CourseChange
// @Tags CourseChange
// @Summary 用id查询CourseChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseChange true "用id查询CourseChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /courseChange/findCourseChange [get]
func FindCourseChange(c *gin.Context) {
	var courseChange model.CourseChange
	_ = c.ShouldBindQuery(&courseChange)
	if err, recourseChange := service.GetCourseChange(courseChange.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recourseChange": recourseChange}, c)
	}
}

// GetCourseChangeList 分页获取CourseChange列表
// @Tags CourseChange
// @Summary 分页获取CourseChange列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CourseChangeSearch true "分页获取CourseChange列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /courseChange/getCourseChangeList [get]
func GetCourseChangeList(c *gin.Context) {
	var pageInfo request.CourseChangeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetCourseChangeInfoList(pageInfo); err != nil {
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
