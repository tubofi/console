
package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"time"
)

func CreateClass(class model.Class) (err error) {
	//组合课程内容字符串
	class.CourseContent = formCourseContent(class.CourseType, class.Stage, class.Phase)
	//组合上课时间字符串
	class.TimeString = formTimeString(class.Weekday, class.Hour, class.Minute)

	var classes []model.Class
	db := global.GVA_DB.Model(&model.Class{})
	if rows := db.Where("`time_string` = ? AND `room` = ?",
		class.TimeString, class.Room).Find(&classes).RowsAffected; rows > 0 {
		return errors.New("该教室在相同时间已经安排班级")
	}

	for _, student := range class.Students {
		student.CourseContent = class.CourseContent
		student.CourseType = class.CourseType
		student.TeacherName = class.TeacherName
	}

	err = global.GVA_DB.Create(&class).Error
	if err != nil { return err }

	return nil
}

func DeleteClass(class model.Class) (err error) {
	var students []model.Student
	err = global.GVA_DB.Where("class_id = ?", class.ID).Find(&students).Error
	if err != nil {return err}
	for _, student := range students {
		student.ClassID = 0
		student.TeacherName = ""
		student.CourseContent = ""
		student.CourseType = ""
		global.GVA_DB.Save(&student)
	}
	err = global.GVA_DB.Delete(&class).Error
	return err
}

func UpdateClass(class model.Class) (err error) {
	var students []model.Student
	err = global.GVA_DB.Where("class_id = ?", class.ID).Find(&students).Error
	if err != nil {return err}
	for _, student := range students {
		student.ClassID = 0
		student.TeacherName = ""
		student.CourseContent = ""
		student.CourseType = ""
		global.GVA_DB.Save(&student)
	}

	for _, student := range class.Students {
		student.TeacherName = class.TeacherName
		student.CourseType = class.CourseType
		student.CourseContent = class.CourseContent
	}
	err = global.GVA_DB.Save(&class).Error
	return err
}

func GetClass(id uint) (err error, class model.Class) {
	err = global.GVA_DB.Preload("Students").Where("id = ?", id).First(&class).Error
	return
}

func GetClassInfoList(info request.ClassSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Class{})
	var classs []model.Class
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CourseContent != "" {
		db = db.Where("`course_content` LIKE ?","%"+ info.CourseContent+"%")
	}
	if info.TimeString != "" {
		db = db.Where("`time_string` LIKE ?","%"+ info.TimeString+"%")
	}
	if info.Room != 0 {
		db = db.Where("`room` = ?",info.Room)
	}
	if info.TeacherName != "" {
		db = db.Where("`teacher_name` LIKE ?","%"+ info.TeacherName+"%")
	}

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Students").Order("time_string").Find(&classs).Error
	return err, classs, total
}

func GetIdleStudents() (err error, students []model.Student) {
	db := global.GVA_DB.Model(&model.Student{})
	err = db.Where("`is_entry` = ? AND `class_id` = ?", 1, 0).Find(&students).Error
	if err != nil {
		return err, nil
	}

	return err, students
}

func GetProfileClass() (err error, profileClasses []model.ProfileClass) {
	times := []string{"8:50 ~ 10:20", "10:30 ~ 12:00", "13:40 ~ 15:10", "15:20 ~ 16:50", "17:00 ~ 18:30", "18:30 ~ 20:00"}

	for _, time := range times {
		pc := model.ProfileClass{}
		pc.Time = time
		profileClasses = append(profileClasses, pc)
	}

	var classes []model.Class
	err = global.GVA_DB.Find(&classes).Error
	if err != nil {return err, nil}
	for _, class := range classes {
		courseType := class.CourseType
		courseDescription := fmt.Sprintf("%s 教室%d %s", class.CourseContent, class.Room, class.TeacherName)
		classDescription := model.ClassDescription{
			Type: 		courseType,
			Describe:  	courseDescription,
		}

		index, err := profileIndex(class.Hour)
		if index == 0 && err != nil {
			return errors.New(fmt.Sprintf("%s: 排课时间存在问题，请检查", courseDescription)), nil
		}

		switch class.Weekday {
		case 0:
			profileClasses[index].Sunday = append(profileClasses[index].Sunday, classDescription)
		case 1:
			profileClasses[index].Monday = append(profileClasses[index].Monday, classDescription)
		case 2:
			profileClasses[index].Tuesday = append(profileClasses[index].Tuesday, classDescription)
		case 3:
			profileClasses[index].Wednesday = append(profileClasses[index].Wednesday, classDescription)
		case 4:
			profileClasses[index].Thursday = append(profileClasses[index].Thursday, classDescription)
		case 5:
			profileClasses[index].Friday = append(profileClasses[index].Friday, classDescription)
		case 6:
			profileClasses[index].Saturday = append(profileClasses[index].Saturday, classDescription)
		default:
			return errors.New(fmt.Sprintf("%s: 排课时间存在问题，请检查", courseDescription)), nil
		}
	}

	return err, profileClasses
}

func profileIndex(hour int) (index int, err error) {
	switch hour {
	case 8:
		index = 0
	case 10:
		index = 1
	case 13:
		index = 2
	case 15:
		index = 3
	case 17:
		index = 4
	case 18:
		index = 5
	default:
		index = 0
		err = errors.New("排课时间存在问题，请检查")
	}
	return index, err
}

func formCourseContent(courseType string, stage int, phase int)(content string){
	switch courseType {
	case "scratch":
		content = fmt.Sprintf("S-%d-%d", stage, phase)
	case "scratchJR":
		content = fmt.Sprintf("J-%d-%d", stage, phase)
	case "osmo":
		content = fmt.Sprintf("O-%d-%d", stage, phase)
	case "python":
		content = fmt.Sprintf("P-%d-%d", stage, phase)
	case "cpp":
		content = fmt.Sprintf("C-%d-%d", stage, phase)
	default:
		content = ""
	}
	return content
}

func formTimeString(weekday time.Weekday, hour int, minute int)(str string){
	var weekdayString string
	switch weekday {
	case 0:
		weekdayString = "星期日"
	case 1:
		weekdayString = "星期一"
	case 2:
		weekdayString = "星期二"
	case 3:
		weekdayString = "星期三"
	case 4:
		weekdayString = "星期四"
	case 5:
		weekdayString = "星期五"
	case 6:
		weekdayString = "星期六"
	default:
		weekdayString = ""
	}
	str = fmt.Sprintf("%s %d:%d", weekdayString, hour, minute)
	return str
}

/*
func addStudent(names []string, class model.Class)(students []model.Student, err error){
	for _, name := range names {
		student := model.Student{}
		err = global.GVA_DB.Where("name = ?", name).First(&student).Error
		if err != nil {
			return nil, errors.New("找不到学生:" + name)
		}
		student.ClassID = class.ID
		student.TeacherName = class.TeacherName
		student.CourseType = class.CourseType
		student.CourseContent = class.CourseContent

		students = append(students, student)
	}
	return students, nil
}
 */