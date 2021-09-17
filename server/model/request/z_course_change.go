package request

import "gin-vue-admin/model"

type CourseChangeSearch struct{
	model.CourseChange
	PageInfo
}