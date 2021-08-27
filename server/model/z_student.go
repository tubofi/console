package model

import (
	"gorm.io/gorm"
	"time"
)

type StudentInfo struct {
	Name 			string			`json:"name" form:"name"`
	Age 			int				`json:"age" form:"age"`
	//BirthDay 		time.Time 		`json:"birthDay" form:"birthDay"`		//出生日期
	Sex 			int				`json:"sex" form:"sex"`
	//Grade			int				`json:"grade" form:"grade"`
	School 			string			`json:"school" form:"school"`			//这里指就读的学校
	Address 		string			`json:"address" form:"address"`

	Guardian 		string			`json:"guardian" form:"guardian"`
	Phone			string			`json:"phone" form:"phone"`
	Referee	    	string			`json:"referee" form:"referee"`			//引荐人：转介绍？招生？
}

type Student struct {
	gorm.Model
	SchoolID    	uint				`json:"schoolID" form:"schoolID"`
	ClassID     	uint				`json:"classID" form:"classID"`
	TeacherName   	string				`json:"teacherName" form:"teacherName"`		//授课教师
	ManagerName		string				`json:"managerName" form:"managerName"`		//管理人
	UserID 			uint 				`json:"userId" form:"userId"`
	//StudentInfo 	StudentInfo 		`gorm:"embedded"`

	Name 			string			`json:"name" form:"name"`
	Age 			int				`json:"age" form:"age"`
	Birthday 		time.Time 		`json:"birthday" form:"birthday"`		//出生日期
	Sex 			string			`json:"sex" form:"sex"`
	//Grade			int				`json:"grade" form:"grade"`
	School 			string			`json:"school" form:"school"`			//这里指就读的学校
	Address 		string			`json:"address" form:"address"`

	Guardian 		string			`json:"guardian" form:"guardian"`
	Phone			string			`json:"phone" form:"phone"`
	Referee	    	string			`json:"referee" form:"referee"`			//引荐人：转介绍？招生？

	IsEntry 		int					`json:"isEntry" form:"isEntry"`				//1为正式，0为体验生
	EntryTime 		time.Time 			`json:"entryTime" form:"entryTime" `							//报名时间
	CourseType 		string				`json:"courseType" form:"courseType" `							//课程类型
	CourseContent   string 				`json:"courseContent" form:"courseContent"`	//课程内容
	CourseRemain	int					`json:"courseRemain" form:"courseRemain" gorm:"default:0"`		//剩余课次
	CreditRemain	int					`json:"creditRemain" form:"creditRemain" gorm:"default:0"`				//剩余积分
	CourseRecords 	[]CourseRecord 		`json:"courseRecords" form:"courseRecords" gorm:"foreignKey:StudentName;references:Name"`   //课堂评分
	Payments		[]Payment			`json:"payments" form:"payments" gorm:"foreignKey:StudentName;references:Name"`
	Gifts			[]Gift 				`json:"gifts" form:"gifts" gorm:"foreignKey:StudentName;references:Name"` 			//选的礼物
	CourseChanges 	[]CourseChange		`json:"courseChanges" form:"courseChanges" gorm:"foreignKey:StudentName;references:Name"`		//课程次数变化记录
	CreditChanges 	[]CreditChange		`json:"creditChanges" form:"creditChanges" gorm:"foreignKey:StudentName;references:Name"`		//积分变化记录

	Comment 		string				`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

func (Student) TableName() string {
	return "students"
}

