package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var CORSMiddleware echo.MiddlewareFunc

// TODO: 修改CORS策略
func initCors() {
	CORSMiddleware = middleware.CORSWithConfig(middleware.DefaultCORSConfig)
}
