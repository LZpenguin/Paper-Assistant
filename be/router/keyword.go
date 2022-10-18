package router

import (
	"git.bingyan.net/doc-aid-re-go/controller"
	"github.com/labstack/echo/v4"
)

func InitKeyword(g *echo.Group) {
	g.GET("/keyword/:keywordid", controller.GetOneKeyword)
	g.GET("/keyword/delete/:keywordid", controller.DeleteOneKeyword)
	g.POST("/keyword/add", controller.CreateOneKeyword)

	g.GET("/keyword/all", controller.GetAllKeywords)
	g.POST("/keyword/addmany", controller.CreateManyKeyWords)
}
