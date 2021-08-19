package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"time"
)

func CreatePotential(potential model.Potential) (err error) {
	db := global.GVA_DB.Model(&model.Potential{})
	if db.Where("`name` = ?", potential.Name).Find(&model.Potential{}).RowsAffected > 0 {
		return errors.New("意向学员姓名已存在")
	}
	err = global.GVA_DB.Create(&potential).Error
	return err
}

func DeletePotential(potential model.Potential) (err error) {
	if potential.Level > 0 {
		return errors.New("请先将学生的意愿降到最低后再删除")
	}
	err = global.GVA_DB.Delete(&potential).Error
	return err
}

func UpdatePotential(potential model.Potential) (err error) {
	var old model.Potential
	err = global.GVA_DB.Where("`ID` = ?", potential.ID).First(&old).Error
	if err != nil {return errors.New("没有查到该学员")}

	if potential.Name != old.Name {
		if global.GVA_DB.Where("`name` = ?", potential.Name).RowsAffected > 0 {
			return errors.New("修改的学员姓名已经存在了")
		}
	}
	err = global.GVA_DB.Save(&potential).Error
	return err
}

func GetPotential(id uint) (err error, potential model.Potential) {
	err = global.GVA_DB.Where("id = ?", id).First(&potential).Error
	return
}

func GetPotentialInfoList(info request.PotentialSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Potential{})
	var potentials []model.Potential
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("level desc").Order("updated_at desc").Find(&potentials).Error
	return err, potentials, total
}

func PotentialToTrial(potential model.Potential) (err error) {
	student := model.Student{
		Name: 		potential.Name,
		Age: 		potential.Age,
		Sex:		potential.Sex,
		Guardian: 	potential.Guardian,
		Phone: 		potential.Phone,
		CourseType: potential.CourseType,
		Comment: 	potential.Comment,

		IsEntry:  	0,
		EntryTime:  time.Now(),
	}

	db := global.GVA_DB.Model(&model.Student{})
	var students []model.Student
	if rows := db.Where("`name` = ?", student.Name).Find(&students).RowsAffected; rows > 0 {
		return errors.New("该姓名已经在学员列表中")
	}
	err = global.GVA_DB.Create(&student).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&potential).Error
	return err
}

func UpgradePotentialLevel(id uint) (err error) {
	var potential model.Potential
	global.GVA_DB.Where("id = ?", id).First(&potential)
	potential.Level = potential.Level + 1
	err = global.GVA_DB.Save(&potential).Error
	return err
}

func DegradePotentialLevel(id uint) (err error) {
	var potential model.Potential
	global.GVA_DB.Where("id = ?", id).First(&potential)
	potential.Level = potential.Level - 1
	err = global.GVA_DB.Save(&potential).Error
	return err
}