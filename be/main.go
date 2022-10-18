package main

import (
	"fmt"
	_ "git.bingyan.net/doc-aid-re-go/middleware"
	_ "git.bingyan.net/doc-aid-re-go/model"
	"git.bingyan.net/doc-aid-re-go/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("hello, this is doc-aid-re-go!")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	router.InitRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
