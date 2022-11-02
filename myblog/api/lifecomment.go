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

//动态评论模块api 下面注释用于swagger生成接口文档

// @Summary 前台添加生活评论
// @Description 前台添加一个生活评论
// @Tags	LifeComment
// @Accept  json
// @Produce json
// @Param 	lifecomment body model.LifeComment false "添加生活评论"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/lifecomment/add [post]
func AddLcomment(c *gin.Context) {
	var data model.LifeComment
	_ = c.ShouldBind(&data)
	code := service.CreateLComment(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台删除动态评论
// @Description 后台删除一个动态评论
// @Tags	Life
// @Accept  json
// @Produce json
// @Param 	id path int true "要删除的动态评论id"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/lifecomment/delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteLcomment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DeleteLcomment(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台查看动态评论列表
// @Description 后台查看动态评论
// @Tags	LifeComment
// @Accept  json
// @Produce json
// @Param 	pagesize query int true "分页尺寸"
// @Param 	pagenum query int true "页码"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/lifecomment/list [get]
func GetAllLcomment(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	data, code := service.GetAllLcomment(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 前台查看动态评论
// @Description 前台查看动态评论
// @Tags	LifeComment
// @Accept  json
// @Produce json
// @Param 	pagesize query int true "分页尺寸"
// @Param 	pagenum query int true "页码"
// @Param 	id path int true "life的id"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/lifecomment/{id} [get]
func GetLifecomment(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := service.GetLifeComment(pageSize, pageNum, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
