package middleware

import (
	"github.com/PedroGao/jerry/utils"
	"github.com/gin-gonic/gin"
)

var (
	RouteMetaInfos map[string]Meta
)

func init() {
	RouteMetaInfos = map[string]Meta{}
}

type Meta struct {
	Auth, Module string
}

func RouteMeta(auth, module string) func(c *gin.Context) {
	//fmt.Println(len(RouteMetaInfos))

	meta := Meta{
		Auth:   auth,
		Module: module,
	}

	return func(c *gin.Context) {
		// get route name
		name := utils.GetRouteName(c)
		RouteMetaInfos[name] = meta
	}
}
