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

//动态模块api 下面注释用于swagger生成接口文档

// @Summary 后台添加分类
// @Description 后台添加一个分类
// @Tags	Category
// @Accept  json
// @Produce json
// @Param 	category body model.Category false "添加动态"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/category/add [post]
// @Security ApiKeyAuth
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBind(&data)
	code := service.CreateCategory(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"message:": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台删除分类
// @Description 后台删除一个分类
// @Tags	Category
// @Accept  json
// @Produce json
// @Param id path int true "要删除分类的id"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/category/delete/{id} [delete]
// @Security ApiKeyAuth
func DelCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DelteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台编辑分类
// @Description 后台编辑一个分类
// @Tags	Category
// @Accept  json
// @Produce json
// @Param 	id path int true "要编辑分类的id"
// @Param 	category body model.Category false "编辑分类"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/category/{id} [put]
// @Security ApiKeyAuth
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBind(&data)
	code := service.EditCategory(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台根据name检索
// @Description 根据name检索id
// @Tags	Category
// @Accept  json
// @Produce json
// @Param name path string true "要检索的name"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/categoryid/{name} [get]
// @Security ApiKeyAuth
func GetCategoryId(c *gin.Context) {
	name := c.Param("name")
	code, id := service.GetCategoryId(name)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"id":      id,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台查看类别列表
// @Description 查看所有类别
// @Tags	Category
// @Accept  json
// @Produce json
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/category [get]
func GetCategory(c *gin.Context) {
	data, code := service.GetCategory()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
