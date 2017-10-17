package handler

import (
	"github.com/labstack/echo"
	"html"
	"net/http"
)

func Login(c echo.Context) error {
	email    := html.EscapeString(c.FormValue("email"))
	password := html.EscapeString(c.FormValue("password"))

	if ! isValidUser(email, password) {
		return c.JSON(http.StatusOK, responseWrap(nil, false))
	}
	return c.JSON(http.StatusOK, responseWrap(nil, true))
}



func isValidUser(email string, password string) bool {

	return true
}
