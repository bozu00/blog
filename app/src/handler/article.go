package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"../models"
)

func Articles(c echo.Context) error {
	media_id, _ := strconv.Atoi(c.Param("media_id"))
	articles, err := models.GetArticles(media_id, 10, 0)
	checkErr(err, "fetch fail")

	return c.JSON(http.StatusOK, responseWrap(err, articles))
}

func Article(c echo.Context) error {
	media_id, _ := strconv.Atoi(c.Param("media_id"))
	article_id, _ := strconv.Atoi(c.Param("article_id"))
	article, err := models.GetArticle(media_id, article_id)
	checkErr(err, "fetch fail")

	return c.JSON(http.StatusOK, responseWrap(err, article))
}
