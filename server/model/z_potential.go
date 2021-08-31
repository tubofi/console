package model

import (
	"gorm.io/gorm"
	"time"
)

type Potential struct {
	gorm.Model
	SchoolID 	uint		`json:"schoolID" form:"schoolID"`
	ManagerName	string		`json:"managerName" form:"managerName"`		//管理人

	Name 		string		`json:"name" form:"name"`
	Age 		float64		`json:"age" form:"age"`
	Sex 		string		`json:"sex" form:"sex"`
	Guardian 	string		`json:"guardian" form:"guardian"`
	Phone 		string		`json:"phone" form:"phone"`
	CourseType 	string		`json:"courseType" form:"courseType" `		//将要试听的课程类型

	Level		int			`json:"level" form:"level" gorm:"default:2"`   //2重视 1一般 0忽略
	Accesses 	[]Access	`json:"accesses" form:"accesses" gorm:"foreignKey:PotentialName;references:Name"`
	Comment 	string 		`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

type Access struct {
	gorm.Model
	PotentialName	string		`json:"potentialName" form:"potentialName"`
	Time 			time.Time	`json:"time" form:"time"`
	Comment 		string 		`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

func (Potential) TableName() string {
	return "potentials"
}
