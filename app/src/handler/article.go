package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"log"
	//"reflect"
	"../models"
	"../responses"
)

func Json(handler func(echo.Context) (interface{}, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := handler(c)
		log.Printf("JSON:エラーをチェックしてここでレスポンスを決める\n")
		return c.JSON(http.StatusOK, responses.SafeResponse(err, res)) 
	}
}

func Html(handler func(echo.Context) (interface{}, error), get_tmpl func(echo.Context) (interface{}, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		//res, err := handler(c)

		//v := reflect.ValueOf(res)

		//data := res.(v.Type())
		//
		//log.Println(v)
		//log.Printf("HTML:エラーをチェックしてここでレスポンスを決める.\n")
		// tmpl := gettemplate(c)
		return c.Render(http.StatusOK, "hello", "world")
	}
}

func ArticlesJson(c echo.Context) error {
	media_id, _ := strconv.Atoi(c.Param("media_id"))
	articles, err := models.GetArticles(media_id, 10, 0)
	checkErr(err, "fetch fail")

	return c.JSON(http.StatusOK, responses.SafeResponse(err, articles))
}

func Articles(c echo.Context) error {
	//media_id, _ := strconv.Atoi(c.Param("media_id"))
	//articles, err := models.GetArticles(media_id, 10, 0)
	//checkErr(err, "fetch fail")

	return c.Render(http.StatusOK, "test", "world")
}

func Article(c echo.Context) error {
	media_id, _ := strconv.Atoi(c.Param("media_id"))
	article_id, _ := strconv.Atoi(c.Param("article_id"))
	article, err := models.GetArticle(media_id, article_id)
	checkErr(err, "fetch fail")

	return c.JSON(http.StatusOK, responses.SafeResponse(err, article))
}
