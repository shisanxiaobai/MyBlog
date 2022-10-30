package service

import (
	"myblog/dao"
	"myblog/errmsg"
	"myblog/model"
)

//<----------项目模块逻辑处理------------>

// 后台创建项目功能
func CreateProject(data *model.Project) int {
	err := dao.DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	} else {
		return errmsg.SUCCESS
	}
}

// 删除项目
func DeleteProject(id int) int {
	var project model.Project
	err := dao.DB.Where("id=?", id).Delete(&project).Error
	if err != nil {
		return errmsg.ERROR
	} else {
		return errmsg.SUCCESS
	}
}

// 更新项目
func EditProject(id int, data *model.Project) int {
	var project model.Project
	var maps = make(map[string]interface{})
	maps["Name"] = data.Name
	maps["Desc"] = data.Desc
	maps["Url"] = data.Url
	maps["Img"] = data.Img
	err := dao.DB.Model(&project).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	} else {
		return errmsg.SUCCESS
	}
}

// 前后通用查看项目列表
func GetAllProject() ([]model.Project, int) {
	var project []model.Project
	err := dao.DB.Find(&project).Error
	if err != nil {
		return nil, errmsg.ERROR
	} else {
		return project, errmsg.SUCCESS
	}
}

// 项目内容
func GetProjectInfo(id int) (int, model.Project) {
	var project model.Project
	err := dao.DB.Where("id=?", id).First(&project).Error
	if err != nil {
		return errmsg.ERROR, project
	} else {
		return errmsg.SUCCESS, project
	}
}

// 根据标题检索项目
func SearchProject(title string) ([]model.Project, int) {
	var ProjectList []model.Project
	err := dao.DB.Select("title", "img", "created_at", "desc", "url").Order("Created_At").Where("title LIKE ?", title+"%").Find(&ProjectList).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return ProjectList, errmsg.SUCCESS
}
