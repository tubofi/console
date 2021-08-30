package model

import (
	"gin-vue-admin/global"
)

type ExaFileUploadAndDownload struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"comment:文件名"` // 文件名
	Url  string `json:"url" form:"url" gorm:"comment:文件地址"` // 文件地址
	Tag  string `json:"tag" form:"tag" gorm:"comment:文件标签"` // 文件标签
	Key  string `json:"key" form:"key" gorm:"comment:编号"`   // 编号
}
