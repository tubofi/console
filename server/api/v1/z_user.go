package v1

import (
	"gin-vue-admin/global"
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
