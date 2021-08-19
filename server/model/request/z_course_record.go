package request

import "gin-vue-admin/model"

type CourseRecordSearch struct{
	model.CourseRecord
	PageInfo
}