package erro

import (
	"fmt"
	"github.com/json-iterator/go"
)

type HttpErr struct {
	ErrorCode int    `json:"error_code"` // specify error code
	Msg       string `json:"msg"`        // error message
	Url       string `json:"url"`        // url the user request
	HttpCode  int    `json:"-"`
}

func NewHttpErr(errorCode, httpCode int, msg string) *HttpErr {
	return &HttpErr{
		ErrorCode: errorCode,
		Msg:       msg,
		HttpCode:  httpCode,
	}
}

func (h HttpErr) Error() string {
	return h.Msg
}

func (h *HttpErr) SetUrl(url string) *HttpErr {
	h.Url = url
	return h
}

func (h *HttpErr) SetHttpCode(httpCode int) *HttpErr {
	h.HttpCode = httpCode
	return h
}

func (h *HttpErr) Clear() *HttpErr {
	h.Msg = ""
	return h
}

func (h *HttpErr) SetMsg(msg string) *HttpErr {
	h.Msg = msg
	return h
}

func (h *HttpErr) Add(msg string) *HttpErr {
	h.Msg += " " + msg
	return h
}

func (h *HttpErr) Addf(format string, args ...interface{}) *HttpErr {
	h.Msg += " " + fmt.Sprintf(format, args...)
	return h
}

func (h *HttpErr) Json(url string) string {
	h.SetUrl(url)
	s, err := jsoniter.MarshalToString(h)
	if err != nil {
		// 如果json解析错误，返回空字符串
		return ""
	}
	return s
}
