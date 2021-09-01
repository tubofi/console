
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

// CreateLivingBill 创建LivingBill
// @Tags LivingBill
// @Summary 创建LivingBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingBill true "创建LivingBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /livingBill/createLivingBill [post]
func CreateLivingBill(c *gin.Context) {
	var livingBill model.LivingBill
	_ = c.ShouldBindJSON(&livingBill)
	if err := service.CreateLivingBill(livingBill); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLivingBill 删除LivingBill
// @Tags LivingBill
// @Summary 删除LivingBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingBill true "删除LivingBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /livingBill/deleteLivingBill [delete]
func DeleteLivingBill(c *gin.Context) {
	var livingBill model.LivingBill
	_ = c.ShouldBindJSON(&livingBill)
	if err := service.DeleteLivingBill(livingBill); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLivingBillByIds 批量删除LivingBill
// @Tags LivingBill
// @Summary 批量删除LivingBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LivingBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /livingBill/deleteLivingBillByIds [delete]
func DeleteLivingBillByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteLivingBillByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLivingBill 更新LivingBill
// @Tags LivingBill
// @Summary 更新LivingBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingBill true "更新LivingBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /livingBill/updateLivingBill [put]
func UpdateLivingBill(c *gin.Context) {
	var livingBill model.LivingBill
	_ = c.ShouldBindJSON(&livingBill)
	if err := service.UpdateLivingBill(livingBill); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLivingBill 用id查询LivingBill
// @Tags LivingBill
// @Summary 用id查询LivingBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingBill true "用id查询LivingBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /livingBill/findLivingBill [get]
func FindLivingBill(c *gin.Context) {
	var livingBill model.LivingBill
	_ = c.ShouldBindQuery(&livingBill)
	if err, relivingBill := service.GetLivingBill(livingBill.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relivingBill": relivingBill}, c)
	}
}

// GetLivingBillList 分页获取LivingBill列表
// @Tags LivingBill
// @Summary 分页获取LivingBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LivingBillSearch true "分页获取LivingBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /livingBill/getLivingBillList [get]
func GetLivingBillList(c *gin.Context) {
	var pageInfo request.LivingBillSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetLivingBillInfoList(pageInfo); err != nil {
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

// CreateLivingCost 创建LivingCost
// @Tags LivingCost
// @Summary 创建LivingCost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingCost true "创建LivingCost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /livingCost/createLivingCost [post]
func CreateLivingCost(c *gin.Context) {
	var livingCost model.LivingCost
	_ = c.ShouldBindJSON(&livingCost)
	if err := service.CreateLivingCost(livingCost); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLivingCost 删除LivingCost
// @Tags LivingCost
// @Summary 删除LivingCost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingCost true "删除LivingCost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /livingCost/deleteLivingCost [delete]
func DeleteLivingCost(c *gin.Context) {
	var livingCost model.LivingCost
	_ = c.ShouldBindJSON(&livingCost)
	if err := service.DeleteLivingCost(livingCost); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLivingCostByIds 批量删除LivingCost
// @Tags LivingCost
// @Summary 批量删除LivingCost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LivingCost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /livingCost/deleteLivingCostByIds [delete]
func DeleteLivingCostByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteLivingCostByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLivingCost 更新LivingCost
// @Tags LivingCost
// @Summary 更新LivingCost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingCost true "更新LivingCost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /livingCost/updateLivingCost [put]
func UpdateLivingCost(c *gin.Context) {
	var livingCost model.LivingCost
	_ = c.ShouldBindJSON(&livingCost)
	if err := service.UpdateLivingCost(livingCost); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLivingCost 用id查询LivingCost
// @Tags LivingCost
// @Summary 用id查询LivingCost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingCost true "用id查询LivingCost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /livingCost/findLivingCost [get]
func FindLivingCost(c *gin.Context) {
	var livingCost model.LivingCost
	_ = c.ShouldBindQuery(&livingCost)
	if err, relivingCost := service.GetLivingCost(livingCost.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relivingCost": relivingCost}, c)
	}
}

// GetLivingCostList 分页获取LivingCost列表
// @Tags LivingCost
// @Summary 分页获取LivingCost列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LivingCostSearch true "分页获取LivingCost列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /livingCost/getLivingCostList [get]
func GetLivingCostList(c *gin.Context) {
	var pageInfo request.LivingCostSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetLivingCostInfoList(pageInfo); err != nil {
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
