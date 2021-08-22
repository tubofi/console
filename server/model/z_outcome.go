package model

import "gorm.io/gorm"

type Outcome struct {
	gorm.Model
	BillID				uint			`json:"billId" form:"billId"`

	Money				float64			`json:"money" form:"money"`
	Comment 			string			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

func (Outcome) TableName() string {
	return "outcomes"
}