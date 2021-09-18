package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gorm.io/gorm"
)

// CreateCourseChange 创建CourseChange记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateCourseChange(courseChange model.CourseChange) (err error) {
	err = global.GVA_DB.Create(&courseChange).Error
	//加上课程次数变化
	err = global.GVA_DB.Model(&model.Student{}).Where(
		"name = ?", courseChange.StudentName).Update(
			"course_remain", gorm.Expr("course_remain + ?", courseChange.Amount)).Error
	return err
}

// DeleteCourseChange 删除CourseChange记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteCourseChange(courseChange model.CourseChange) (err error) {
	oldChange := model.CourseChange{}
	global.GVA_DB.Where("id = ?", courseChange.ID).First(&oldChange)
	err = global.GVA_DB.Delete(&courseChange).Error

	//减掉课程次数变化
	err = global.GVA_DB.Model(&model.Student{}).Where(
		"name = ?", oldChange.StudentName).Update(
		"course_remain", gorm.Expr("course_remain - ?", oldChange.Amount)).Error
	return err
}

// DeleteCourseChangeByIds 批量删除CourseChange记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteCourseChangeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.CourseChange{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCourseChange 更新CourseChange记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateCourseChange(courseChange model.CourseChange) (err error) {
	oldChange := model.CourseChange{}
	err = global.GVA_DB.Where("id = ?", courseChange.ID).First(&oldChange).Error

	err = global.GVA_DB.Save(&courseChange).Error

	//先减掉旧的，再加上新的
	err = global.GVA_DB.Model(&model.Student{}).Where(
		"name = ?", courseChange.StudentName).Update(
		"course_remain", gorm.Expr("course_remain - ? + ?", oldChange.Amount, courseChange.Amount)).Error
	return err
}

// GetCourseChange 根据id获取CourseChange记录
// Author [piexlmax](https://github.com/piexlmax)
func GetCourseChange(id uint) (err error, courseChange model.CourseChange) {
	err = global.GVA_DB.Where("id = ?", id).First(&courseChange).Error
	return
}

// GetCourseChangeInfoList 分页获取CourseChange记录
// Author [piexlmax](https://github.com/piexlmax)
func GetCourseChangeInfoList(info request.CourseChangeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.CourseChange{})
	var courseChanges []model.CourseChange
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StudentName != "" {
		db = db.Where("`student_name` LIKE ?","%"+ info.StudentName+"%")
	}
	if info.Type != 0 {
		db = db.Where("`type` = ?", info.Type)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&courseChanges).Error
	return err, courseChanges, total
}
