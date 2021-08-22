package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateBill 创建Bill记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateBill(bill model.Bill) (err error) {
	err = global.GVA_DB.Create(&bill).Error
	return err
}

// DeleteBill 删除Bill记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteBill(bill model.Bill) (err error) {
	err = global.GVA_DB.Delete(&bill).Error
	return err
}

// DeleteBillByIds 批量删除Bill记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteBillByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Bill{},"id in ?",ids.Ids).Error
	return err
}

// UpdateBill 更新Bill记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateBill(bill model.Bill) (err error) {
	err = global.GVA_DB.Save(&bill).Error
	return err
}

// GetBill 根据id获取Bill记录
// Author [piexlmax](https://github.com/piexlmax)
func GetBill(id uint) (err error, bill model.Bill) {
	err = global.GVA_DB.Where("id = ?", id).First(&bill).Error
	return
}

// GetBillInfoList 分页获取Bill记录
// Author [piexlmax](https://github.com/piexlmax)
func GetBillInfoList(info request.BillSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Bill{})
	var bills []model.Bill
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&bills).Error
	return err, bills, total
}
