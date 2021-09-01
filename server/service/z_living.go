
package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateLivingBill 创建LivingBill记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateLivingBill(livingBill model.LivingBill) (err error) {
	err = global.GVA_DB.Create(&livingBill).Error
	return err
}

// DeleteLivingBill 删除LivingBill记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteLivingBill(livingBill model.LivingBill) (err error) {
	err = global.GVA_DB.Delete(&livingBill).Error
	return err
}

// DeleteLivingBillByIds 批量删除LivingBill记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteLivingBillByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.LivingBill{},"id in ?",ids.Ids).Error
	return err
}

// UpdateLivingBill 更新LivingBill记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateLivingBill(livingBill model.LivingBill) (err error) {
	err = global.GVA_DB.Save(&livingBill).Error
	return err
}

// GetLivingBill 根据id获取LivingBill记录
// Author [piexlmax](https://github.com/piexlmax)
func GetLivingBill(id uint) (err error, livingBill model.LivingBill) {
	err = global.GVA_DB.Where("id = ?", id).First(&livingBill).Error
	return
}

// GetLivingBillInfoList 分页获取LivingBill记录
// Author [piexlmax](https://github.com/piexlmax)
func GetLivingBillInfoList(info request.LivingBillSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.LivingBill{})
	var livingBills []model.LivingBill
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Year != 0 {
		db = db.Where("`year` = ?",info.Year)
	}
	if info.Month != 0 {
		db = db.Where("`month` = ?",info.Month)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("month desc").Order("year desc").Find(&livingBills).Error
	return err, livingBills, total
}

// CreateLivingCost 创建LivingCost记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateLivingCost(livingCost model.LivingCost) (err error) {
	bill := model.LivingBill{}
	if global.GVA_DB.Where(`year = ? AND month = ?`, livingCost.Time.Year(), livingCost.Time.Month()).First(&bill).RowsAffected == 1 {
		bill.All += livingCost.Money
		global.GVA_DB.Save(&bill)
	} else {
		bill.Year = livingCost.Time.Year()
		bill.Month = livingCost.Time.Month()
		bill.All = livingCost.Money
		global.GVA_DB.Create(&bill)
	}
	livingCost.LivingBillID = bill.ID
	err = global.GVA_DB.Create(&livingCost).Error
	return err
}

// DeleteLivingCost 删除LivingCost记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteLivingCost(livingCost model.LivingCost) (err error) {
	bill := model.LivingBill{}
	if global.GVA_DB.Where(`year = ? AND month = ?`, livingCost.Time.Year(), livingCost.Time.Month()).First(&bill).RowsAffected == 1 {
		bill.All = bill.All - livingCost.Money
		global.GVA_DB.Save(&bill)
	} else {
		return errors.New("找不到对应的月账单")
	}
	err = global.GVA_DB.Delete(&livingCost).Error
	return err
}

// DeleteLivingCostByIds 批量删除LivingCost记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteLivingCostByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.LivingCost{},"id in ?",ids.Ids).Error
	return err
}

// UpdateLivingCost 更新LivingCost记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateLivingCost(livingCost model.LivingCost) (err error) {
	oldLivingCost := model.LivingCost{}
	global.GVA_DB.Where(`id = ?`, livingCost.ID).First(&oldLivingCost)

	oldBill := model.LivingBill{}
	if global.GVA_DB.Where(`year = ? AND month = ?`, oldLivingCost.Time.Year(), oldLivingCost.Time.Month()).First(&oldBill).RowsAffected == 1 {
		oldBill.All = oldBill.All - oldLivingCost.Money
		global.GVA_DB.Save(&oldBill)
	} else {
		return errors.New("找不到对应的月账单")
	}

	newBill := model.LivingBill{}
	if global.GVA_DB.Where(`year = ? AND month = ?`, livingCost.Time.Year(), livingCost.Time.Month()).First(&newBill).RowsAffected == 1 {
		newBill.All = newBill.All + livingCost.Money
		global.GVA_DB.Save(&newBill)
	} else {
		return errors.New("找不到对应的月账单")
	}

	livingCost.LivingBillID = newBill.ID
	err = global.GVA_DB.Save(&livingCost).Error
	return err
}

// GetLivingCost 根据id获取LivingCost记录
// Author [piexlmax](https://github.com/piexlmax)
func GetLivingCost(id uint) (err error, livingCost model.LivingCost) {
	err = global.GVA_DB.Where("id = ?", id).First(&livingCost).Error
	return
}

// GetLivingCostInfoList 分页获取LivingCost记录
// Author [piexlmax](https://github.com/piexlmax)
func GetLivingCostInfoList(info request.LivingCostSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.LivingCost{})
	var livingCosts []model.LivingCost
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Year != 0 {
		db = db.Where("`year` = ?",info.Year)
	}
	if info.Month != 0 {
		db = db.Where("`month` = ?",info.Month)
	}
	if info.Type != 0 {
		db = db.Where("`type` = ?",info.Type)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("time desc").Find(&livingCosts).Error
	return err, livingCosts, total
}

