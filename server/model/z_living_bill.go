package model

import (
	"gorm.io/gorm"
	"time"
)

//生活账单
type LivingBill struct {
	gorm.Model
	Year			int				`json:"year" form:"year"`
	Month 			time.Month		`json:"month" form:"month"`
	All				float32			`json:"all" form:"all" gorm:"default:0.0"`
	Spends			[]*Spend		`json:"spends" form:"spends"`
	Comment 		string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

type Spend struct {
	gorm.Model
	CostID			uint			`json:"costId" form:"costId"`
	Time 			time.Time		`json:"time" form:"time"`
	Money			float32			`json:"money" form:"money" gorm:"default:0.0"`
	Comment 		string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}
