package logic

import (
	"github.com/labstack/echo"
	"strconv"
	"log"
	"../models"
)

func checkErr(err error, msg string) bool {
	if err != nil {
		log.Println(msg, err)
		return false
	}
	return true
}


func Articles(c echo.Context) (interface{}, error) {
	media_id, _ := strconv.Atoi(c.Param("media_id"))
	articles, err := models.GetArticles(media_id, 10, 0)

	checkErr(err, "fetch fail")
	return articles, err
	//checkErr(err, "fetch fail")

	//return c.JSON(http.StatusOK, responses.SafeResponse(err, articles))
}

func Article(c echo.Context) (interface{}, error) {
	media_id, _ := strconv.Atoi(c.Param("media_id"))
	article_id, _ := strconv.Atoi(c.Param("article_id"))
	article, err := models.GetArticle(media_id, article_id)
	checkErr(err, "fetch fail")

	return article, err
	//return c.JSON(http.StatusOK, responses.SafeResponse(err, article))
}
