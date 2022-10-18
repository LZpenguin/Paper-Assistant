package router

import (
	"git.bingyan.net/doc-aid-re-go/controller"
	"github.com/labstack/echo/v4"
)

func InitMagazine(g *echo.Group) {
	g.GET("/magazine/:magazineid", controller.GetOneMagazine)
	g.GET("/magazine/delete/:magazineid", controller.DeleteOneMagazine)
	g.POST("/magazine/add", controller.AddOneMagazine)
}
