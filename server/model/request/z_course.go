package request

import "gin-vue-admin/model"

type CourseSearch struct{
	model.ResCourse
	PageInfo
}