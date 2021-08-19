
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

// CreateCourseRecord 创建CourseRecord
// @Tags CourseRecord
// @Summary 创建CourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseRecord true "创建CourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /CourseRecord/createCourseRecord [post]
func CreateCourseRecord(c *gin.Context) {
	var CourseRecord model.CourseRecord
	_ = c.ShouldBindJSON(&CourseRecord)
	if err := service.CreateCourseRecord(CourseRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCourseRecord 删除CourseRecord
// @Tags CourseRecord
// @Summary 删除CourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseRecord true "删除CourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /CourseRecord/deleteCourseRecord [delete]
func DeleteCourseRecord(c *gin.Context) {
	var CourseRecord model.CourseRecord
	_ = c.ShouldBindJSON(&CourseRecord)
	if err := service.DeleteCourseRecord(CourseRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCourseRecordByIds 批量删除CourseRecord
// @Tags CourseRecord
// @Summary 批量删除CourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /CourseRecord/deleteCourseRecordByIds [delete]
func DeleteCourseRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteCourseRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCourseRecord 更新CourseRecord
// @Tags CourseRecord
// @Summary 更新CourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseRecord true "更新CourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /CourseRecord/updateCourseRecord [put]
func UpdateCourseRecord(c *gin.Context) {
	var CourseRecord model.CourseRecord
	_ = c.ShouldBindJSON(&CourseRecord)
	if err := service.UpdateCourseRecord(CourseRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCourseRecord 用id查询CourseRecord
// @Tags CourseRecord
// @Summary 用id查询CourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseRecord true "用id查询CourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /CourseRecord/findCourseRecord [get]
func FindCourseRecord(c *gin.Context) {
	var CourseRecord model.CourseRecord
	_ = c.ShouldBindQuery(&CourseRecord)
	if err, reCourseRecord := service.GetCourseRecord(CourseRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败: " + err.Error(), c)
	} else {
		response.OkWithData(gin.H{"recourseRecord": reCourseRecord}, c)
	}
}

// GetCourseRecordList 分页获取CourseRecord列表
// @Tags CourseRecord
// @Summary 分页获取CourseRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CourseRecordSearch true "分页获取CourseRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /CourseRecord/getCourseRecordList [get]
func GetCourseRecordList(c *gin.Context) {
	var pageInfo request.CourseRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetCourseRecordInfoList(pageInfo); err != nil {
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

func GetCramCourseRecordList(c *gin.Context) {
	var pageInfo request.CourseRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetCramCourseRecordInfoList(pageInfo); err != nil {
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

func CramCourseRecord(c *gin.Context) {
	var CourseRecord model.CourseRecord
	_ = c.ShouldBindJSON(&CourseRecord)
	if err := service.CramCourseRecord(CourseRecord); err != nil {
		global.GVA_LOG.Error("补课失败!", zap.Any("err", err))
		response.FailWithMessage("补课失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("补课成功", c)
	}
}

func IgnoreCramCourseRecord(c *gin.Context) {
	var CourseRecord model.CourseRecord
	_ = c.ShouldBindJSON(&CourseRecord)
	if err := service.IgnoreCramCourseRecord(CourseRecord); err != nil {
		global.GVA_LOG.Error("忽略补课失败!", zap.Any("err", err))
		response.FailWithMessage("忽略补课失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("已忽略", c)
	}
}

func GetFeedbackCourseRecordList(c *gin.Context) {
	var pageInfo request.CourseRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetFeedbackCourseRecordInfoList(pageInfo); err != nil {
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

func FeedbackCourseRecord(c *gin.Context) {
	var CourseRecord model.CourseRecord
	_ = c.ShouldBindJSON(&CourseRecord)
	if err := service.FeedbackCourseRecord(CourseRecord); err != nil {
		global.GVA_LOG.Error("补课失败!", zap.Any("err", err))
		response.FailWithMessage("补课失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("补课成功", c)
	}
}
