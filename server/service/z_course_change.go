
package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateCourseChange 创建CourseChange记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateCourseChange(courseChange model.CourseChange) (err error) {
	err = global.GVA_DB.Create(&courseChange).Error
	return err
}

// DeleteCourseChange 删除CourseChange记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteCourseChange(courseChange model.CourseChange) (err error) {
	err = global.GVA_DB.Delete(&courseChange).Error
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
	err = global.GVA_DB.Save(&courseChange).Error
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
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&courseChanges).Error
	return err, courseChanges, total
}
