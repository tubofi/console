package model

import (
	"gorm.io/gorm"
	"time"
)

type Class struct {
	gorm.Model
	SchoolID  		uint				`json:"schoolId" form:"schoolId"`
	TeacherName 	string				`json:"teacherName" form:"teacherName"`
	//ClassInfo 	ClassInfo 			`gorm:"embedded;embeddedPrefix:class_"`
	//Name 			string				`json:"name" form:"name"`	//S-2-3-Saturday-10:20-1
	//CreatedTime 	time.Time 			`json:"createdTime" form:"createdTime"`
	Room 			int					`json:"room" form:"room"`
	CourseType		string				`json:"courseType" form:"courseType"` 		//
	Stage			int					`json:"stage" form:"stage"`
	Phase			int					`json:"phase" form:"phase"`		//第几期
	CourseContent	string				`json:"courseContent" form:"courseContent"`
	TimeString 		string 				`json:"timeString" form:"timeString"`
	Weekday 		time.Weekday		`json:"weekday" form:"weekday"` //周几上课？
	Hour 			int					`json:"hour" form:"hour"`		//上课的时
	Minute			int					`json:"minute" form:"minute"`	//分钟

	Students		[]Student			`json:"students" form:"students"`
	Courses			[]Course			`json:"courses" form:"courses"`
	CourseRecords	[]CourseRecord		`json:"courseRecords" form:"courseRecords"`

	Comment 		string 				`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

func (Class) TableName() string {
	return "classes"
}
type ClassInfo struct {
	//
}
type ClassDescription struct {
	Type 		string			`json:"type" form:"type"`
	Describe 	string			`json:"describe" form:"describe"`
}

type ProfileClass struct {
	Time 		string					`json:"time" form:"time"`
	Monday 		[]ClassDescription		`json:"monday" form:"monday"`
	Tuesday 	[]ClassDescription		`json:"tuesday" form:"tuesday"`
	Wednesday 	[]ClassDescription		`json:"wednesday" form:"wednesday"`
	Thursday 	[]ClassDescription		`json:"thursday" form:"thursday"`
	Friday 		[]ClassDescription		`json:"friday" form:"friday"`
	Saturday 	[]ClassDescription		`json:"saturday" form:"saturday"`
	Sunday 		[]ClassDescription		`json:"sunday" form:"sunday"`
}