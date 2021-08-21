package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	BillID				uint			`json:"billId" form:"billId"`
	StudentName			string			`json:"studentName" form:"studentName"`

	Money 				float64			`json:"money" form:"money"`
	Pattern 			int 			`json:"pattern" form:"pattern"`		//1新单 2续费 3试听
	Times 				int 			`json:"times" form:"times"`			//课程次数

	Comment 			string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

type OtherIncome struct {
	gorm.Model
	BillID				uint			`json:"billId" form:"billId"`
	Money 				float64			`json:"money" form:"money"`

	Comment 			string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}
