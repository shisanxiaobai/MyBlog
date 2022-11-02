package router

import (
	"myblog/api"
	"myblog/conf"
	"myblog/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouterStart() {
	r := gin.New() //新建一个不使用任何中间件的路由

	//静态资源访问
	r.Static("/assets", "./assets")

	//gin自定义使用路由中间件
	r.Use(gin.Recovery())      //处理错误中间件，直接使用gin默认自带的
	r.Use(middleware.Cors())   //跨域中间件
	r.Use(middleware.Logger()) //路由中间件

	//跳转根页面,用于开发测试，前后端分根页面跳转交给vue就行
	r.GET("/", api.Index)
	r.GET("/admin", api.Admin)

	//<------------------front 前台路由--------------------->

	router := r.Group("api")
	//前台文章模块
	router.GET("article/list", api.GetArtList)
	router.GET("article/:id", api.GetArtInfo)
	router.GET("articlecate/:name", api.GetCateArticle)
	//前台文章评论模块
	router.POST("comment/add", api.AddComment)
	router.GET("comment/:id", api.GetCommentInfo)

	//前台分类模块
	router.GET("category", api.GetCategory)

	//前台生活模块
	router.GET("life", api.GetAllLife)
	router.GET("life/:id", api.GetLifeInfo)
	//前台生活评论模块
	router.POST("lifecomment/add", api.AddLcomment)
	router.GET("lifecomment/:id", api.GetLifecomment)

	//前台项目模块
	router.GET("project", api.GetAllProject)
	router.GET("project/:id", api.GetProjectInfo)

	//前台留言板模块
	router.POST("board/add", api.AddBoard)
	router.GET("board", api.GetBoard)

	//admin后台登录
	router.POST("login", api.Login)
	//<------------------admin 后台路由---------------------->

	auth := r.Group("api")
	auth.Use(middleware.JwtToken()) //auth使用jwt中间件

	//后台分类模块
	auth.POST("category/add", api.AddCategory)
	auth.DELETE("category/delete/:id", api.DelCategory)
	auth.PUT("category/:id", api.EditCategory)
	auth.GET("categoryid/:name", api.GetCategoryId)

	//后台文章模块
	auth.POST("article/add", api.AddArticle)
	auth.DELETE("article/delete/:id", api.DeleteArticle)
	auth.PUT("article/edit/:id", api.EditArticle)
	auth.GET("article/search", api.SearchArticle)

	//后台文章评论模块
	auth.DELETE("comment/delete/:id", api.DeleteComment)
	auth.GET("artcomment/list", api.GetAllComment)
	

	//后台内动态模块路由
	auth.POST("life/add", api.AddLife)
	auth.DELETE("life/delete/:id", api.DeleteLife)
	auth.PUT("life/edit/:id", api.EditLife)
	auth.GET("life/search", api.SearchLife)
	//后台管理动态评论模块
	auth.DELETE("lifecomment/delete/:id", api.DeleteLcomment)
	auth.GET("lifecomment/list", api.GetAllLcomment)

	//后项目模块路由
	auth.POST("project/add", api.AddProject)
	auth.DELETE("project/delete/:id", api.DeleteProject)
	auth.PUT("project/edit/:id", api.EditProject)
	auth.GET("project/search", api.SearchProject)

	//后台留言板模块
	auth.DELETE("board/delete/:id", api.DeleteBoard)

	//<------------------------------------------------------------------>

	//注册swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//启动路由，端口号为配置端口
	r.Run(conf.Config.HttpPort)
}
