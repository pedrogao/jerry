package router

import (
	"github.com/PedroGao/jerry/controller/web"
	"github.com/PedroGao/jerry/libs/err"
	"github.com/PedroGao/jerry/middleware"
	"github.com/gin-gonic/gin"
)

func Load(app *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// handle no route
	app.NoRoute(func(c *gin.Context) {
		c.JSON(err.NoRouteMatched.HTTPCode, err.NoRouteMatched.SetUrl(c.Request.URL.String()))
	})

	// handle no method
	app.NoMethod(func(c *gin.Context) {
		c.JSON(err.NoMethodMatched.HTTPCode, err.NoMethodMatched.SetUrl(c.Request.URL.String()))
	})

	// apply middleware
	app.Use(middleware.CORS)
	app.Use(middleware.NoCache)
	app.Use(middleware.Secure)
	app.Use(mw...)

	//app.Use(middleware.ErrHandler)

	// mount routes
	// Example for binding JSON ({"user": "manu", "password": "123"})
	app.POST("/login", middleware.RouteMeta("登陆", "用户"), web.Login)

	user := app.Group("/user")
	user.GET("/", middleware.RouteMeta("查询所有用户", "用户"), middleware.LoginRequired, web.GetUsers)

	book := app.Group("/book")
	book.GET("/search", middleware.RouteMeta("搜索书籍", "书籍"), web.SearchBook)

	book.GET("/id/:id",
		middleware.RouteMeta("查询书籍", "书籍"),
		web.GetBook,
		middleware.Logger("hhhh"))

	return app
}
