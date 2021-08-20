package model

import (
	"gorm.io/gorm"
	"time"
)

type Wage struct {
	gorm.Model
	BillID				uint			`json:"billId" form:"billId"`
	TeacherName			string			`json:"teacherName" form:"teacherName"`
	Year				int				`json:"year" form:"year"`
	Month 				time.Month 		`json:"month" form:"month"`

	All					float32			`json:"all" form:"all" gorm:"default:0.0"` //月工资
	BaseWage 			float32 		`json:"baseWage" form:"baseWage" gorm:"default:1800.0"` //基础工资
	PostWage 			float32			`json:"postWage" form:"postWage" gorm:"default:1000.0"` //岗位工资
	Allowance 			float32 		`json:"allowance" form:"allowance" gorm:"default:200.0"` //岗位补贴
	CourseAll 			float32 		`json:"courseAll" form:"courseAll" gorm:"default:0.0"` //课时费
	NewStudent 			float32			`json:"newStudent" form:"newStudent" gorm:"default:0.0"` //新单提成
	OldStudent 			float32 		`json:"oldStudent" form:"oldStudent" gorm:"default:0.0"` //续费提成
	NewTeam 			float32			`json:"newTeam" form:"newTeam" gorm:"default:0.0"` //新单团队提成
	OldTeam			  	float32 		`json:"oldTeam" form:"oldTeam" gorm:"default:0.0"` //续费团队提成
	Bonus				float32			`json:"bonus" form:"bonus" gorm:"default:0.0"` //奖金

	Comment 			string			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}
