package service

import (
	"fmt"
	"myblog/dao"
	"myblog/errmsg"
	"myblog/model"
)

//<----------动态评论模块逻辑处理------------>

// 前台添加动态评论
func CreateLComment(data *model.LifeComment) int {
	err := dao.DB.Create(&data).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 后台删除评论
func DeleteLcomment(id int) int {
	var lifecomment model.LifeComment
	err := dao.DB.Where("id=?", id).Delete(&lifecomment).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 后台获取所有评论
func GetAllLcomment(pageSize int, pageNum int) ([]model.LifeComment, int) {
	var lifecomment []model.LifeComment
	err := dao.DB.Joins("Life").Select("life_comment.id", "Life.title", "name", "life_comment.created_at", "life_comment.content").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("life_comment.created_at DESC").Find(&lifecomment).Error
	if err != nil {
		fmt.Println(err)
		return nil, errmsg.ERROR
	}
	return lifecomment, errmsg.SUCCESS
}

// 前台查看评论
func GetLifeComment(pageSize int, pageNum int, id int) ([]model.LifeComment, int) {
	var lifecomment []model.LifeComment
	err := dao.DB.Joins("Life").Select("Life.title", "name", "life_comment.created_at", "life_comment.content").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("lid=?", id).Order("created_at Desc").Find(&lifecomment).Error
	if err != nil {
		fmt.Println(err)
		return nil, errmsg.ERROR
	}
	return lifecomment, errmsg.SUCCESS
}
