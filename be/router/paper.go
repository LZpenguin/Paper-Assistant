package router

import (
	"git.bingyan.net/doc-aid-re-go/controller"
	"github.com/labstack/echo/v4"
)

func InitPaper(g *echo.Group) {
	g.GET("/document/:documentid", controller.GetOnePaper)
	g.GET("/document/delete/:documentid", controller.DeleteOnePaper)
	g.POST("/document/add", controller.AddOnePaper)
}
