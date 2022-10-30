package service

import (
	"myblog/dao"
	"myblog/errmsg"
	"myblog/model"
)

//<----------留言板模块逻辑处理------------>

//创建留言
func CreatBoard(data *model.Board) int {
	err := dao.DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//后台删除留言
func DeleteBoard(id int) int {
	var board model.Board
	err := dao.DB.Where("id=?", id).Delete(&board).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//后台查看留言
func GetBoard(pageSize int, pageNum int) ([]model.Board, int) {
	var board []model.Board
	err := dao.DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_at DESC").Find(&board).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return board, errmsg.SUCCESS
}
