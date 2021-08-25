package model

import "gorm.io/gorm"

type Upload struct {
	gorm.Model
	CourseType 				string 		`json:"courseType" form:"courseType"`
	CourseName 				string 		`json:"courseName" form:"courseName"`

	Bucket					string		`json:"bucket" form:"bucket"`
	Region					string 		`json:"region" form:"region"`

	IsUploadSourceCode 		int			`json:"isUploadSourceCode" form:"isUploadSourceCode"`	//是否上传源码
	SourceCodeFolder		string 		`json:"sourceCodeFolder" form:"sourceCodeFolder"`
	SourceCodeUrl 			string		`json:"sourceCodeUrl" form:"sourceCodeUrl"`				//保存源码的链接

	IsUploadFodder 			int			`json:"isUploadFodder" form:"isUploadFodder"`			//是否上传素材
	FodderFolder 			string 		`json:"fodderFolder" form:"fodderFolder"`
	FodderUrl 				string		`json:"fodderUrl" form:"fodderUrl"`						//保存素材的链接

	Comment 				string 		`json:"comment" form:"comment" gorm:"type:varchar(512)"`
}

func (Upload) TableName() string {
	return "uploads"
}
