package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateIncome 创建Income记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateIncome(income model.Income) (err error) {
	err = global.GVA_DB.Create(&income).Error
	return err
}

// DeleteIncome 删除Income记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteIncome(income model.Income) (err error) {
	err = global.GVA_DB.Delete(&income).Error
	return err
}

// DeleteIncomeByIds 批量删除Income记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteIncomeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Income{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIncome 更新Income记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateIncome(income model.Income) (err error) {
	err = global.GVA_DB.Save(&income).Error
	return err
}

// GetIncome 根据id获取Income记录
// Author [piexlmax](https://github.com/piexlmax)
func GetIncome(id uint) (err error, income model.Income) {
	err = global.GVA_DB.Where("id = ?", id).First(&income).Error
	return
}

// GetIncomeInfoList 分页获取Income记录
// Author [piexlmax](https://github.com/piexlmax)
func GetIncomeInfoList(info request.IncomeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Income{})
	var incomes []model.Income
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&incomes).Error
	return err, incomes, total
}
