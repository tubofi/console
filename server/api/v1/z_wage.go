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

// CreateWage 创建Wage
// @Tags Wage
// @Summary 创建Wage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wage true "创建Wage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wage/createWage [post]
func CreateWage(c *gin.Context) {
	var wage model.Wage
	_ = c.ShouldBindJSON(&wage)
	if err := service.CreateWage(wage); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWage 删除Wage
// @Tags Wage
// @Summary 删除Wage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wage true "删除Wage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wage/deleteWage [delete]
func DeleteWage(c *gin.Context) {
	var wage model.Wage
	_ = c.ShouldBindJSON(&wage)
	if err := service.DeleteWage(wage); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWageByIds 批量删除Wage
// @Tags Wage
// @Summary 批量删除Wage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Wage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wage/deleteWageByIds [delete]
func DeleteWageByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteWageByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWage 更新Wage
// @Tags Wage
// @Summary 更新Wage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wage true "更新Wage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wage/updateWage [put]
func UpdateWage(c *gin.Context) {
	var wage model.Wage
	_ = c.ShouldBindJSON(&wage)
	if err := service.UpdateWage(wage); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWage 用id查询Wage
// @Tags Wage
// @Summary 用id查询Wage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wage true "用id查询Wage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wage/findWage [get]
func FindWage(c *gin.Context) {
	var wage model.Wage
	_ = c.ShouldBindQuery(&wage)
	if err, rewage := service.GetWage(wage.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败: " + err.Error(), c)
	} else {
		response.OkWithData(gin.H{"rewage": rewage}, c)
	}
}

// GetWageList 分页获取Wage列表
// @Tags Wage
// @Summary 分页获取Wage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.WageSearch true "分页获取Wage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wage/getWageList [get]
func GetWageList(c *gin.Context) {
	var pageInfo request.WageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetWageInfoList(pageInfo); err != nil {
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
