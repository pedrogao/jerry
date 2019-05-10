package erro

import "net/http"

const (
	OKCode              = 0
	NoRouteMatchedCode  = 10000
	UnKnownCode         = 10050
	NoMethodMatchedCode = 20000
	UserNotFoundCode    = 30000
	ParamsErrCode       = 40000
	ForbiddenCode       = 60000
	UnauthorizedCode    = 70000
)

var (
	OK              = NewHttpErr(OKCode, http.StatusOK, "OK!")
	NoRouteMatched  = NewHttpErr(NoRouteMatchedCode, http.StatusNotFound, "路由不存在")
	NoMethodMatched = NewHttpErr(NoMethodMatchedCode, http.StatusForbidden, "请求方法不允许")
	UnKnown         = NewHttpErr(UnKnownCode, http.StatusBadRequest, "服务器未知错误")
	UserNotFound    = NewHttpErr(UserNotFoundCode, http.StatusNotFound, "没有找到用户")
	ParamsErr       = NewHttpErr(ParamsErrCode, http.StatusBadRequest, "参数错误")
	Forbidden       = NewHttpErr(ForbiddenCode, http.StatusForbidden, "禁止访问")
	Unauthorized    = NewHttpErr(UnauthorizedCode, http.StatusUnauthorized, "认证失败")
)
