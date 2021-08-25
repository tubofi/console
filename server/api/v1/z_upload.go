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

func GetTmpSecret (c *gin.Context) {
	if err, res := service.GetTmpSecret(""); err != nil {
		global.GVA_LOG.Error("获取临时密钥失败!", zap.Any("err", err))
		response.FailWithMessage("获取临时密钥失败: " + err.Error(), c)
	} else {
		response.OkWithData(res, c)
	}
}

// CreateUpload 创建Upload
// @Tags Upload
// @Summary 创建Upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Upload true "创建Upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /upload/createUpload [post]
func CreateUpload(c *gin.Context) {
	var upload model.Upload
	_ = c.ShouldBindJSON(&upload)
	if err := service.CreateUpload(upload); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUpload 删除Upload
// @Tags Upload
// @Summary 删除Upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Upload true "删除Upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /upload/deleteUpload [delete]
func DeleteUpload(c *gin.Context) {
	var upload model.Upload
	_ = c.ShouldBindJSON(&upload)
	if err := service.DeleteUpload(upload); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUploadByIds 批量删除Upload
// @Tags Upload
// @Summary 批量删除Upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /upload/deleteUploadByIds [delete]
func DeleteUploadByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteUploadByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUpload 更新Upload
// @Tags Upload
// @Summary 更新Upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Upload true "更新Upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /upload/updateUpload [put]
func UpdateUpload(c *gin.Context) {
	var upload model.Upload
	_ = c.ShouldBindJSON(&upload)
	if err := service.UpdateUpload(upload); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUpload 用id查询Upload
// @Tags Upload
// @Summary 用id查询Upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Upload true "用id查询Upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /upload/findUpload [get]
func FindUpload(c *gin.Context) {
	var upload model.Upload
	_ = c.ShouldBindQuery(&upload)
	if err, reupload := service.GetUpload(upload.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reupload": reupload}, c)
	}
}

// GetUploadList 分页获取Upload列表
// @Tags Upload
// @Summary 分页获取Upload列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UploadSearch true "分页获取Upload列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /upload/getUploadList [get]
func GetUploadList(c *gin.Context) {
	var pageInfo request.UploadSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetUploadInfoList(pageInfo); err != nil {
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

