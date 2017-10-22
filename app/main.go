package main

import (
	"github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"


	"./src/handler"
	"./src/logic"
	"./src/services"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/**.html")),
	}
    e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// for swagger
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:8888"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	store, err := session.NewRedisStore(32, "tcp", "session-redis:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	e.Use(session.Sessions("auth_token", store))

	// e.GET("/", handler.Hello)
	e.GET("/private", services.MustAuth(handler.Private))
	e.GET("/countup", handler.CountUp)


	e.GET("/media/:media_id/api/hello",    handler.Hello)

	//e.GET("/media/:media_id/articles",     handler.Html(logic.Articles, logic.Articles))
	e.GET("/media/:media_id/articles",     handler.Articles)
	e.GET("/media/:media_id/api/articles", handler.Json(logic.Articles))

	e.GET("/media/:media_id/api/article/:article_id", handler.Json(logic.Articles))

	e.POST("/media/:media_id/api/login", handler.Login)
	e.POST("/media/:media_id/api/logout", handler.Logout)
	e.POST("/media/:media_id/api/create_user", handler.CreateUser)

	e.Logger.Fatal(e.Start(":1323"))
}
