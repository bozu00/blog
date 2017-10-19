package main

import (
	"github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"./src/handler"
	"./src/services"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// for swagger
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:8888"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	store, err := session.NewRedisStore(32, "tcp", "redis:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	e.Use(session.Sessions("auth_token", store))

	e.GET("/", handler.Hello)
	e.GET("/private", services.MustAuth(handler.Private))
	e.GET("/countup", handler.CountUp)
	e.GET("/api/:media_id/articles", handler.Articles)
	e.GET("/api/:media_id/article/:article_id", handler.Article)

	e.POST("/api/login", handler.Login)
	e.POST("/api/logout", handler.Logout)
	e.POST("/api/create_user", handler.CreateUser)

	e.Logger.Fatal(e.Start(":1323"))
}
