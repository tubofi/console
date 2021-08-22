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

	All					float64			`json:"all" form:"all"` //月工资
	//基础工资
	BaseWage 			float64 		`json:"baseWage" form:"baseWage"` //基础工资
	PostWage 			float64			`json:"postWage" form:"postWage"` //岗位工资
	Allowance 			float64 		`json:"allowance" form:"allowance"` //岗位补贴
	//教师岗位工资
	CourseFee 			float64 		`json:"courseFee" form:"courseFee"` //课时费
	TurnFee		 		float64			`json:"turnFee" form:"turnFee"`		//转化提成
	RenewFee	 		float64			`json:"renewFee" form:"renewFee"`	//续费提成
	//招生岗位工资
	NewStudent 			float64			`json:"newStudent" form:"newStudent"` //新单提成
	OldStudent 			float64 		`json:"oldStudent" form:"oldStudent"` //续费提成
	NewTeam 			float64			`json:"newTeam" form:"newTeam"` //新单团队提成
	OldTeam			  	float64 		`json:"oldTeam" form:"oldTeam"` //续费团队提成
	Bonus				float64			`json:"bonus" form:"bonus"` //奖金

	Comment 			string			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

func (Wage) TableName() string {
	return "wages"
}
