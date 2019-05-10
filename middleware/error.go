package middleware

import (
	"github.com/PedroGao/jerry/libs/erro"
	lv "github.com/PedroGao/jerry/libs/validator"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"unicode"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	// 取最后一个Error为返回的Error
	length := len(c.Errors)
	if length > 0 {
		e := c.Errors[length-1]
		switch e1 := e.Err.(type) {
		case *erro.HttpErr:
			writeHttpError(c, *e1)
		case erro.HttpErr:
			writeHttpError(c, e1)
		case validator.ValidationErrors:
			writeParamError(c, e1)
		case *validator.ValidationErrors:
			writeParamError(c, *e1)
		default:
			writeError(c, e.Err.Error())
		}
	}
}

func writeError(ctx *gin.Context, errString interface{}) {
	status := http.StatusBadRequest
	if ctx.Writer.Status() != http.StatusOK {
		status = ctx.Writer.Status()
	}
	s := erro.UnKnown.SetMsg(errString).SetUrl(ctx.Request.URL.String())
	ctx.JSON(status, s)
}

func writeParamError(ctx *gin.Context, e1 validator.ValidationErrors) {
	mapErrors := make(map[string]string)
	var (
		finalStr string
		s        *erro.HttpErr
	)
	for _, err := range e1 {
		finalStr = err.Translate(lv.Trans)
		runes := []rune(err.StructField())
		runes[0] = unicode.ToLower(runes[0])
		fieldName := string(runes)
		mapErrors[fieldName] = finalStr
	}
	status := http.StatusBadRequest
	if ctx.Writer.Status() != http.StatusOK {
		status = ctx.Writer.Status()
	}
	if len(mapErrors) > 1 {
		s = erro.ParamsErr.SetMsg(mapErrors).SetUrl(ctx.Request.URL.String())
	} else {
		s = erro.ParamsErr.SetMsg(finalStr).SetUrl(ctx.Request.URL.String())
	}
	ctx.JSON(status, s)
}

func writeHttpError(ctx *gin.Context, httpErr erro.HttpErr) {
	s := httpErr.SetUrl(ctx.Request.URL.String())
	ctx.JSON(httpErr.HttpCode, s)
}
