package service

import (
	"fmt"
	"myblog/dao"
	"myblog/errmsg"
	"myblog/model"
)

//<----------分类模块逻辑处理------------>

// 创建一个分类
func CreateCategory(data *model.Category) int {
	err := dao.DB.Create(&data).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除一个分类
func DelteCategory(id int) int {
	var category model.Category
	err := dao.DB.Where("id=?", id).Delete(&category).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑一个分类
func EditCategory(id int, data *model.Category) int {
	var category model.Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := dao.DB.Model(&category).Where("id=?", id).Updates(maps).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 根据分类名字检索id
func GetCategoryId(name string) (int, int) {
	var category model.Category
	err := dao.DB.Select("ID").Where("name=?", name).First(&category).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR, 0
	}
	return errmsg.SUCCESS, category.ID
}

// 分类列表
func GetCategory() ([]model.Category, int) {
	var category []model.Category
	err := dao.DB.Find(&category).Error
	if err != nil {
		fmt.Println(err)
		return nil, errmsg.ERROR
	}
	return category, errmsg.SUCCESS
}
