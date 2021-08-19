package model

import (
	"gorm.io/gorm"
	"time"
)

//记录课程变化
type CourseChange struct {
	gorm.Model
	StudentName 	string		`json:"studentName" form:"studentName"`
	Time 			time.Time	`json:"time" form:"time"`
	Amount 			int			`json:"amount" form:"amount"`
	Type 			int 		`json:"type" form:"type"`		//0新缴费；1续费；2推荐；3转介绍；4赠送；5积分

	Comment 		string		`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

type CreditChange struct {
	gorm.Model
	StudentName 	string		`json:"studentName" form:"studentName"`
	Time 			time.Time	`json:"time" form:"time"`
	Amount 			int			`json:"amount" form:"amount"`
	Type 			int 		`json:"type" form:"type"`		//0签到；1推荐；2转介绍；3其它

	Comment 		string		`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}
