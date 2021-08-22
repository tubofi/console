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

//重新计算月账单
func UpdateBill(bill model.Bill) (err error) {
	newBill := model.Bill{}
	global.GVA_DB.Where(
		`ID = ?`, bill.ID).Preload(
			"Payments").Preload(
				"Incomes").Preload(
				"Outcomes").Preload(
					"Wages").First(&newBill)

	//计算收入
	newBill.PaymentIncome = 0
	for _, pay := range newBill.Payments {
		newBill.PaymentIncome += pay.Money
	}
	newBill.OtherIncome = 0
	for _, pay := range newBill.Incomes {
		newBill.OtherIncome += pay.Money
	}
	newBill.AllIncome = newBill.PaymentIncome + newBill.OtherIncome

	//计算支出
	newBill.WageOutcome = 0
	for _, pay := range newBill.Wages {
		newBill.WageOutcome += pay.All
	}
	newBill.OtherOutcome = 0
	for _, pay := range newBill.Outcomes {
		newBill.OtherOutcome += pay.Money
	}
	newBill.AllOutcome = newBill.WageOutcome + newBill.OtherOutcome

	//计算利润
	newBill.Profit = newBill.AllIncome - newBill.AllOutcome
	newBill.IsFinished = 1

	err = global.GVA_DB.Save(&newBill).Error
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

func GetAllStudents() (err error, students []model.Student) {
	db := global.GVA_DB.Model(&model.Student{})
	err = db.Where("`is_entry` = ?", 1).Find(&students).Error
	if err != nil {
		return err, nil
	}

	return err, students
}

