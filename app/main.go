package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ipfans/echo-session"

	"./src/handler"
	"./src/services"
	"./src/context"
)


func main() {
	e := echo.New()

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.CustomContext{c, -1} // -1は存在しないユーザ
			return h(cc)
		}
	})


	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	store, err := session.NewRedisStore(32, "tcp", "redis:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	e.Use(session.Sessions("auth_token", store))


	e.GET("/"                                 , handler.Hello)
	e.GET("/private"                          , services.MustAuth(handler.Private))
	e.GET("/countup"                          , handler.CountUp)
    e.GET("/api/:media_id/articles"           , handler.Articles)
    e.GET("/api/:media_id/article/:article_id", handler.Article)

    e.POST("/api/login"                       , handler.Login)
    e.POST("/api/create_user"                 , handler.CreateUser)

	e.Logger.Fatal(e.Start(":1323"))
}
