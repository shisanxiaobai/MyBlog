package api

import "github.com/gin-gonic/gin"

//根路径的跳转处理，用于开发测试接口
func Index(c *gin.Context) {
	c.HTML(200, "front", nil)
}
