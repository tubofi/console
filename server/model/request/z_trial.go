package request

import "gin-vue-admin/model"

type TrialCourseRecordSearch struct{
	model.TrialCourseRecord
	PageInfo
}