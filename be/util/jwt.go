package util

import (
	"git.bingyan.net/doc-aid-re-go/config"
	"github.com/golang-jwt/jwt" // 这好像是github.com/dgrijalva/jwt-go的社区维护版
	"github.com/labstack/echo/v4"
	"time"
)

// 待修改
type JWTClaims struct {
	Openid string `json:"openid,omitempty"`
	jwt.StandardClaims
}

// GenerateJWTToken 根据键值对生成jwt token
func GenerateJWTToken(claims JWTClaims) (string, error) {
	claims.Audience = "BingYan"
	claims.ExpiresAt = time.Now().Add(time.Hour * 24 * 14).Unix() // 14天
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(config.C.JWT.Secret))
}

func GetOpenID(c echo.Context) string {
	user := c.Get("user")
	if user == nil {
		return ""
	}
	token := user.(*jwt.Token)
	claims := token.Claims.(*JWTClaims)
	return claims.Openid
}
