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

func CreateStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindJSON(&student)
	if err := service.CreateStudent(student); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败：" + err.Error(), c)
	} else {
		response.OkWithMessage("创建成功!", c)
	}
}

func TurnStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindJSON(&student)
	if err := service.TurnStudent(student.ID); err != nil {
		global.GVA_LOG.Error("转正失败!", zap.Any("err", err))
		response.FailWithMessage("转正失败：" + err.Error(), c)
	} else {
		response.OkWithMessage("转正成功!", c)
	}
}

func DeleteStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindJSON(&student)
	if err := service.DeleteStudent(student); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败：" + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功!", c)
	}
}

func UpdateStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindJSON(&student)
	if err := service.UpdateStudent(student); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("更新成功!", c)
	}
}

func FindStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindQuery(&student)
	if err, restudent := service.GetStudent(student.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败: " + err.Error(), c)
	} else {
		response.OkWithData(gin.H{"restudent": restudent}, c)
	}
}

func GetStudentList(c *gin.Context) {
	var pageInfo request.StudentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetStudentInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败: " + err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功!", c)
	}
}

func GetTrialStudentList(c *gin.Context) {
	var pageInfo request.StudentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTrialStudentInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败: " + err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功!", c)
	}
}





