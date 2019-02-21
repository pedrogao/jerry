package middleware

import (
	"github.com/gin-gonic/gin"
)

const (
	REG_XP = "[{](.*?)[}]"
)

func Logger(template string) func(c *gin.Context) {
	// 1. parse template

	// 2. get vars from user response request

	// 3. combine them into template

	// 4. write into db
	return func(c *gin.Context) {

	}
}
