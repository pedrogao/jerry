package middleware

import (
	Err "github.com/PedroGao/jerry/libs/err"
	"github.com/PedroGao/jerry/libs/token"
	"github.com/PedroGao/jerry/model"
	"github.com/gin-gonic/gin"
)

const (
	AuthHeader = "Authorization"
)

// access when user is logined
func LoginRequired(c *gin.Context) {
	authHeader := c.Request.Header.Get(AuthHeader)
	if authHeader == "" {
		c.JSON(Err.Forbidden.HTTPCode, Err.Forbidden.SetUrl(c.Request.URL.String()))
		c.Abort()
	} else {
		indentify, err := token.JwtInstance.VerifyAccessToken(authHeader)
		if err != nil {
			c.JSON(Err.ParamsErr.HTTPCode, Err.ParamsErr.SetUrl(c.Request.URL.String()).Clear().Add(err.Error()))
			c.Abort()
		} else {
			user := &model.UserModel{
				Username: indentify,
			}
			ok, _ := model.DB.Get(user)
			if !ok {
				c.JSON(Err.ParamsErr.HTTPCode, Err.ParamsErr.SetUrl(c.Request.URL.String()).Clear().Add("用户未找到"))
				c.Abort()
			} else {
				c.Set("user", user)
				c.Next()
			}
		}
	}
}

func AdminRequired(c *gin.Context) {
	authHeader := c.Request.Header.Get(AuthHeader)
	if authHeader == "" {
		c.JSON(Err.Forbidden.HTTPCode, Err.Forbidden.SetUrl(c.Request.URL.String()))
		c.Abort()
	} else {
		indentify, err := token.JwtInstance.VerifyAccessToken(authHeader)
		if err != nil {
			c.JSON(Err.ParamsErr.HTTPCode, Err.ParamsErr.SetUrl(c.Request.URL.String()).Clear().Add(err.Error()))
			c.Abort()
		} else {
			user := &model.UserModel{
				Username: indentify,
			}
			ok, _ := model.DB.Get(user)
			if !ok {
				c.JSON(Err.ParamsErr.HTTPCode, Err.ParamsErr.SetUrl(c.Request.URL.String()).Clear().Add("用户未找到"))
				c.Abort()
			} else {
				if !user.IsSuper() {
					c.JSON(Err.Forbidden.HTTPCode, Err.Forbidden.SetUrl(c.Request.URL.String()).Clear().Add("只有超级管理员可访问"))
					c.Abort()
				} else {
					c.Set("user", user)
					c.Next()
				}
			}
		}
	}
}

func RefreshRequired(c *gin.Context) {
	authHeader := c.Request.Header.Get(AuthHeader)
	if authHeader == "" {
		c.JSON(Err.Forbidden.HTTPCode, Err.Forbidden.SetUrl(c.Request.URL.String()))
		c.Abort()
	} else {
		indentify, err := token.JwtInstance.VerifyRefreshToken(authHeader)
		if err != nil {
			c.JSON(Err.ParamsErr.HTTPCode, Err.ParamsErr.SetUrl(c.Request.URL.String()).Clear().Add(err.Error()))
			c.Abort()
		} else {
			user := &model.UserModel{
				Username: indentify,
			}
			ok, _ := model.DB.Get(user)
			if !ok {
				c.JSON(Err.ParamsErr.HTTPCode, Err.ParamsErr.SetUrl(c.Request.URL.String()).Clear().Add("用户未找到"))
				c.Abort()
			} else {
				c.Set("user", user)
				c.Next()
			}
		}
	}
}
