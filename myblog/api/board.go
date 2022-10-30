package api

import (
	_ "myblog/docs"
	"myblog/errmsg"
	"myblog/model"
	"myblog/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//留言板模块api 下面注释用于swagger生成接口文档

// @Summar	添加留言
// @Description 前台添加一个留言
// @Tags	Board
// @Accept  json
// @Produce json
// @Param 	board body model.Board false "添加留言"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/board/add [post]
func AddBoard(c *gin.Context) {
	var data model.Board
	_ = c.ShouldBind(&data)
	code := service.CreatBoard(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台删除留言
// @Description 后台删除一个留言
// @Tags	Board
// @Accept  json
// @Produce json
// @Param 	id path int true "要删除留言的id"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/board/delete/{id} [delete]
func DeleteBoard(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DeleteBoard(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 前台查看留言
// @Description 前台查看留言
// @Tags	Board
// @Accept  json
// @Produce json
// @Param 	pagesize query int true "分页尺寸"
// @Param 	pagenum query int true "页码"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/board [get]
func GetBoard(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	data, code := service.GetBoard(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
