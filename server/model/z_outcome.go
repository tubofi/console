package model

import (
	"gorm.io/gorm"
)

type Outcome struct {
	gorm.Model
	BillID				uint			`json:"billId" form:"billId"`

	Money				float32			`json:"money" form:"money"`
	Comment 			string			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}
