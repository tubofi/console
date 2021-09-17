package model

import (
	"gorm.io/gorm"
)

//记录课程变化
type CourseChange struct {
	gorm.Model
	StudentName 	string		`json:"studentName" form:"studentName"`
	Amount 			int			`json:"amount" form:"amount"`
	Type 			int 		`json:"type" form:"type"`		//1报名；2续费；3推荐；4赠送；5积分；11修正

	Comment 		string		`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}


