package middleware

import (
	"fmt"
	"github.com/PedroGao/jerry/libs/erro"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"strings"
	"unicode"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	// 取最后一个Error为返回的Error
	length := len(c.Errors)
	if length > 0 {
		e := c.Errors[length-1]
		switch e.Type {
		case gin.ErrorTypeBind:
			errs, ok := e.Err.(validator.ValidationErrors)
			if !ok {
				writeError(c, e.Error())
				return
			}
			var stringErrors []string
			for _, err := range errs {
				stringErrors = append(stringErrors, validationErrorToText(err))
			}
			writeError(c, strings.Join(stringErrors, "; "))
		case gin.ErrorTypePrivate:
			httpErr, ok := e.Err.(erro.HttpErr)
			if !ok {
				httpErr, ok := e.Err.(*erro.HttpErr)
				if !ok {
					writeError(c, e.Error())
					return
				}
				writeHttpError(c, *httpErr)
				return
			}
			writeHttpError(c, httpErr)
		default:
			writeError(c, e.Err.Error())
		}
	}
}

func validationErrorToText(e *validator.FieldError) string {
	runes := []rune(e.Field)
	runes[0] = unicode.ToLower(runes[0])
	fieldName := string(runes)
	switch e.Tag {
	case "required":
		return fmt.Sprintf("Field '%s' is required", fieldName)
	case "max":
		return fmt.Sprintf("Field '%s' must be less or equal to %s", fieldName, e.Param)
	case "min":
		return fmt.Sprintf("Field '%s' must be more or equal to %s", fieldName, e.Param)
	}
	return fmt.Sprintf("Field '%s' is not valid", fieldName)
}

func writeError(ctx *gin.Context, errString string) {
	status := http.StatusBadRequest
	if ctx.Writer.Status() != http.StatusOK {
		status = ctx.Writer.Status()
	}
	s := erro.UnKnown.SetMsg(errString).SetUrl(ctx.Request.URL.String())
	ctx.JSON(status, s)
}

func writeHttpError(ctx *gin.Context, httpErr erro.HttpErr) {
	s := httpErr.SetUrl(ctx.Request.URL.String())
	ctx.JSON(httpErr.HttpCode, s)
}
