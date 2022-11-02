package api

import (
	_ "myblog/docs"
	"myblog/errmsg"
	"myblog/middleware"
	"myblog/model"
	"myblog/service"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//登录模块api 下面注释用于swagger生成接口文档

// @Summary	登录
// @Description 登录
// @Tags	Login
// @Accept  json
// @Produce json
// @Param 	login body model.User true "登录"
// @Success 200 {object} model.Reponse
// @Failure 400 {object} model.Reponse
// @Failure 404 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/login [post]
func Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)
	var token string
	code := service.CheckLogin(user.Username, user.Password)
	if code == errmsg.SUCCESS {
		setToken(c, user.Username)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    user.Username,
			"message": errmsg.GetErrMsg(code),
			"token":   token,
		})
	}
}

// token生成函数，用于跨域鉴权
func setToken(c *gin.Context, username string) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,    //签发时间
			ExpiresAt: time.Now().Unix() + 604800, //过期时间
			Issuer:    "LMY",                      //签发人
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    username,
		"message": errmsg.GetErrMsg(200),
		"token":   token,
	})

}
