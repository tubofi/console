package model

import (
	"gorm.io/gorm"
	"time"
)

//月账单
type Bill struct {
	gorm.Model
	SchoolID 		uint			`json:"schoolId" form:"schoolId"`
	Year			int				`json:"year" form:"year"`
	Month 			time.Month		`json:"month" form:"month"`
	IsFinished 		int 			`json:"isFinished" form:"isFinished"`			//0是未结算	1为已结算
	AllIn 			float64			`json:"allIn" form:"allIn"`
	AllOut 			float64			`json:"allOut" form:"allOut"`
	Profit 			float64			`json:"profit" form:"profit"`

	Payments 		[]Payment		`json:"payments" form:"payments"`
	Incomes			[]Income		`json:"incomes" form:"incomes"`
	Outcomes 		[]Outcome		`json:"outcomes" form:"outcomes"`
	Wages			[]Wage			`json:"wages" form:"wages"`

	Comment 		string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}
func (Bill) TableName() string {
	return "bills"
}










