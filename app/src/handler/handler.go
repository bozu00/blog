package handler

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func checkErr(err error, msg string) bool {
	if err != nil {
		// log.Fatalln(msg, err)
		log.Println(msg, err)
		return false
	}
	return true
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

