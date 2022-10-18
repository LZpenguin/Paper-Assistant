package router

import (
	"git.bingyan.net/doc-aid-re-go/controller"
	"git.bingyan.net/doc-aid-re-go/middleware"
	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	InitDefault(e)
	e.Use(middleware.CORSMiddleware)

	e.GET("/document", controller.GetDocument, middleware.JwtMiddelware)

	e.GET("/document/:documentId", controller.GetOneDocumentById, middleware.JwtMiddelware)

	e.GET("/document/search/:keywords", controller.SearchDocumentByKeywords, middleware.JwtMiddelware)

	e.GET("/fav", controller.GetUserFavorite, middleware.JwtMiddelware)

	e.DELETE("/fav/:folderName", controller.DeleteFavoritesFromFolder, middleware.JwtMiddelware)

	e.POST("/fav/:folderName", controller.PostFavoritesToFolder, middleware.JwtMiddelware)

	e.PUT("/folder", controller.SortFolder, middleware.JwtMiddelware)

	e.DELETE("/folder/:folderName", controller.DeleteFolder, middleware.JwtMiddelware)

	e.GET("/folder/:folderName", controller.GetAllURLFromFolder, middleware.JwtMiddelware)

	e.POST("/folder/:folderName", controller.PostFolder, middleware.JwtMiddelware)

	e.DELETE("/keyword", controller.DeleteKeyword, middleware.JwtMiddelware)

	e.GET("/keyword", controller.GetKeywords, middleware.JwtMiddelware)

	e.POST("/keyword", controller.PostKeywords, middleware.JwtMiddelware)

	e.GET("/keyword/search/:keywords", controller.SearchKeywords, middleware.JwtMiddelware)

	e.GET("/magazine", controller.GetMagazine)

	e.GET("/magazine/:issueName/:topicName", controller.GetMagazinesWithTopic, middleware.JwtMiddelware)

	e.GET("/magazine/:magazineName", controller.GetMagazineInfoWithMagazineName, middleware.JwtMiddelware)

	e.DELETE("/sub", controller.DeleteSub, middleware.JwtMiddelware)

	e.GET("/sub", controller.GetSub, middleware.JwtMiddelware)

	e.POST("/sub/:magazineName", controller.SubMagazineWithMagazineName, middleware.JwtMiddelware)

	e.POST("/user/feedback", controller.GatherFeedback, middleware.JwtMiddelware)

	// 爬虫用
	e.POST("/magazine", controller.CreateMagazine)
	e.POST("/document", controller.PostDocument)

	e.POST("/user/login", controller.UserLogin)

	var g *echo.Group = e.Group("/debug")
	//if !config.C.EnableDebug {
	//	g = e.Group(config.C.AppInfo.APIPrefix)
	//} else {
	//	log.Println("Debug mod is in use")
	//	g = e.Group("")
	//}
	//g.Use(middleware.JwtMiddelware)
	InitUser(g)
	InitPaper(g)
	InitKeyword(g)
	InitMagazine(g)
	InitFavorite(g)
}
func InitDefault(e *echo.Echo) {
	e.GET("/ping", controller.Ping)
	e.GET("/user/debug", controller.GetJWT)
}
