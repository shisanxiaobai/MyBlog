package service

import (
	"fmt"
	"myblog/dao"
	"myblog/errmsg"
	"myblog/model"
)

//<----------文章评论模块逻辑处理------------>

// 创建一个文章评论
func CreateComment(data *model.ArtComment) int {
	err := dao.DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除一个文章评论
func DeleteComment(id int) int {
	var comment model.ArtComment
	err := dao.DB.Where("id=?", id).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 后台获取评论列表
func GetAllComment() ([]model.ArtComment, int) {
	var artcomment []model.ArtComment
	err := dao.DB.Select("art_comment.id", "name", "art_comment.content", "Article.title").Joins("Article").Order("art_comment.created_at DESC").Find(&artcomment).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return artcomment, errmsg.SUCCESS
}

// 前台查看评论
func GetCommentInfo(pageSize int, pageNum int, id int) ([]model.ArtComment, int) {
	var artcomment []model.ArtComment
	err := dao.DB.Joins("Article").Select("name", "art_comment.created_at", "art_comment.content").Where("aid=?", id).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&artcomment).Error
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return nil, errmsg.ERROR
	}
	return artcomment, errmsg.SUCCESS
}
