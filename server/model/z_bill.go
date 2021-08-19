package model

import (
	"gorm.io/gorm"
	"time"
)

//月账单
type Bill struct {
	gorm.Model
	SchoolID 		uint			`json:"schoolId" form:"schoolId"`
	BillInfo 		BillInfo 		`gorm:"embedded;embeddedPrefix:bill_"`

	Payments 		[]Payment		`json:"payments" form:"payments"`
	OtherIncomes	[]OtherIncome	`json:"otherIncomes" form:"otherIncomes"`
	Outcomes 		[]Outcome		`json:"outcomes" form:"outcomes"`
	Wages			[]Wage			`json:"wages" form:"wages"`

	Comment 		string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

type BillInfo struct {
	Year			int				`json:"year" form:"year"`
	Month 			time.Month		`json:"month" form:"month"`
	AllIn 			float32			`json:"allIn" form:"allIn" gorm:"default:0.0"`
	AllOut 			float32			`json:"allOut" form:"allOut" gorm:"default:0.0"`
	Profit 			float32			`json:"profit" form:"profit" gorm:"default:0.0"`
}


