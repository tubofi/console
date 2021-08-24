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
	CourseContent 	string 			`json:"courseContent" form:"courseContent"`
	TimeString 		string 			`json:"timeString" form:"timeString"`
	Room			int				`json:"room" form:"room"`

	Students 		string			`json:"students" form:"students"`				//参课学生
	AbsentStudents 	string			`json:"absentStudents" form:"absentStudents"`	//请假学生
	CourseRecords 	[]CourseRecord 	`json:"courseRecords" form:"courseRecords"`

	IsUploadSourceCode 	int			`json:"isUploadSourceCode" form:"isUploadSourceCode"`	//是否上传源码
	IsUploadFodder 		int			`json:"isUploadFodder" form:"isUploadFodder"`			//是否上传素材
	SourceCodeUrl 	string			`json:"sourceCodeUrl" form:"sourceCodeUrl"`		//保存源码的链接
	FodderUrl 		string			`json:"fodderUrl" form:"fodderUrl"`				//保存素材的链接

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
	CourseContent 	string 			`json:"courseContent" form:"courseContent"`
	TimeString 		string 			`json:"timeString" form:"timeString"`
	Room			int				`json:"room" form:"room"`

	Students 		[]string		`json:"students" form:"students"`				//参课学生
	AbsentStudents 	[]string		`json:"absentStudents" form:"absentStudents"`	//请假学生
	CourseRecords 	[]CourseRecord 	`json:"courseRecords" form:"courseRecords"`

	IsUploadSourceCode 	int			`json:"isUploadSourceCode" form:"isUploadSourceCode"`	//是否上传源码
	IsUploadFodder 		int			`json:"isUploadFodder" form:"isUploadFodder"`			//是否上传素材
	SourceCodeUrl 	string			`json:"sourceCodeUrl" form:"sourceCodeUrl"`		//保存源码的链接
	FodderUrl 		string			`json:"fodderUrl" form:"fodderUrl"`				//保存素材的链接

	Comment 		string 			`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

