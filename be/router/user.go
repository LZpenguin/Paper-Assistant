package router

import (
	"git.bingyan.net/doc-aid-re-go/controller"
	"github.com/labstack/echo/v4"
)

func InitUser(eg *echo.Group) {

	ug := eg.Group("/user")
	initUserDebug(ug)
	ug.POST("/login", controller.Login)
	ug.GET("/debug", controller.GetJWT)
}
func initUserDebug(ug *echo.Group) {
	ug.POST("/add", controller.CreateUser)
	ug.GET("/all", controller.ListAllUsers)
}
