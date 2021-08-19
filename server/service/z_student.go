package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"time"
)

func CreateStudent(student model.Student) (err error) {
	db := global.GVA_DB.Model(&model.Student{})
	var students []model.Student
	if rows := db.Where("`name` = ?", student.Name).Find(&students).RowsAffected; rows > 0 {
		return errors.New("姓名已存在")
	}

	err = global.GVA_DB.Create(&student).Error
	return err
}

func TurnStudent(id uint) (err error) {
	var student model.Student
	err = global.GVA_DB.Where("id = ?", id).First(&student).Error
	student.IsEntry = 1
	student.EntryTime = time.Now()
	err = global.GVA_DB.Save(&student).Error
	return
}

func DeleteStudent(student model.Student) (err error) {
	if student.ClassID != 0 {
		return errors.New("请先将该学生从对应的班级中删除")
	}
	err = global.GVA_DB.Delete(&student).Error
	return err
}

func UpdateStudent(student model.Student) (err error) {
	var oldStu model.Student
	err = global.GVA_DB.Where("ID = ?", student.ID).First(&oldStu).Error
	if err != nil {return errors.New("没有查到该学生")}

	if student.Name != oldStu.Name {
		if global.GVA_DB.Where("name = ?", student.Name).Find(&[]model.Student{}).RowsAffected > 0 {
			return errors.New("修改的学生姓名已经存在了")
		}
	}
	err = global.GVA_DB.Save(&student).Error
	return err
}

func GetStudent(id uint) (err error, student model.Student) {
	err = global.GVA_DB.Where("id = ?", id).First(&student).Error
	return
}

func GetStudentInfoList(info request.StudentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Student{})
	var students []model.Student
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("`is_entry` = ?", 1)		//根据条件查找正式学员
	if info.TeacherName != "" {
		db = db.Where("`teacher_name` LIKE ?","%"+ info.TeacherName+"%")
	}
	if info.Name != "" {
		db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
	}
	if info.Age != 0 {
		switch info.Age {
		case 6:
			db = db.Where("`age` <= ?",6)
		case 10:
			db = db.Where("`age` > ? AND `age` <= ?",6, 10)
		case 15:
			db = db.Where("`age` > ? AND `age` <= ?",10, 15)
		case 16:
			db = db.Where("`age` > ?",15)
		default:
			//
		}
	}
	if info.CourseRemain != 0 {
		switch info.CourseRemain {
		case 3:
			db = db.Where("`course_remain` <= ?",3)
		case 6:
			db = db.Where("`course_remain` > ? AND `course_remain` <= ?",3, 6)
		case 16:
			db = db.Where("`course_remain` > ? AND `course_remain` <= ?",6, 16)
		case 17:
			db = db.Where("`course_remain` > ?",16)
		default:
			//
		}
	}
	if info.CreditRemain != 0 {
		switch info.CreditRemain {
		case 50:
			db = db.Where("`credit_remain` <= ?",50)
		case 100:
			db = db.Where("`credit_remain` > ? AND `credit_remain` <= ?",50, 100)
		case 200:
			db = db.Where("`credit_remain` > ? AND `credit_remain` <= ?",100, 200)
		case 201:
			db = db.Where("`credit_remain` > ?",200)
		default:
			//
		}
	}
	if info.CourseType != "" {
		db = db.Where("`course_type` = ?",info.CourseType)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&students).Error
	return err, students, total
}

func GetTrialStudentInfoList(info request.StudentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Student{})
	var students []model.Student
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("`is_entry` = ?", 0)		//根据条件查找正式学员
	if info.Name != "" {
		db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
	}
	if info.Phone != "" {
		db = db.Where("`phone` LIKE ?","%"+ info.Phone+"%")
	}
	if info.Age != 0 {
		switch info.Age {
		case 6:
			db = db.Where("`age` <= ?",6)
		case 10:
			db = db.Where("`age` > ? AND `age` <= ?",6, 10)
		case 15:
			db = db.Where("`age` > ? AND `age` <= ?",10, 15)
		case 16:
			db = db.Where("`age` > ?",15)
		default:
			//
		}
	}
	if info.CourseType != "" {
		db = db.Where("`course_type` = ?",info.CourseType)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("entry_time desc").Find(&students).Error
	return err, students, total
}
