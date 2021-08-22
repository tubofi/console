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

// CreatePayment 创建Payment
// @Tags Payment
// @Summary 创建Payment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Payment true "创建Payment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payment/createPayment [post]
func CreatePayment(c *gin.Context) {
	var payment model.Payment
	_ = c.ShouldBindJSON(&payment)
	if err := service.CreatePayment(payment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePayment 删除Payment
// @Tags Payment
// @Summary 删除Payment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Payment true "删除Payment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payment/deletePayment [delete]
func DeletePayment(c *gin.Context) {
	var payment model.Payment
	_ = c.ShouldBindJSON(&payment)
	if err := service.DeletePayment(payment); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePaymentByIds 批量删除Payment
// @Tags Payment
// @Summary 批量删除Payment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Payment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /payment/deletePaymentByIds [delete]
func DeletePaymentByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePaymentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePayment 更新Payment
// @Tags Payment
// @Summary 更新Payment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Payment true "更新Payment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payment/updatePayment [put]
func UpdatePayment(c *gin.Context) {
	var payment model.Payment
	_ = c.ShouldBindJSON(&payment)
	if err := service.UpdatePayment(payment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPayment 用id查询Payment
// @Tags Payment
// @Summary 用id查询Payment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Payment true "用id查询Payment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payment/findPayment [get]
func FindPayment(c *gin.Context) {
	var payment model.Payment
	_ = c.ShouldBindQuery(&payment)
	if err, repayment := service.GetPayment(payment.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败: " + err.Error(), c)
	} else {
		response.OkWithData(gin.H{"repayment": repayment}, c)
	}
}

// GetPaymentList 分页获取Payment列表
// @Tags Payment
// @Summary 分页获取Payment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PaymentSearch true "分页获取Payment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payment/getPaymentList [get]
func GetPaymentList(c *gin.Context) {
	var pageInfo request.PaymentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPaymentInfoList(pageInfo); err != nil {
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
