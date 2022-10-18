package middleware

import (
	"git.bingyan.net/doc-aid-re-go/config"
	"git.bingyan.net/doc-aid-re-go/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var JwtMiddelware echo.MiddlewareFunc

func initJwt() {
	c := middleware.JWTConfig{
		SigningKey: []byte(config.C.JWT.Secret),
		Claims:     &util.JWTClaims{},
	}
	JwtMiddelware = middleware.JWTWithConfig(c)
}
