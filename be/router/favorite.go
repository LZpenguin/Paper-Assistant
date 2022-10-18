package router

import (
	"git.bingyan.net/doc-aid-re-go/controller"
	"github.com/labstack/echo/v4"
)

func InitFavorite(g *echo.Group) {
	g.POST("/favorite/add", controller.CreateFavorite)

	g.POST("/favorite/delete", controller.DeleteFavorite)

	g.POST("/favorite/folder/add", controller.CreateFavFolder)
	g.POST("/favorite/folder/fav", controller.AddFavToFolder)
	//g.POST("/favorite/folder/favs", controller.AddFavsToFolder)
	g.POST("/favorite/folder/delete", controller.DeleteFavorite)
}
