package model

import (
	"gorm.io/gorm"
	"time"
)

type School struct {
	gorm.Model

	Name 			string		`json:"name" form:"name"`
	Setup 			time.Time	`json:"setup" form:"setup"`
	Address 		string		`json:"address" form:"address"`
	Telephone 		string		`json:"telephone" form:"telephone"`

	Classes 		[]Class		`json:"classes" form:"classes"`
	Courses			[]Course	`json:"courses" form:"courses"`
	Students 		[]Student	`json:"students" form:"students"`
	Potentials 		[]Potential	`json:"potentials" form:"potentials"`

	Bills 			[]Bill		`json:"bills" form:"bills"`
	goods			[]Goods	    `json:"goods" form:"goods"`
	Teachers		[]Teacher	`json:"teachers" form:"teachers"`

	Comment 		string		`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}
