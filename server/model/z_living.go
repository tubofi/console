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
	IsClose 		int 			`json:"isClose" form:"isClose"`
	All				float64			`json:"all" form:"all"`
	LivingCosts		[]LivingCost	`json:"livingCosts" form:"livingCosts"`
	Comment 		string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

type LivingCost struct {
	gorm.Model
	LivingBillID	uint			`json:"livingBillID" form:"livingBillID"`
	Year			int				`json:"year" form:"year"`
	Month 			time.Month		`json:"month" form:"month"`
	Time 			time.Time		`json:"time" form:"time"`
	Money			float64			`json:"money" form:"money"`
	Type 			int 		 	`json:"type" form:"type"`				//1衣裤 2饮食 3居住 4交通 5生活用品 6家用电器 7电子设备 8化妆品 9药物 10宠物 11游玩 12书本文具 13其它
	Comment 		string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

func (LivingBill) TableName() string {
	return "living_bills"
}

func (LivingCost) TableName() string {
	return "living_costs"
}

