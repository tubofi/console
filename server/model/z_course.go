package model

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	SchoolID   		uint			`json:"schoolId" form:"schoolId"`
	ClassID    		uint			`json:"classId" form:"classId"`
	TeacherName 	string 			`json:"teacherName" form:"teacherName"`

	Name 			string			`json:"name" form:"name"`			//课程名称
	CourseType		string			`json:"courseType" form:"courseType"`
	CourseContent 	string 			`json:"courseContent" form:"courseContent"`
	TimeString 		string 			`json:"timeString" form:"timeString"`
	Room			int				`json:"room" form:"room"`

	Students 		string			`json:"students" form:"students"`				//参课学生
	AbsentStudents 	string			`json:"absentStudents" form:"absentStudents"`	//请假学生
	CourseRecords 	[]CourseRecord 	`json:"courseRecords" form:"courseRecords"`

	Comment 		string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

func (Course) TableName() string {
	return "courses"
}

//用来向前端反馈
type ResCourse struct {
	gorm.Model
	SchoolID   		uint			`json:"schoolId" form:"schoolId"`
	ClassID    		uint			`json:"classId" form:"classId"`
	TeacherName 	string 			`json:"teacherName" form:"teacherName"`

	Name 			string			`json:"name" form:"name"`						//课程名称
	CourseType		string			`json:"courseType" form:"courseType"`
	CourseContent 	string 			`json:"courseContent" form:"courseContent"`
	TimeString 		string 			`json:"timeString" form:"timeString"`
	Room			int				`json:"room" form:"room"`

	Students 		[]string		`json:"students" form:"students"`				//参课学生
	AbsentStudents 	[]string		`json:"absentStudents" form:"absentStudents"`	//请假学生
	CourseRecords 	[]CourseRecord 	`json:"courseRecords" form:"courseRecords"`

	Comment 		string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

