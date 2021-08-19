
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

func CreatePotential(c *gin.Context) {
	var potential model.Potential
	_ = c.ShouldBindJSON(&potential)
	if err := service.CreatePotential(potential); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("创建成功!", c)
	}
}

func DeletePotential(c *gin.Context) {
	var potential model.Potential
	_ = c.ShouldBindJSON(&potential)
	if err := service.DeletePotential(potential); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功!", c)
	}
}

func UpdatePotential(c *gin.Context) {
	var potential model.Potential
	_ = c.ShouldBindJSON(&potential)
	if err := service.UpdatePotential(potential); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func FindPotential(c *gin.Context) {
	var potential model.Potential
	_ = c.ShouldBindQuery(&potential)
	if err, repotential := service.GetPotential(potential.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败: " + err.Error(), c)
	} else {
		response.OkWithData(gin.H{"repotential": repotential}, c)
	}
}

func GetPotentialList(c *gin.Context) {
	var pageInfo request.PotentialSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPotentialInfoList(pageInfo); err != nil {
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

func PotentialToTrial(c *gin.Context) {
	var potential model.Potential
	_ = c.ShouldBindJSON(&potential)
	if err := service.PotentialToTrial(potential); err != nil {
		global.GVA_LOG.Error("试听失败!", zap.Any("err", err))
		response.FailWithMessage("试听失败: "  + err.Error(), c)
	} else {
		response.OkWithMessage("试听成功", c)
	}
}

func UpgradePotentialLevel(c *gin.Context) {
	var potential model.Potential
	_ = c.ShouldBindJSON(&potential)
	if err := service.UpgradePotentialLevel(potential.ID); err != nil {
		global.GVA_LOG.Error("升级失败!", zap.Any("err", err))
		response.FailWithMessage("升级失败: "  + err.Error(), c)
	} else {
		response.OkWithMessage("升级成功!", c)
	}
}

func DegradePotentialLevel(c *gin.Context) {
	var potential model.Potential
	_ = c.ShouldBindJSON(&potential)
	if err := service.DegradePotentialLevel(potential.ID); err != nil {
		global.GVA_LOG.Error("降级失败!", zap.Any("err", err))
		response.FailWithMessage("降级失败: "  + err.Error(), c)
	} else {
		response.OkWithMessage("降级成功!", c)
	}
}
