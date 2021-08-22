package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateWage 创建Wage记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateWage(wage model.Wage) (err error) {
	err = global.GVA_DB.Create(&wage).Error
	return err
}

// DeleteWage 删除Wage记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteWage(wage model.Wage) (err error) {
	err = global.GVA_DB.Delete(&wage).Error
	return err
}

// DeleteWageByIds 批量删除Wage记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteWageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Wage{},"id in ?",ids.Ids).Error
	return err
}

// UpdateWage 更新Wage记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateWage(wage model.Wage) (err error) {
	err = global.GVA_DB.Save(&wage).Error
	return err
}

// GetWage 根据id获取Wage记录
// Author [piexlmax](https://github.com/piexlmax)
func GetWage(id uint) (err error, wage model.Wage) {
	err = global.GVA_DB.Where("id = ?", id).First(&wage).Error
	return
}

// GetWageInfoList 分页获取Wage记录
// Author [piexlmax](https://github.com/piexlmax)
func GetWageInfoList(info request.WageSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Wage{})
	var wages []model.Wage
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&wages).Error
	return err, wages, total
}
