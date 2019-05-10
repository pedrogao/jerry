package token

import (
	"errors"
	"github.com/spf13/viper"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"strings"
	"time"
)

const (
	ACCESS  = "ACCESS"
	REFRESH = "REFRESH"
)

var (
	JwtInstance *Jwt
)

type Jwt struct {
	Key                    []byte
	AccessTokenExpireTime  time.Duration
	RefreshTokenExpireTime time.Duration
}

func New() {
	secret := viper.GetString("secret")
	accessExpire := viper.GetInt("access_expire")
	JwtInstance = &Jwt{
		Key:                    []byte(secret),
		AccessTokenExpireTime:  time.Hour * time.Duration(accessExpire),
		RefreshTokenExpireTime: time.Hour * 24 * 30,
	}
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
	return j.generateToken(j.AccessTokenExpireTime, ACCESS, identify)
}

func GenerateAccessToken(identify string) (string, error) {
	return JwtInstance.GenerateAccessToken(identify)
}

func (j *Jwt) GenerateRefreshToken(identify string) (string, error) {
	return j.generateToken(j.RefreshTokenExpireTime, REFRESH, identify)
}

func (j *Jwt) generateToken(timeOut time.Duration, tokenType, identify string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expire := time.Now().Add(timeOut)
	// exp express token expire time
	claims["exp"] = expire.Unix()
	claims["type"] = tokenType
	// identify express username
	claims["identify"] = identify
	// get the exp time
	//exp := int64(claims["exp"].(float64))
	return token.SignedString(j.Key)
}

func (j *Jwt) VerifyAccessTokenInHeader(authHeader string) (jwt.MapClaims, error) {
	return j.verifyTokenInHeader(authHeader, ACCESS)
}

func VerifyAccessTokenInHeader(authHeader string) (jwt.MapClaims, error) {
	return JwtInstance.VerifyAccessTokenInHeader(authHeader)
}

func (j *Jwt) VerifyRefreshToken(authHeader string) (jwt.MapClaims, error) {
	return j.verifyToken(authHeader, REFRESH)
}

func (j *Jwt) verifyTokenInHeader(authHeader, tokenType string) (jwt.MapClaims, error) {
	var (
		tokenStr string
	)
	parts := strings.SplitN(authHeader, " ", 2)

	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, errors.New("请填写正确的Bearer字段")
	}
	tokenStr = parts[1]
	return j.verifyToken(tokenStr, tokenType)
}

func (j *Jwt) verifyToken(tokenStr, tokenType string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		//if jwt.SigningMethodHS256 != t.Method {
		//	return nil, errors.New("token解析错误")
		//}
		claims := t.Claims.(jwt.MapClaims)
		innerTokenType := claims["type"].(string)
		if innerTokenType != tokenType {
			return nil, errors.New("令牌类型错误")
		}
		return j.Key, nil
	})
	if token.Valid {
		return token.Claims.(jwt.MapClaims), nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New("令牌损坏")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return nil, errors.New("令牌过期")
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
