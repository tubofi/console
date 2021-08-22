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

	PaymentIncome   float64			`json:"paymentIncome" form:"paymentIncome"`
	OtherIncome 	float64			`json:"otherIncome" form:"otherIncome"`
	AllIncome 		float64			`json:"allIncome" form:"allIncome"`

	WageOutcome		float64			`json:"wageOutcome" form:"wageOutcome"`
	OtherOutcome 	float64			`json:"otherOutcome" form:"otherOutcome"`
	AllOutcome 		float64			`json:"allOutcome" form:"allOutcome"`

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










