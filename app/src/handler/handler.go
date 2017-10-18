package handler

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"github.com/ipfans/echo-session"
	"../responses"
)

func checkErr(err error, msg string) bool {
	if err != nil {
		log.Println(msg, err)
		return false
	}
	return true
}


// Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Private(c echo.Context) error {
	return c.String(http.StatusOK, "Private, World!")
}

func CountUp(c echo.Context) error {
	session := session.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count += 1
	}
	session.Set("count", count)
	session.Save()
	return c.JSON(http.StatusOK, responses.SafeResponse(nil, count))
}
