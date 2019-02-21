package test

import (
	"github.com/PedroGao/jerry/router"
	"github.com/gin-gonic/gin"
)

func setupApp() *gin.Engine {
	app := gin.Default()

	// load middleware and routes
	router.Load(app)

	// test api
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "greeting from pedro",
		})
	})

	return app
}
