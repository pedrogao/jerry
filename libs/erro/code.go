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
)

var (
	OK              = NewHttpErr(OKCode, http.StatusOK, "OK!")
	NoRouteMatched  = NewHttpErr(NoRouteMatchedCode, http.StatusNotFound, "sorry, there is no route the url matched")
	NoMethodMatched = NewHttpErr(NoMethodMatchedCode, http.StatusForbidden, "sorry, there is no methods the action matched")
	UnKnown         = NewHttpErr(UnKnownCode, http.StatusBadRequest, "sorry, ")
	UserNotFound    = NewHttpErr(UserNotFoundCode, http.StatusNotFound, "sorry, no user found")
	ParamsErr       = NewHttpErr(ParamsErrCode, http.StatusBadRequest, "parameter is erro")
	Forbidden       = NewHttpErr(ForbiddenCode, http.StatusOK, "forbidden！you can't access the resource")
)
