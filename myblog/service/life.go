package service

import (
	"myblog/dao"
	"myblog/errmsg"
	"myblog/model"
)

//<----------动态模块逻辑处理------------>

// 后台创建动态功能
func CreateLife(data *model.Life) int {
	err := dao.DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	} else {
		return errmsg.SUCCESS
	}
}

// 删除动态
func DeleteLife(id int) int {
	var life model.Life
	err := dao.DB.Where("id=?", id).Delete(&life).Error
	if err != nil {
		return errmsg.ERROR
	} else {
		return errmsg.SUCCESS
	}
}

// 更新动态
func EditLife(id int, data *model.Life) int {
	var life model.Life
	var maps = make(map[string]interface{})
	maps["Title"] = data.Title
	maps["Desc"] = data.Desc
	maps["Content"] = data.Content
	maps["Img"] = data.Img
	maps["ReadCount"] = data.ReadCount
	err := dao.DB.Model(&life).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	} else {
		return errmsg.SUCCESS
	}
}

// 前后通用查看动态列表
func GetAllLife() ([]model.Life, int) {
	var life []model.Life
	err := dao.DB.Find(&life).Error
	if err != nil {
		return nil, errmsg.ERROR
	} else {
		return life, errmsg.SUCCESS
	}
}

// 前台查看动态内容
func GetLifeInfo(id int) (int, model.Life) {
	var life model.Life
	err := dao.DB.Where("id=?", id).First(&life).Error
	if err != nil {
		return errmsg.ERROR, life
	} else {
		return errmsg.SUCCESS, life
	}
}


// 根据标题检索动态
func SearchLife(title string) ([]model.Life, int) {
	var LifeList []model.Life
	err := dao.DB.Select("title", "img", "created_at", `desc`, "read_count").Order("Created_At").Where("title LIKE ?", title+"%").Find(&LifeList).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return LifeList, errmsg.SUCCESS
}
