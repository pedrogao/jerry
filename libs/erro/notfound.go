package erro

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// NotFound creates a gin middleware for handling page not found.
func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := NoRouteMatched.SetUrl(c.Request.URL.String())
		c.JSON(http.StatusNotFound, s)
	}
}
