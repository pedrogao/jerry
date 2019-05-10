package token

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestGenerateAccessToken(t *testing.T) {
	viper.Set("secret", "\x88W\xf09\x91\x07\x98\x89\x87\x96\xa0A\xc68\xf9\xecJJU\x17\xc5V\xbe\x8b\xef\xd7\xd8\xd3\xe6\x95*4")
	viper.Set("access_expire", 3)
	token, e := GenerateAccessToken("pedro")
	if e != nil {
		assert.Error(t, e)
	}
	log.Println(token)
	assert.NotEqual(t, token, "")
}

func TestGenerateAccessToken2(t *testing.T) {
	jwt := NewJwt([]byte("sssss"), 3*time.Second, time.Second)
	token, e := jwt.GenerateAccessToken("pedro")
	if e != nil {
		log.Println(e)
	}
	log.Println(token)
	assert.NotEqual(t, token, "")
	time.Sleep(2 * time.Second)
	claims, e := jwt.verifyToken(token, ACCESS)
	if e != nil {
		log.Println(e)
	}
	log.Println(claims)
}
