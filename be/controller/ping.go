package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func TestJWT(c echo.Context) error {
	user := c.Get("user")
	return c.String(http.StatusOK, fmt.Sprintf("Get:%#v", user))
}
