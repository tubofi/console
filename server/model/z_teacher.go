package model

import (
	"gorm.io/gorm"
	"time"
)

type Teacher struct {
	gorm.Model
	SchoolID    		uint			`json:"schoolID" form:"schoolID"`

	//个人信息
	Name 				string			`json:"name" form:"name"`
	Sex 				string			`json:"sex" form:"sex"` //女/男
	Address 			string 			`json:"address" form:"address"` //居住地址
	IdentityNumber 		string 			`json:"identityNumber" form:"identityNumber"` //身份证号
	Phone 				string			`json:"phone" form:"phone"`
	EntryTime 			time.Time 		`json:"entryTime" form:"entryTime"` //入职时间

	//基础工资
	BaseWage 			float64 		`json:"baseWage" form:"baseWage"` //基础工资
	PostWage 			float64			`json:"postWage" form:"postWage"` //岗位工资
	Allowance 			float64 		`json:"allowance" form:"allowance"` //岗位补贴
	//教师岗位工资
	CoursePer 			float64 		`json:"coursePer" form:"coursePer"` //课时费
	TurnPercent 		float64			`json:"turnPercent" form:"turnPercent"`
	RenewPercent 		float64			`json:"renewPercent" form:"renewPercent"`
	//招生岗位工资
	NewPercent 			float64			`json:"newPercent" form:"newPercent"` //新单提成
	OldPercent 			float64 		`json:"oldPercent" form:"oldPercent"` //续费提成
	NewTeamPercent 		float64			`json:"newTeamPercent" form:"newTeamPercent"` //新单团队提成
	OldTeamPercent  	float64 		`json:"oldTeamPercent" form:"oldTeamPercent"` //续费团队提成

	//教师岗位
	Classes 			[]Class			`json:"classes" form:"classes" gorm:"foreignKey:TeacherName;references:Name"`
	CourseRecords		[]CourseRecord 	`json:"courseRecords" form:"courseRecords" gorm:"foreignKey:TeacherName;references:Name"`
	Students 			[]Student		`json:"students" form:"students" gorm:"foreignKey:TeacherName;references:Name"`		//包括正式学员和试听学员
	TrialCourseRecords  []TrialCourseRecord		`json:"trialCourseRecords" form:"trialCourseRecords" gorm:"foreignKey:TeacherName;references:Name"`
	//招生岗位
	ManageStudents 		[]Student		`json:"manageStudents" form:"manageStudents" gorm:"foreignKey:ManagerName;references:Name"`
	Potentials			[]Potential		`json:"potentials" form:"potentials" gorm:"foreignKey:ManagerName;references:Name"`
	//工资
	Wages				[]Wage			`json:"wages" form:"wages" gorm:"foreignKey:TeacherName;references:Name"`

	Comment 			string			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

func (Teacher) TableName() string {
	return "teachers"
}



