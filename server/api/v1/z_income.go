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

// CreateIncome 创建Income
// @Tags Income
// @Summary 创建Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Income true "创建Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /income/createIncome [post]
func CreateIncome(c *gin.Context) {
	var income model.Income
	_ = c.ShouldBindJSON(&income)
	if err := service.CreateIncome(income); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIncome 删除Income
// @Tags Income
// @Summary 删除Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Income true "删除Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /income/deleteIncome [delete]
func DeleteIncome(c *gin.Context) {
	var income model.Income
	_ = c.ShouldBindJSON(&income)
	if err := service.DeleteIncome(income); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIncomeByIds 批量删除Income
// @Tags Income
// @Summary 批量删除Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /income/deleteIncomeByIds [delete]
func DeleteIncomeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteIncomeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIncome 更新Income
// @Tags Income
// @Summary 更新Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Income true "更新Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /income/updateIncome [put]
func UpdateIncome(c *gin.Context) {
	var income model.Income
	_ = c.ShouldBindJSON(&income)
	if err := service.UpdateIncome(income); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIncome 用id查询Income
// @Tags Income
// @Summary 用id查询Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Income true "用id查询Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /income/findIncome [get]
func FindIncome(c *gin.Context) {
	var income model.Income
	_ = c.ShouldBindQuery(&income)
	if err, reincome := service.GetIncome(income.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败: " + err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reincome": reincome}, c)
	}
}

// GetIncomeList 分页获取Income列表
// @Tags Income
// @Summary 分页获取Income列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IncomeSearch true "分页获取Income列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /income/getIncomeList [get]
func GetIncomeList(c *gin.Context) {
	var pageInfo request.IncomeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetIncomeInfoList(pageInfo); err != nil {
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
