package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreatePayment 创建Payment记录
// Author [piexlmax](https://github.com/piexlmax)
func CreatePayment(payment model.Payment) (err error) {
	err = global.GVA_DB.Create(&payment).Error
	return err
}

// DeletePayment 删除Payment记录
// Author [piexlmax](https://github.com/piexlmax)
func DeletePayment(payment model.Payment) (err error) {
	err = global.GVA_DB.Delete(&payment).Error
	return err
}

// DeletePaymentByIds 批量删除Payment记录
// Author [piexlmax](https://github.com/piexlmax)
func DeletePaymentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Payment{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePayment 更新Payment记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdatePayment(payment model.Payment) (err error) {
	err = global.GVA_DB.Save(&payment).Error
	return err
}

// GetPayment 根据id获取Payment记录
// Author [piexlmax](https://github.com/piexlmax)
func GetPayment(id uint) (err error, payment model.Payment) {
	err = global.GVA_DB.Where("id = ?", id).First(&payment).Error
	return
}

// GetPaymentInfoList 分页获取Payment记录
// Author [piexlmax](https://github.com/piexlmax)
func GetPaymentInfoList(info request.PaymentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Payment{})
	var payments []model.Payment
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order(`created_at desc`).Find(&payments).Error
	return err, payments, total
}
