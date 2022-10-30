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

// @Summary 检索文章
// @Description 根据标题检索文章
// @Tags	Article
// @Accept  json
// @Produce json
// @Param title query string true "文章标题"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/article/search [get]
func SearchArticle(c *gin.Context) {
	title := c.Query("title")
	data, code := service.SearchArticle(title)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台添加文章
// @Description 后台添加一个文章
// @Tags	Article
// @Accept  json
// @Produce json
// @Param article body model.Article false "添加文章"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/article/add [post]
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBind(&data)

	code := service.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台删除文章
// @Description 后台删除一个文章
// @Tags	Article
// @Accept  json
// @Produce json
// @Param id path int true "要删除文章的id"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/article/delete/{id} [delete]
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := service.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 后台更新文章
// @Description 后台编辑一个动态
// @Tags	Article
// @Accept  json
// @Produce json
// @Param id path int true "要编辑分类的id"
// @Param category body model.Category false "编辑分类"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/article/edit/{id} [put]
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&data)

	code := service.UpdateArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 前台查看分类文章
// @Description 查看分类文章
// @Tags	Article
// @Accept  json
// @Produce json
// @Param name path string true "分类名字"
// @Param pagesize query int true "分页尺寸"
// @Param pagenum query int true "页码"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/articlecate/{name} [get]
func GetCateArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	name := c.Param("name")
	data, code := service.GetCateArt(name, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 前台查看文章内容
// @Description 前台查看文章内容
// @Tags	Article
// @Accept  json
// @Produce json
// @Param id path int true "要查看的id"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/article/{id} [get]
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := service.GetArticleInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// @Summary 查看文章列表
// @Description 查看文章列表
// @Tags	Article
// @Accept  json
// @Produce json
// @Param title query string false "title"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/article/list [get]
func GetArtList(c *gin.Context) {
	title := c.Query("title")

	if len(title) == 0 {
		data, code := service.GetArticleList()
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		data, code := service.SearchArticle(title)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	}
}
