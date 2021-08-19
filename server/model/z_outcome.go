package model

import (
	"gorm.io/gorm"
	"time"
)

type Outcome struct {
	gorm.Model
	BillID				uint			`json:"billId" form:"billId"`
	Time				time.Time		`json:"time" form:"time"`

	Use 				int				`json:"use" form:"use"`
	Money				float32			`json:"money" form:"money" gorm:"default:0.0"`
	Comment 			string			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}
