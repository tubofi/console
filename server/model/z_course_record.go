package model

import (
	"gorm.io/gorm"
	"time"
)

//课堂评分
type CourseRecord struct {
	gorm.Model
	ClassID    		uint		`json:"classId" form:"classId"`
	CourseID   		uint		`json:"courseId" form:"courseId"`
	StudentName  	string		`json:"studentName" form:"studentName"`
	TeacherName  	string		`json:"teacherName" form:"teacherName"`

	CourseContent 	string 		`json:"courseContent" form:"courseContent"`
	CourseName 		string		`json:"courseName" form:"courseName"`
	AttendTime 		time.Time 	`json:"attendTime" form:"attendTime"`		//上课时间
	IsAbsentee 		int 		`json:"isAbsentee" form:"isAbsentee"`		//0为正式上课，1为补课
	NeedCram 		int 		`json:"needCram" form:"needCram"`			//0为不需要补课，1为需要补课
	NeedFeedback 	int 		`json:"needFeedback" form:"needFeedback"`		//0为已反馈，1为未反馈

	Total 			int 		`json:"total" form:"total" gorm:"default:0"`	 				//总分 0~100
	Punctuality		int			`json:"punctuality" form:"punctuality" gorm:"default:0"`		//遵纪守时0~5，是否有迟到早退情况
	Discipline 		int 		`json:"discipline" form:"discipline" gorm:"default:0"`		//课堂纪律0~10，嬉笑打闹，随便走动，吃零食
	Concentration	int			`json:"concentration" form:"concentration" gorm:"default:0"`	//专注度0~15，反应课堂参与度。紧跟老师节奏？积极思考？踊跃回答问题？
	Innovation 		int 		`json:"innovation" form:"innovation" gorm:"default:0"`	 	//创新能力0~15 课堂上提出不一样的想法？对某种做法产生疑问？不一样的实现方法
	Logic 			int 		`json:"logic" form:"logic" gorm:"default:0"` 				//逻辑性 0~20 对课堂任务和目标是否清晰？执行是否有条理？对程序逻辑的理解程度？
	Mathematics 	int 		`json:"mathematics" form:"mathematics" gorm:"default:0"`	 	//数学知识 0~20 对作品中的数学知识掌握程度，是否能够理解程序算法
	Complete 		int 		`json:"complete" form:"complete" gorm:"default:0"`	 		//完成度 0~15。反应独立完成作品能力，自己探索的部分有多少？是不是跟着老师做下来的？
	Review 			string		`json:"review" form:"review" gorm:"type:varchar(512)"`		//教师点评

	Comment 		string		`json:"comment" form:"comment" gorm:"type:varchar(512)"`	//教师点评
}

func (CourseRecord) TableName() string {
	return "course_records"
}
