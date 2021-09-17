package model

import (
	"gorm.io/gorm"
	"time"
)

//记录积分变化
type CreditChange struct {
	gorm.Model
	StudentName 	string		`json:"studentName" form:"studentName"`
	Time 			time.Time	`json:"time" form:"time"`
	Amount 			int			`json:"amount" form:"amount"`
	Type 			int 		`json:"type" form:"type"`		//0签到；1推荐；2转介绍；3其它

	Comment 		string		`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

func (CourseChange) TableName() string {
	return "course_changes"
}