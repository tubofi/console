
package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"strings"
)

// CreateCourseRecord 创建CourseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateCourseRecord(CourseRecord model.CourseRecord) (err error) {
	err = global.GVA_DB.Create(&CourseRecord).Error
	return err
}

// DeleteCourseRecord 删除CourseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteCourseRecord(CourseRecord model.CourseRecord) (err error) {
	err = global.GVA_DB.Delete(&CourseRecord).Error
	return err
}

// DeleteCourseRecordByIds 批量删除CourseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteCourseRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.CourseRecord{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCourseRecord 更新CourseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateCourseRecord(CourseRecord model.CourseRecord) (err error) {
	err = global.GVA_DB.Save(&CourseRecord).Error
	return err
}

// GetCourseRecord 根据id获取CourseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func GetCourseRecord(id uint) (err error, CourseRecord model.CourseRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&CourseRecord).Error
	CourseRecord = RecordToRate(CourseRecord)
	return
}

// GetCourseRecordInfoList 分页获取CourseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func GetCourseRecordInfoList(info request.CourseRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.CourseRecord{})
	var CourseRecords []model.CourseRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StudentName != "" {
		db = db.Where("`student_name` LIKE ?","%"+ info.StudentName+"%")
	}
	if info.TeacherName != "" {
		db = db.Where("`teacher_name` LIKE ?","%"+ info.TeacherName+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order(
		"need_cram desc").Order(
		"need_feedback desc").Order(
		"attend_time desc").Find(&CourseRecords).Error
	return err, CourseRecords, total
}

func GetCramCourseRecordInfoList(info request.CourseRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.CourseRecord{})
	var CourseRecords []model.CourseRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StudentName != "" {
		db = db.Where("`student_name` LIKE ?","%"+ info.StudentName+"%")
	}
	if info.TeacherName != "" {
		db = db.Where("`teacher_name` LIKE ?","%"+ info.TeacherName+"%")
	}
	if info.CourseName != "" {
		db = db.Where("`course_name` LIKE ?","%"+ info.CourseName+"%")
	}
	//err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("need_cram desc").Order("attend_time desc").Where("`is_absentee` = ?",1).Find(&CourseRecords).Error
	total = int64(len(CourseRecords))		//正常是要获取表总长度，此处以is_absentee=1的数量代替
	return err, CourseRecords, total
}

func CramCourseRecord(CourseRecord model.CourseRecord) (err error) {
	CourseRecord.NeedCram = 0
	err = global.GVA_DB.Save(&CourseRecord).Error

/*	//补课后减少学员课次
	student := model.Student{}
	global.GVA_DB.Where(`name = ?`, CourseRecord.StudentName).First(&student)
	student.CourseRemain = student.CourseRemain - 1
	global.GVA_DB.Save(&student)*/

	//创建课程次数变更记录
	description, courseType := "", ""
	//判断课程类型！垃圾方法
	switch string(CourseRecord.CourseContent[0]) {
	case "P":
		courseType = "python"
	case "S":
		courseType = "scratch"
	case "O":
		courseType = "osmo"
	case "J":
		courseType = "scratchJR"
	case "C":
		courseType = "cpp"
	default:
		courseType = "course"
	}
	//判断课程名称中是否存在书名号
	if strings.Index(CourseRecord.CourseName, "《") != -1 && strings.Index(CourseRecord.CourseName, "》") != -1 {
		description = courseType + ":" + CourseRecord.CourseName
	} else {
		description = courseType + ":《" + CourseRecord.CourseName + "》"
	}
	courseChange := model.CourseChange{
		StudentName: 	CourseRecord.StudentName,
		Amount:  		-1,			//次数减少1
		Type:  			-1,			//类型为消课
		Description:  	description,
	}
	err = CreateCourseChange(courseChange)

	return err
}

func IgnoreCramCourseRecord(CourseRecord model.CourseRecord) (err error) {
	global.GVA_DB.First(&CourseRecord)
	CourseRecord.NeedCram = 0
	CourseRecord.NeedFeedback = 0
	err = global.GVA_DB.Save(&CourseRecord).Error
	return err
}

func GetFeedbackCourseRecordInfoList(info request.CourseRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.CourseRecord{})
	var CourseRecords []model.CourseRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StudentName != "" {
		db = db.Where("`student_name` LIKE ?","%"+ info.StudentName+"%")
	}
	if info.TeacherName != "" {
		db = db.Where("`teacher_name` LIKE ?","%"+ info.TeacherName+"%")
	}
	if info.CourseName != "" {
		db = db.Where("`course_name` LIKE ?","%"+ info.CourseName+"%")
	}
	//err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("need_feedback desc").Order("attend_time desc").Where("`need_cram` = ?", 0).Find(&CourseRecords).Error
	total = int64(len(CourseRecords))		//已经被忽略，need_cram=0的不加入统计了
	return err, CourseRecords, total
}

func FeedbackCourseRecord(CourseRecord model.CourseRecord) (err error) {
	CourseRecord.NeedFeedback = 0
	CourseRecord = RateToRecord(CourseRecord)
	err = global.GVA_DB.Save(&CourseRecord).Error
	return err
}

func RateToRecord(CourseRecord model.CourseRecord) (res model.CourseRecord){
	CourseRecord.Punctuality = CourseRecord.Punctuality / 5 * 5
	CourseRecord.Discipline  = CourseRecord.Discipline / 5 * 10
	CourseRecord.Concentration = CourseRecord.Concentration / 5 * 15
	CourseRecord.Innovation = CourseRecord.Innovation / 5 * 15
	CourseRecord.Logic = CourseRecord.Logic / 5 * 20
	CourseRecord.Mathematics = CourseRecord.Mathematics / 5 * 20
	CourseRecord.Complete = CourseRecord.Complete / 5 * 15

	CourseRecord.Total =
		CourseRecord.Punctuality +
			CourseRecord.Discipline +
			CourseRecord.Concentration +
			CourseRecord.Innovation +
			CourseRecord.Logic +
			CourseRecord.Mathematics +
			CourseRecord.Complete

	res = CourseRecord
	return res
}

func RecordToRate(CourseRecord model.CourseRecord) (res model.CourseRecord) {
	CourseRecord.Punctuality = CourseRecord.Punctuality * 5 / 5
	CourseRecord.Discipline  = CourseRecord.Discipline * 5 / 10
	CourseRecord.Concentration = CourseRecord.Concentration * 5 / 15
	CourseRecord.Innovation = CourseRecord.Innovation * 5 / 15
	CourseRecord.Logic = CourseRecord.Logic * 5 / 20
	CourseRecord.Mathematics = CourseRecord.Mathematics * 5 / 20
	CourseRecord.Complete = CourseRecord.Complete * 5 / 15

	CourseRecord.Total =
		CourseRecord.Punctuality +
			CourseRecord.Discipline +
			CourseRecord.Concentration +
			CourseRecord.Innovation +
			CourseRecord.Logic +
			CourseRecord.Mathematics +
			CourseRecord.Complete
	res = CourseRecord
	return
}

