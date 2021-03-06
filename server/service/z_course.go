package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"strings"
	"time"
)

func CreateCourse(resCourse model.ResCourse) (err error) {
	resCourse = createCourseRecords(resCourse)			//创建本班级学员的课程记录列表，包括参课和请假学员
	course := resCourseToCourse(resCourse)
	err = global.GVA_DB.Create(&course).Error

/*	//所有参课学员课程次数-1
	var students []model.Student
	global.GVA_DB.Where(`name IN ?`, resCourse.Students).Find(&students)
	for i, _ := range students {
		students[i].CourseRemain = students[i].CourseRemain - 1
		global.GVA_DB.Save(&students[i])
	}*/

	//修改为通过CourseChange表来变更学员课程次数
	for _, name :=  range resCourse.Students {
		description := ""
		//判断课程名称中是否存在书名号
		if strings.Index(resCourse.Name, "《") != -1 && strings.Index(resCourse.Name, "》") != -1 {
			description = resCourse.CourseType + ":" + resCourse.Name
		} else {
			description = resCourse.CourseType + ":《" + resCourse.Name + "》"
		}
		courseChange := model.CourseChange{
			StudentName: 	name,
			Amount:  		-1,			//次数减少1
			Type:  			-1,			//类型为消课
			Description:  	description,
		}
		err = CreateCourseChange(courseChange)
	}

	//判断upload是否存在，如果没有那么创建
	creatUpload(course)
	return err
}

func DeleteCourse(resCourse model.ResCourse) (err error) {
	deleteCourseRecords(resCourse.ID)
	course := resCourseToCourse(resCourse)
	err = global.GVA_DB.Delete(&course).Error
	return err
}

func UpdateCourse(resCourse model.ResCourse) (err error) {
	//更新时删除所有原记录并重新创建，评分点评什么的也都没有了
	deleteCourseRecords(resCourse.ID)
	resCourse = createCourseRecords(resCourse)
	course := resCourseToCourse(resCourse)
	err = global.GVA_DB.Save(&course).Error
	return err
}

func GetCourse(id uint) (err error, resCourse model.ResCourse) {
	course := model.Course{}
	err = global.GVA_DB.Where("id = ?", id).First(&course).Error
	resCourse = courseToResCourse(course)
	return
}

func GetCourseInfoList(info request.CourseSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&model.Course{})
	var courses []model.Course
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ClassID != 0 {
		db = db.Where("`class_id` = ?",info.ClassID)
	}
	if info.TimeString != "" {
		db = db.Where("`time_string` LIKE ?","%"+ info.TimeString+"%")
	}
	if info.Name != "" {
		db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
	}
	if info.Room != 0 {
		db = db.Where("`room` = ?",info.Room)
	}
	if info.TeacherName != "" {
		db = db.Where("`teacher_name` LIKE ?","%"+ info.TeacherName+"%")
	}
	if len(info.Students) > 0 && info.Students[0] != "" {
		db = db.Where("`students` LIKE ?","%"+ info.Students[0]+"%")
	}
	if len(info.AbsentStudents) > 0 && info.AbsentStudents[0] != "" {
		db = db.Where("`absent_students` LIKE ?","%"+ info.AbsentStudents[0]+"%")
	}

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&courses).Error

	var resCourses []model.ResCourse
	for _, course := range courses {
		resCourses = append(resCourses, courseToResCourse(course))
	}
	return err, resCourses, total
}

func GetClassListFromCourse()(err error, list []model.Class){
	err = global.GVA_DB.Preload("Students").Find(&list).Error
	return
}


func resCourseToCourse(resCourse model.ResCourse) model.Course {
	course := model.Course{
		Model:			resCourse.Model,
		SchoolID:   	resCourse.SchoolID,
		ClassID:  		resCourse.ClassID,
		Name: 			resCourse.Name,
		CourseType:     resCourse.CourseType,
		CourseContent:  resCourse.CourseContent,
		TimeString:  	resCourse.TimeString,
		Room:  			resCourse.Room,
		TeacherName:  	resCourse.TeacherName,
		Students:  		strings.Join(resCourse.Students, ";"),
		AbsentStudents: strings.Join(resCourse.AbsentStudents, ";"),	//将字符串组以指定分隔符凑成字符串
		CourseRecords	:resCourse.CourseRecords,
		Comment:  		resCourse.Comment,
	}
	return course
}

func courseToResCourse(course model.Course) model.ResCourse{
	//直接分割字符串组会出现AbsentStudents = [""]，长度为1
	absentStudents := strings.Split(course.AbsentStudents, ";")
	if len(absentStudents) == 1 && absentStudents[0] == "" {
		absentStudents = []string{}
	}
	resCourse := model.ResCourse{
		Model:			course.Model,
		SchoolID:   	course.SchoolID,
		ClassID:  		course.ClassID,
		Name: 			course.Name,
		CourseType:     course.CourseType,
		CourseContent:  course.CourseContent,
		TimeString:  	course.TimeString,
		Room:  			course.Room,
		TeacherName:  	course.TeacherName,
		Students:  		strings.Split(course.Students, ";"),		//将字符分割为字符串组
		AbsentStudents: absentStudents,
		CourseRecords:	course.CourseRecords,
		Comment:  		course.Comment,
	}
	return resCourse
}

func createCourseRecords(reCourse model.ResCourse)(res model.ResCourse){
	var records []model.CourseRecord
	for _, name := range reCourse.Students {
		record := model.CourseRecord{
			ClassID:  		reCourse.ClassID,
			StudentName:  	name,
			TeacherName:  	reCourse.TeacherName,
			CourseContent:  reCourse.CourseContent,
			CourseName:  	reCourse.Name,
			AttendTime:  	time.Now(),
			IsAbsentee:  	0,
			NeedCram:  		0,
			NeedFeedback:  	1,
		}
		records = append(records, record)
	}

	for _, name := range reCourse.AbsentStudents {
		record := model.CourseRecord{
			ClassID:  		reCourse.ClassID,
			StudentName:  	name,
			TeacherName:  	reCourse.TeacherName,		//补课的老师待定
			CourseContent:  reCourse.CourseContent,
			CourseName:  	reCourse.Name,
			AttendTime:  	time.Now(),					//补课的时间待定
			IsAbsentee:  	1,
			NeedCram:  		1,
			NeedFeedback:  	1,
		}
		records = append(records, record)
	}

	res = reCourse
	res.CourseRecords = records
	return
}

func deleteCourseRecords(id uint) {
	global.GVA_DB.Where("course_id = ?", id).Delete(&model.CourseRecord{})
}

func creatUpload(course model.Course) {
	if global.GVA_DB.Where(`course_type = ? AND course_name = ?`,
		course.CourseType, course.Name).Find(&[]model.Upload{}).RowsAffected > 0 {
		return
	}
	upload := model.Upload{
		CourseType:			course.CourseType,
		CourseName:  		course.Name,
		Bucket:  			global.GVA_CONFIG.TencentCOS.Bucket,
		Region:  			global.GVA_CONFIG.TencentCOS.Region,
		IsUploadFodder:  	0,
		IsUploadSourceCode: 0,
		SourceCodeFolder:  	"source-code/",
		FodderFolder:  		"fodder/",
		SourceCodeUrl:  	"",
		FodderUrl:  		"",
	}
	global.GVA_DB.Create(&upload)
}

