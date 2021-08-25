package request

import "gin-vue-admin/model"

type UploadSearch struct{
	model.Upload
	PageInfo
}