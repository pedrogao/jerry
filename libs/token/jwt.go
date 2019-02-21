package token

import (
	"errors"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"strings"
	"time"
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

var (
	JwtInstance *Jwt
)

func init() {
	JwtInstance = NewJwt([]byte("jhhlXjjokkjhiopipigio"), time.Hour, time.Hour*24*30)
}

func (j *Jwt) GenerateTokens(identify string) (string, string, error) {
	var (
		accessToken, refreshToken string
		err                       error
	)
	accessToken, err = j.generateAccessToken(identify)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = j.generateRefreshToken(identify)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (j *Jwt) generateAccessToken(identify string) (string, error) {
	return j.generateToken(j.AccessTokenExpireTime, identify)
}

func (j *Jwt) generateRefreshToken(identify string) (string, error) {
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

func (j *Jwt) VerifyRefreshToken(authHeader string) (string, error) {
	return j.verifyToken(authHeader, j.RefreshTokenExpireTime)
}

func (j *Jwt) verifyToken(authHeader string, timeOut time.Duration) (string, error) {
	var (
		tokenStr string
	)
	parts := strings.SplitN(authHeader, " ", 2)

	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("请求认证的header字段错误")
	}
	tokenStr = parts[1]
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != t.Method {
			return nil, errors.New("生成token的哈希函数不一致")
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
			return nil, errors.New("生成token的哈希函数不一致")
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
