package err

import "fmt"

type HTTPErr struct {
	ErrorCode int    `json:"error_code"` // specify error code
	Msg       string `json:"msg"`        // error message
	Url       string `json:"url"`        // url the user request
	HTTPCode  int    `json:"-"`
}

func NewHTTPErr(errorCode, httpCode int, msg, url string) *HTTPErr {
	return &HTTPErr{
		ErrorCode: errorCode,
		Msg:       msg,
		Url:       url,
		HTTPCode:  httpCode,
	}
}

func (h HTTPErr) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", h.ErrorCode, h.Msg, h.Url)
}

func (h *HTTPErr) SetUrl(url string) *HTTPErr {
	h.Url = url
	return h
}

func (h *HTTPErr) SetHTTPCode(httpCode int) *HTTPErr {
	h.HTTPCode = httpCode
	return h
}

func (h *HTTPErr) Clear() *HTTPErr {
	h.Msg = ""
	return h
}

func (h *HTTPErr) Add(msg string) *HTTPErr {
	h.Msg += " " + msg
	return h
}

func (h *HTTPErr) Addf(format string, args ...interface{}) *HTTPErr {
	h.Msg += " " + fmt.Sprintf(format, args...)
	return h
}
