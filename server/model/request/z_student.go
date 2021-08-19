package request

import "gin-vue-admin/model"

type StudentSearch struct{
	model.Student
	PageInfo
}