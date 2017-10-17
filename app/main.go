package main

import (
	"github.com/labstack/echo"

	"./src/handler"
)




func main() {
	e := echo.New()

	e.GET("/", handler.Hello)
    e.GET("/api/:media_id/articles", handler.Articles)
    e.GET("/api/:media_id/article/:article_id", handler.Article)

    e.POST("/api/login",       handler.Login)
    e.POST("/api/create_user", handler.CreateUser)

	e.Logger.Fatal(e.Start(":1323"))
}
