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

//项目模块api 下面注释用于swagger生成接口文档

// @Summary 检索项目
// @Description 根据标题检索项目
// @Tags	Project
// @Accept  json
// @Produce json
// @Param 	title query string true "查找项目"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/project/search [get]
// @Security ApiKeyAuth
func SearchProject(c *gin.Context) {
	title := c.Query("title")
	data, code := service.SearchProject(title)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台添加项目
// @Description 后台添加一个项目
// @Tags	Project
// @Accept  json
// @Produce json
// @Param 	project body model.Project false "添加项目"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/project/add [post]
// @Security ApiKeyAuth
func AddProject(c *gin.Context) {
	var data model.Project
	_ = c.ShouldBind(&data)
	code := service.CreateProject(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台删除项目
// @Description 后台删除一个项目
// @Tags	Project
// @Accept  json
// @Produce json
// @Param 	id path int true "删除项目的id"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/project/delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteProject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DeleteProject(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台更新项目
// @Description 后台编辑一个项目
// @Tags	Project
// @Accept  json
// @Produce json
// @Param 	id path int true "要编辑项目的id"
// @Param 	project body model.Project false "编辑项目"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/project/edit/{id} [put]
// @Security ApiKeyAuth
func EditProject(c *gin.Context) {
	var data model.Project
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&data)
	code := service.EditProject(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 查看所有项目
// @Description 前后台通用查看所有项目
// @Tags	Project
// @Accept  json
// @Produce json
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/project [get]
func GetAllProject(c *gin.Context) {
	data, code := service.GetAllProject()
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"data":     data,
		"msessage": errmsg.GetErrMsg(code),
	})
}

// @Summary 前台查看项目内容
// @Description 前台查看项目内容
// @Tags	Project
// @Accept  json
// @Produce json
// @Param 	id path int true "要查看项目的id"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/project/{id} [get]
func GetProjectInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code, data := service.GetProjectInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
