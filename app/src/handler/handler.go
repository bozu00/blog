package handler

import (
	"net/http"
	"github.com/labstack/echo"
	model "../models"
	"strconv"
	"log"
)


func checkErr(err error, msg string) {
    if err != nil {
        // log.Fatalln(msg, err)
		log.Println(msg,err)
    }
}


type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}
func responseWrap(err error, res interface{}) Response {
	return Response{
		200,
		"OK",
		res,
	}
}

// Handler
func Hello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
}

func Articles(c echo.Context) error {
    media_id, _ := strconv.Atoi(c.Param("media_id"))
	articles, err := model.GetArticles(media_id, 10, 0)
	checkErr(err, "fetch fail")

	return c.JSON(http.StatusOK, responseWrap(err, articles))
}

func Article(c echo.Context) error {
    media_id, _ := strconv.Atoi(c.Param("media_id"))
    article_id, _ := strconv.Atoi(c.Param("article_id"))
	article, err := model.GetArticle(media_id, article_id)
	checkErr(err, "fetch fail")
	
	return c.JSON(http.StatusOK, responseWrap(err, article))
}
