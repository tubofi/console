
package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateTrialCourseRecord 创建TrialCourseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateTrialCourseRecord(trialCourseRecord model.TrialCourseRecord) (err error) {
	err = global.GVA_DB.Create(&trialCourseRecord).Error
	return err
}

// DeleteTrialCourseRecord 删除TrialCourseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteTrialCourseRecord(trialCourseRecord model.TrialCourseRecord) (err error) {
	err = global.GVA_DB.Delete(&trialCourseRecord).Error
	return err
}

// DeleteTrialCourseRecordByIds 批量删除TrialCourseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteTrialCourseRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.TrialCourseRecord{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTrialCourseRecord 更新TrialCourseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateTrialCourseRecord(trialCourseRecord model.TrialCourseRecord) (err error) {
	err = global.GVA_DB.Save(&trialCourseRecord).Error
	return err
}

// GetTrialCourseRecord 根据id获取TrialCourseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func GetTrialCourseRecord(id uint) (err error, trialCourseRecord model.TrialCourseRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&trialCourseRecord).Error
	return
}

// GetTrialCourseRecordInfoList 分页获取TrialCourseRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func GetTrialCourseRecordInfoList(info request.TrialCourseRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.TrialCourseRecord{})
	var trialCourseRecords []model.TrialCourseRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StudentName != "" {
		db = db.Where("`student_name` LIKE ?","%"+ info.StudentName+"%")
	}
	if info.CourseName != "" {
		db = db.Where("`course_name` LIKE ?","%"+ info.CourseName+"%")
	}
	if info.TeacherName != "" {
		db = db.Where("`teacher_name` LIKE ?","%"+ info.TeacherName+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("need_feedback desc").Order("time desc").Find(&trialCourseRecords).Error
	return err, trialCourseRecords, total
}

func GetAllTrialStudents() (err error, trialStudents []model.Student) {
	db := global.GVA_DB.Model(&model.Student{})
	err = db.Where("`is_entry` = ?", 0).Find(&trialStudents).Error
	if err != nil {
		return err, nil
	}

	return err, trialStudents
}

func FeedbackTrialCourseRecord(trialCourseRecord model.TrialCourseRecord) (err error) {
	trialCourseRecord.NeedFeedback = 0
	trialCourseRecord.Punctuality = trialCourseRecord.Punctuality / 5 * 5
	trialCourseRecord.Discipline  = trialCourseRecord.Discipline / 5 * 10
	trialCourseRecord.Concentration = trialCourseRecord.Concentration / 5 * 15
	trialCourseRecord.Innovation = trialCourseRecord.Innovation / 5 * 15
	trialCourseRecord.Logic = trialCourseRecord.Logic / 5 * 20
	trialCourseRecord.Mathematics = trialCourseRecord.Mathematics / 5 * 20
	trialCourseRecord.Complete = trialCourseRecord.Complete / 5 * 15

	trialCourseRecord.Total =
			trialCourseRecord.Punctuality +
			trialCourseRecord.Discipline +
			trialCourseRecord.Concentration +
			trialCourseRecord.Innovation +
			trialCourseRecord.Logic +
			trialCourseRecord.Mathematics +
			trialCourseRecord.Complete
	err = global.GVA_DB.Save(&trialCourseRecord).Error
	return err
}


