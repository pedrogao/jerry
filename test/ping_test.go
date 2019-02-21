package test

import (
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPing(t *testing.T) {
	app := setupApp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bytes := w.Body.Bytes()
	msg := jsoniter.Get(bytes, "msg")
	assert.Equal(t, "greeting from pedro", msg.ToString())
}

func TestUserLogin(t *testing.T) {
	app := setupApp()

	w := httptest.NewRecorder()

	s, _ := jsoniter.MarshalToString(map[string]string{
		"nickname": "pedro",
		"password": "123456",
	})
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(s))
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bytes := w.Body.Bytes()
	access := jsoniter.Get(bytes, "access_token")
	refresh := jsoniter.Get(bytes, "refresh_token")
	assert.NotEqual(t, "", refresh.ToString())
	assert.NotEqual(t, "greeting from pedro", access.ToString())
}
