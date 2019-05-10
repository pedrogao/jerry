package middleware

import (
	"github.com/PedroGao/jerry/libs/erro"
	"github.com/PedroGao/jerry/libs/token"
	"github.com/PedroGao/jerry/model"
	"github.com/gin-gonic/gin"
)

const (
	AuthHeader = "Authorization"
)

// authorization middleware, access the resource if token is correct or not
func LoginRequired(c *gin.Context) {
	authHeader := c.Request.Header.Get(AuthHeader)
	if authHeader == "" {
		c.Error(erro.Unauthorized)
		c.Abort()
	} else {
		claims, err := token.VerifyAccessTokenInHeader(authHeader)
		if err != nil {
			c.Error(erro.Unauthorized.SetMsg(err.Error()))
			c.Abort()
		} else {
			indentify := claims["identify"].(string)
			user := &model.UserModel{
				Username: indentify,
			}
			ok, _ := model.DB.Get(user)
			if !ok {
				c.Error(erro.Unauthorized.SetMsg("用户未找到"))
				c.Abort()
			} else {
				c.Set("user", user)
				c.Next()
			}
		}
	}
}
