package api

import (
	"myblog/errmsg"
	"myblog/model"
	"myblog/service"
	"strconv"

	_ "myblog/docs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//动态模块api 下面注释用于swagger生成接口文档

// @Summary 检索动态
// @Description 根据标题检索动态
// @Tags	Life
// @Accept  json
// @Produce json
// @Param 	title query string true "查找动态"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/life/search [get]
// @Security ApiKeyAuth
func SearchLife(c *gin.Context) {
	title := c.Query("title")
	data, code := service.SearchLife(title)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台添加动态
// @Description 后台添加一个动态
// @Tags	Life
// @Accept  json
// @Produce json
// @Param 	life body model.Life false "添加动态"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/life/add [post]
// @Security ApiKeyAuth
func AddLife(c *gin.Context) {
	var data model.Life
	_ = c.ShouldBind(&data)
	code := service.CreateLife(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台删除动态
// @Description 后台删除一个动态
// @Tags	Life
// @Accept  json
// @Produce json
// @Param 	id path int true "要删除动态的id"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/life/delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteLife(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DeleteLife(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台更新动态
// @Description 后台编辑一个动态
// @Tags	Life
// @Accept  json
// @Produce json
// @Param 	id path int true "要编辑动态的id"
// @Param 	life body model.Life false "编辑动态"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/life/edit/{id} [put]
// @Security ApiKeyAuth
func EditLife(c *gin.Context) {
	var data model.Life
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&data)
	code := service.EditLife(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 查看所有动态
// @Description 前后台通用查看所有动态
// @Tags	Life
// @Accept  json
// @Produce json
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/life [get]
func GetAllLife(c *gin.Context) {
	data, code := service.GetAllLife()
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"data":     data,
		"msessage": errmsg.GetErrMsg(code),
	})
}

// @Summary 前台查看动态内容
// @Description 前台查看动态内容
// @Tags	Life
// @Accept  json
// @Produce json
// @Param id path int true "要查看的id"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/life/{id} [get]
func GetLifeInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code, data := service.GetLifeInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
