
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

func CreateClass(c *gin.Context) {
	var class model.Class
	_ = c.ShouldBindJSON(&class)
	if err := service.CreateClass(class); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func DeleteClass(c *gin.Context) {
	var class model.Class
	_ = c.ShouldBindJSON(&class)
	if err := service.DeleteClass(class); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func UpdateClass(c *gin.Context) {
	var class model.Class
	_ = c.ShouldBindJSON(&class)
	if err := service.UpdateClass(class); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func FindClass(c *gin.Context) {
	var class model.Class
	_ = c.ShouldBindQuery(&class)
	if err, reclass := service.GetClass(class.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reclass": reclass}, c)
	}
}

func GetClassList(c *gin.Context) {
	var pageInfo request.ClassSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetClassInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:" + err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func GetIdleStudents(c *gin.Context) {
	if err, students := service.GetIdleStudents(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败: " + err.Error(), c)
	} else {
		response.OkWithDetailed(students, "获取成功!", c)
	}
}

func GetProfileClass(c *gin.Context) {
	if err, profileClass := service.GetProfileClass(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败: " + err.Error(), c)
	} else {
		response.OkWithDetailed(profileClass, "获取成功!", c)
	}
}
