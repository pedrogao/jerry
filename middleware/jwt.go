package middleware

import (
	"github.com/PedroGao/jerry/libs/erro"
	"github.com/PedroGao/jerry/libs/token"
	"github.com/PedroGao/jerry/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	AuthHeader = "Authorization"
)

// access when user is logined
func LoginRequired(c *gin.Context) {
	authHeader := c.Request.Header.Get(AuthHeader)
	if authHeader == "" {
		c.Error(erro.Forbidden)
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		indentify, err := token.VerifyAccessToken(authHeader)
		if err != nil {
			c.Error(erro.ParamsErr.SetMsg(err.Error()))
			c.Abort()
		} else {
			user := &model.UserModel{
				Username: indentify,
			}
			ok, _ := model.DB.Get(user)
			if !ok {
				c.Error(erro.ParamsErr.SetMsg("用户未找到"))
				c.Abort()
			} else {
				c.Set("user", user)
				c.Next()
			}
		}
	}
}
