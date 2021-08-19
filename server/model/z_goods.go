package model

import (
	"gorm.io/gorm"
	"time"
)

//单个礼物
type Gift struct {
	gorm.Model
	GoodsID		uint		`json:"goodsID" form:"goodsID"`

	Name			string		`json:"name" form:"name"`		//物品的名字
	StudentName		string 		`json:"studentName" form:"studentName"`
	TakeTime 		time.Time	`json:"time" form:"time"`		//领取时间
	Comment 		string 		`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

//一类礼物
type Goods struct {
	gorm.Model
	SchoolID  		uint		`json:"schoolID" form:"schoolID"`
	Price			float32		`json:"price" form:"price"`
	Amount 			int			`json:"amount" form:"amount" gorm:"default:0"`
	PurchaseTime	time.Time	`json:"purchaseTime" form:"purchaseTime"` //购买时间

	Gifts			[]Gift		`json:"gifts" form:"gifts"`		//单项礼物

	Comment 		string		`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}
