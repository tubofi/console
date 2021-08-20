package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateTeacher 创建Teacher记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateTeacher(teacher model.Teacher) (err error) {
	err = global.GVA_DB.Create(&teacher).Error
	return err
}

// DeleteTeacher 删除Teacher记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteTeacher(teacher model.Teacher) (err error) {
	err = global.GVA_DB.Delete(&teacher).Error
	return err
}

// DeleteTeacherByIds 批量删除Teacher记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteTeacherByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Teacher{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTeacher 更新Teacher记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateTeacher(teacher model.Teacher) (err error) {
	err = global.GVA_DB.Save(&teacher).Error
	return err
}

// GetTeacher 根据id获取Teacher记录
// Author [piexlmax](https://github.com/piexlmax)
func GetTeacher(id uint) (err error, teacher model.Teacher) {
	err = global.GVA_DB.Where("id = ?", id).First(&teacher).Error
	return
}

// GetTeacherInfoList 分页获取Teacher记录
// Author [piexlmax](https://github.com/piexlmax)
func GetTeacherInfoList(info request.TeacherSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Teacher{})
	var teachers []model.Teacher
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&teachers).Error
	return err, teachers, total
}

func GetAllTeachers() (err error, teachers []model.Teacher) {
	db := global.GVA_DB.Model(&model.Teacher{})
	err = db.Find(&teachers).Error
	if err != nil {
		return err, nil
	}

	return err, teachers
}
