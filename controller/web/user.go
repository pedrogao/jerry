package web

import (
	"fmt"
	"github.com/PedroGao/jerry/form"
	Err "github.com/PedroGao/jerry/libs/err"
	"github.com/PedroGao/jerry/libs/token"
	"github.com/PedroGao/jerry/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var (
		login                     form.Login
		accessToken, refreshToken string
		err                       error
	)
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(Err.ParamsErr.HTTPCode, Err.ParamsErr.SetUrl(c.Request.URL.String()).Clear().Add("参数错误，请检查参数"))
		return
	}

	if err := login.ValidateNameAndPassword(); err != nil {
		c.JSON(Err.ParamsErr.HTTPCode, Err.ParamsErr.SetUrl(c.Request.URL.String()).Clear().Add(err.Error()))
		return
	}

	accessToken, refreshToken, err = token.JwtInstance.GenerateTokens(login.NickName)

	if err != nil {
		c.JSON(Err.ParamsErr.HTTPCode, Err.ParamsErr.SetUrl(c.Request.URL.String()).Clear().Add(err.Error()))
	}

	c.JSON(Err.OK.HTTPCode, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func GetUsers(c *gin.Context) {
	infos, e := service.ListUser()
	value, exists := c.Get("user")
	if exists {
		fmt.Println(value)
	}
	if e != nil {
		c.JSON(Err.UserNotFound.HTTPCode, Err.UserNotFound.SetUrl(c.Request.URL.String()))
		return
	}
	c.JSON(http.StatusOK, infos)
}
