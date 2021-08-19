package model

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	BillID				uint			`json:"billId" form:"billId"`
	StudentName			string			`json:"studentName" form:"studentName"`

	Pay 				float32			`json:"pay" form:"pay" gorm:"default:0.0"`
	PayTime 			time.Time		`json:"payTime" form:"payTime"`
	Pattern 			int 			`json:"pattern" form:"pattern" gorm:"default:0"`		//0新单 1续费
	Times 				int 			`json:"times" form:"times" gorm:"default:0"`			//课程次数

	Comment 			string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

type OtherIncome struct {
	gorm.Model
	BillID				uint			`json:"billId" form:"billId"`
	Money 				float32			`json:"money" form:"money" gorm:"default:0.0"`
	Time 				time.Time		`json:"time" form:"time"`

	Comment 			string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}
