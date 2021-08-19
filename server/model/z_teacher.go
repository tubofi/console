package model

import (
	"gorm.io/gorm"
	"time"
)

type Teacher struct {
	gorm.Model
	SchoolID    		uint			`json:"schoolID" form:"schoolID"`
	TeacherInfo 		TeacherInfo 	`gorm:"embedded"`
	WageInfo    		WageInfo    	`gorm:"embedded"`

	//教师岗位
	Classes 			[]Class			`json:"classes" form:"classes" gorm:"foreignKey:TeacherName;references:Name"`
	CourseRecords		[]CourseRecord 	`json:"courseRecords" form:"courseRecords" gorm:"foreignKey:TeacherName;references:Name"`
	Students 			[]Student		`json:"students" form:"students" gorm:"foreignKey:TeacherName;references:Name"`

	//招生岗位
	Potentials			[]Potential		`json:"potentials" form:"potentials" gorm:"foreignKey:ManagerName;references:Name"`

	//工资
	Wages				[]Wage			`json:"wages" form:"wages" gorm:"foreignKey:TeacherName;references:Name"`

	Comment 			string			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

type TeacherInfo struct {
	Name 				string			`json:"name" form:"name"`
	Sex 				int				`json:"sex" form:"sex"` //女/男
	Address 			string 			`json:"address" form:"address"` //居住地址
	IdentityNumber 		string 			`json:"identityNumber" form:"identityNumber"` //身份证号
	Phone 				string			`json:"phone" form:"phone"`
	EntryTime 			time.Time 		`json:"entryTime" form:"entryTime"` //入职时间
}

type WageInfo struct {
	BaseWage 			float32 		`json:"baseWage" form:"baseWage" gorm:"default:1800.0"` //基础工资
	PostWage 			float32			`json:"postWage" form:"postWage" gorm:"default:1000.0"` //岗位工资
	Allowance 			float32 		`json:"allowance" form:"allowance" gorm:"default:200.0"` //岗位补贴
	CoursePer 			float32 		`json:"coursePer" form:"coursePer" gorm:"default:10.0"` //课时费
	NewPercent 			float32			`json:"newPercent" form:"newPercent" gorm:"default:0.1"` //新单提成
	OldPercent 			float32 		`json:"oldPercent" form:"oldPercent" gorm:"default:0.05"` //续费提成
	NewTeamPercent 		float32			`json:"newTeamPercent" form:"newTeamPercent" gorm:"default:0.05"` //新单团队提成
	OldTeamPercent  	float32 		`json:"oldTeamPercent" form:"oldTeamPercent" gorm:"default:0.03"` //续费团队提成
}


