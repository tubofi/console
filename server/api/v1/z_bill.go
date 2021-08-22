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

// CreateBill 创建Bill
// @Tags Bill
// @Summary 创建Bill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Bill true "创建Bill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bill/createBill [post]
func CreateBill(c *gin.Context) {
	var bill model.Bill
	_ = c.ShouldBindJSON(&bill)
	if err := service.CreateBill(bill); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBill 删除Bill
// @Tags Bill
// @Summary 删除Bill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Bill true "删除Bill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bill/deleteBill [delete]
func DeleteBill(c *gin.Context) {
	var bill model.Bill
	_ = c.ShouldBindJSON(&bill)
	if err := service.DeleteBill(bill); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBillByIds 批量删除Bill
// @Tags Bill
// @Summary 批量删除Bill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Bill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bill/deleteBillByIds [delete]
func DeleteBillByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteBillByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBill 更新Bill
// @Tags Bill
// @Summary 更新Bill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Bill true "更新Bill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bill/updateBill [put]
func UpdateBill(c *gin.Context) {
	var bill model.Bill
	_ = c.ShouldBindJSON(&bill)
	if err := service.UpdateBill(bill); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBill 用id查询Bill
// @Tags Bill
// @Summary 用id查询Bill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Bill true "用id查询Bill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bill/findBill [get]
func FindBill(c *gin.Context) {
	var bill model.Bill
	_ = c.ShouldBindQuery(&bill)
	if err, rebill := service.GetBill(bill.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebill": rebill}, c)
	}
}

// GetBillList 分页获取Bill列表
// @Tags Bill
// @Summary 分页获取Bill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BillSearch true "分页获取Bill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bill/getBillList [get]
func GetBillList(c *gin.Context) {
	var pageInfo request.BillSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetBillInfoList(pageInfo); err != nil {
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
