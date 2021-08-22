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

// CreateOutcome 创建Outcome
// @Tags Outcome
// @Summary 创建Outcome
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Outcome true "创建Outcome"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /outcome/createOutcome [post]
func CreateOutcome(c *gin.Context) {
	var outcome model.Outcome
	_ = c.ShouldBindJSON(&outcome)
	if err := service.CreateOutcome(outcome); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOutcome 删除Outcome
// @Tags Outcome
// @Summary 删除Outcome
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Outcome true "删除Outcome"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /outcome/deleteOutcome [delete]
func DeleteOutcome(c *gin.Context) {
	var outcome model.Outcome
	_ = c.ShouldBindJSON(&outcome)
	if err := service.DeleteOutcome(outcome); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOutcomeByIds 批量删除Outcome
// @Tags Outcome
// @Summary 批量删除Outcome
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Outcome"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /outcome/deleteOutcomeByIds [delete]
func DeleteOutcomeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteOutcomeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOutcome 更新Outcome
// @Tags Outcome
// @Summary 更新Outcome
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Outcome true "更新Outcome"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /outcome/updateOutcome [put]
func UpdateOutcome(c *gin.Context) {
	var outcome model.Outcome
	_ = c.ShouldBindJSON(&outcome)
	if err := service.UpdateOutcome(outcome); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOutcome 用id查询Outcome
// @Tags Outcome
// @Summary 用id查询Outcome
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Outcome true "用id查询Outcome"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /outcome/findOutcome [get]
func FindOutcome(c *gin.Context) {
	var outcome model.Outcome
	_ = c.ShouldBindQuery(&outcome)
	if err, reoutcome := service.GetOutcome(outcome.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reoutcome": reoutcome}, c)
	}
}

// GetOutcomeList 分页获取Outcome列表
// @Tags Outcome
// @Summary 分页获取Outcome列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.OutcomeSearch true "分页获取Outcome列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /outcome/getOutcomeList [get]
func GetOutcomeList(c *gin.Context) {
	var pageInfo request.OutcomeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetOutcomeInfoList(pageInfo); err != nil {
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
