package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateOutcome 创建Outcome记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateOutcome(outcome model.Outcome) (err error) {
	err = global.GVA_DB.Create(&outcome).Error
	return err
}

// DeleteOutcome 删除Outcome记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteOutcome(outcome model.Outcome) (err error) {
	err = global.GVA_DB.Delete(&outcome).Error
	return err
}

// DeleteOutcomeByIds 批量删除Outcome记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteOutcomeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Outcome{},"id in ?",ids.Ids).Error
	return err
}

// UpdateOutcome 更新Outcome记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateOutcome(outcome model.Outcome) (err error) {
	err = global.GVA_DB.Save(&outcome).Error
	return err
}

// GetOutcome 根据id获取Outcome记录
// Author [piexlmax](https://github.com/piexlmax)
func GetOutcome(id uint) (err error, outcome model.Outcome) {
	err = global.GVA_DB.Where("id = ?", id).First(&outcome).Error
	return
}

// GetOutcomeInfoList 分页获取Outcome记录
// Author [piexlmax](https://github.com/piexlmax)
func GetOutcomeInfoList(info request.OutcomeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Outcome{})
	var outcomes []model.Outcome
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&outcomes).Error
	return err, outcomes, total
}
