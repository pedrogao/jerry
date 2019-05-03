package token

import (
	"errors"
	"github.com/spf13/viper"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"strings"
	"time"
)

var (
	JwtInstance *Jwt
)

type Jwt struct {
	Key                    []byte
	AccessTokenExpireTime  time.Duration
	RefreshTokenExpireTime time.Duration
}

func NewJwt(key []byte, accessTokenExpireTime, refreshTokenExpireTime time.Duration) *Jwt {
	return &Jwt{
		Key:                    key,
		AccessTokenExpireTime:  accessTokenExpireTime,
		RefreshTokenExpireTime: refreshTokenExpireTime,
	}
}

func init() {
	secret := viper.GetString("secret")
	accessExpire := viper.GetInt("access_expire")
	JwtInstance = NewJwt([]byte(secret), time.Hour*time.Duration(accessExpire), time.Hour*24*30)
}

func (j *Jwt) GenerateTokens(identify string) (string, string, error) {
	var (
		accessToken, refreshToken string
		err                       error
	)
	accessToken, err = j.GenerateAccessToken(identify)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = j.GenerateRefreshToken(identify)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (j *Jwt) GenerateAccessToken(identify string) (string, error) {
	return j.generateToken(j.AccessTokenExpireTime, identify)
}

func GenerateAccessToken(identify string) (string, error) {
	return JwtInstance.GenerateAccessToken(identify)
}

func (j *Jwt) GenerateRefreshToken(identify string) (string, error) {
	return j.generateToken(j.RefreshTokenExpireTime, identify)
}

func (j *Jwt) generateToken(timeOut time.Duration, identify string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expire := time.Now().Add(timeOut)
	// exp express token expire time
	claims["exp"] = expire.Unix()
	// identify express username
	claims["identify"] = identify
	// get the exp time
	//exp := int64(claims["exp"].(float64))
	return token.SignedString(j.Key)
}

func (j *Jwt) VerifyAccessToken(authHeader string) (string, error) {
	return j.verifyToken(authHeader, j.AccessTokenExpireTime)
}

func VerifyAccessToken(authHeader string) (string, error) {
	return JwtInstance.VerifyAccessToken(authHeader)
}

func (j *Jwt) VerifyRefreshToken(authHeader string) (string, error) {
	return j.verifyToken(authHeader, j.RefreshTokenExpireTime)
}

func (j *Jwt) verifyToken(authHeader string, timeOut time.Duration) (string, error) {
	var (
		tokenStr string
	)
	parts := strings.SplitN(authHeader, " ", 2)

	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("请填写正确的Bearer字段")
	}
	tokenStr = parts[1]
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != t.Method {
			return nil, errors.New("token解析错误")
		}
		return j.Key, nil
	})
	if err != nil {
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	expire := int64(claims["exp"].(float64))

	if expire > time.Now().Add(timeOut).Unix() {
		return "", errors.New("令牌过期")
	}
	return claims["identify"].(string), nil
}

func (j *Jwt) VerifySingleToken(tokenStr string) (string, error) {
	return j.verifySingleToken(tokenStr, j.AccessTokenExpireTime)
}

func (j *Jwt) verifySingleToken(tokenStr string, timeOut time.Duration) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != t.Method {
			return nil, errors.New("token解析错误")
		}
		return j.Key, nil
	})
	if err != nil {
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	expire := int64(claims["exp"].(float64))

	if expire > time.Now().Add(timeOut).Unix() {
		return "", errors.New("令牌过期")
	}
	return claims["identify"].(string), nil
}
