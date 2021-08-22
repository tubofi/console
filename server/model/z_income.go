package model

import "gorm.io/gorm"

type Income struct {
	gorm.Model
	BillID				uint			`json:"billId" form:"billId"`
	Money 				float64			`json:"money" form:"money"`

	Comment 			string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

func (Income) TableName() string {
	return "incomes"
}