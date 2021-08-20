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

// CreateTrialCourseRecord 创建TrialCourseRecord
// @Tags TrialCourseRecord
// @Summary 创建TrialCourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TrialCourseRecord true "创建TrialCourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /trialCourseRecord/createTrialCourseRecord [post]
func CreateTrialCourseRecord(c *gin.Context) {
	var trialCourseRecord model.TrialCourseRecord
	_ = c.ShouldBindJSON(&trialCourseRecord)
	if err := service.CreateTrialCourseRecord(trialCourseRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTrialCourseRecord 删除TrialCourseRecord
// @Tags TrialCourseRecord
// @Summary 删除TrialCourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TrialCourseRecord true "删除TrialCourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /trialCourseRecord/deleteTrialCourseRecord [delete]
func DeleteTrialCourseRecord(c *gin.Context) {
	var trialCourseRecord model.TrialCourseRecord
	_ = c.ShouldBindJSON(&trialCourseRecord)
	if err := service.DeleteTrialCourseRecord(trialCourseRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTrialCourseRecordByIds 批量删除TrialCourseRecord
// @Tags TrialCourseRecord
// @Summary 批量删除TrialCourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TrialCourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /trialCourseRecord/deleteTrialCourseRecordByIds [delete]
func DeleteTrialCourseRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTrialCourseRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTrialCourseRecord 更新TrialCourseRecord
// @Tags TrialCourseRecord
// @Summary 更新TrialCourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TrialCourseRecord true "更新TrialCourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /trialCourseRecord/updateTrialCourseRecord [put]
func UpdateTrialCourseRecord(c *gin.Context) {
	var trialCourseRecord model.TrialCourseRecord
	_ = c.ShouldBindJSON(&trialCourseRecord)
	if err := service.UpdateTrialCourseRecord(trialCourseRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTrialCourseRecord 用id查询TrialCourseRecord
// @Tags TrialCourseRecord
// @Summary 用id查询TrialCourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TrialCourseRecord true "用id查询TrialCourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /trialCourseRecord/findTrialCourseRecord [get]
func FindTrialCourseRecord(c *gin.Context) {
	var trialCourseRecord model.TrialCourseRecord
	_ = c.ShouldBindQuery(&trialCourseRecord)
	if err, retrialCourseRecord := service.GetTrialCourseRecord(trialCourseRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败: " + err.Error(), c)
	} else {
		response.OkWithData(gin.H{"retrialCourseRecord": retrialCourseRecord}, c)
	}
}

// GetTrialCourseRecordList 分页获取TrialCourseRecord列表
// @Tags TrialCourseRecord
// @Summary 分页获取TrialCourseRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TrialCourseRecordSearch true "分页获取TrialCourseRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /trialCourseRecord/getTrialCourseRecordList [get]
func GetTrialCourseRecordList(c *gin.Context) {
	var pageInfo request.TrialCourseRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTrialCourseRecordInfoList(pageInfo); err != nil {
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

func GetAllTrialStudents(c *gin.Context) {
	if err, trialStudents := service.GetAllTrialStudents(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败: " + err.Error(), c)
	} else {
		response.OkWithDetailed(trialStudents, "获取成功", c)
	}
}

func FeedbackTrialCourseRecord(c *gin.Context) {
	var trialCourseRecord model.TrialCourseRecord
	_ = c.ShouldBindJSON(&trialCourseRecord)
	if err := service.FeedbackTrialCourseRecord(trialCourseRecord); err != nil {
		global.GVA_LOG.Error("反馈失败!", zap.Any("err", err))
		response.FailWithMessage("反馈失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("反馈成功", c)
	}
}


