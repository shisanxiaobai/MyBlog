package main

import (
	_ "myblog/docs"
	"myblog/router"
)

//<------------------项目启动入口----------------------->
//<----------下面文档注释用于生成swagger文档------------->

// @title myblog gin+gorm
// @version 1.0
// @description gin+gorm+jwt+cors+viper等等
// @license.name Apache 2.0
// @contact.name go-swagger帮助文档
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host localhost:8000
// @Basepath /

func main() {

	router.RouterStart()

}
