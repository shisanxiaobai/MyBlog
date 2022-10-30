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

//项目模块api 下面注释用于swagger生成接口文档

// @Summary 前台添加文章评论
// @Description 前台添加一个文章评论
// @Tags	ArtComment
// @Accept  json
// @Produce json
// @Param 	artcomment body model.ArtComment false "添加文章评论"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/comment/add [post]
func AddComment(c *gin.Context) {
	var data model.ArtComment
	_ = c.ShouldBind(&data)
	code := service.CreateComment(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台删除文章评论
// @Description 后台删除一个文章评论
// @Tags	ArtComment
// @Accept  json
// @Produce json
// @Param 	id path int true "要删除的动态评论id"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/comment/delete/{id} [delete]
func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DeleteComment(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台查看动态评论列表
// @Description 后台查看动态评论
// @Tags	ArtComment
// @Accept  json
// @Produce json
// @Param  	id path int true "文章id"
// @Param 	pagesize query int true "分页尺寸"
// @Param 	pagenum query int true "页码"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/comment/{id} [get]
func GetCommentInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, code := service.GetCommentInfo(pageSize, pageNum, id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 查看所有文章评论
// @Description 后台查看所有文章评论
// @Tags	ArtComment
// @Accept  json
// @Produce json
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/artcomment/list [get]
func GetAllComment(c *gin.Context) {
	data, code := service.GetAllComment()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台查看动态评论列表
// @Description 后台查看动态评论
// @Tags	ArtComment
// @Accept  json
// @Produce json
// @Param  	id path int true "id"
// @Param 	pagesize query int true "分页尺寸"
// @Param 	pagenum query int true "页码"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/artcomment/{id} [get]
func GetArtComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, code := service.GetCommentInfo(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
