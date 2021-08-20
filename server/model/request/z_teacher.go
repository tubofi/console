package request

import "gin-vue-admin/model"

type TeacherSearch struct{
	model.Teacher
	PageInfo
}