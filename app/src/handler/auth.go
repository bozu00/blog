package handler

import (
	"github.com/labstack/echo"
	"html"
	"net/http"

	"../services"
)

var (
	STRETCHCOUNT int = 500
)

func Login(c echo.Context) error {
	email    := html.EscapeString(c.FormValue("email"))
	password := html.EscapeString(c.FormValue("password"))

	if ! services.IsValidUser(email, password) {
		return c.JSON(http.StatusOK, responseWrap(nil, false))
	}
	return c.JSON(http.StatusOK, responseWrap(nil, true))
}


