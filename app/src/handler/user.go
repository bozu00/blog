package handler

import (
	model "../models"
	"github.com/labstack/echo"
	"html"
	"net/http"
	"strconv"
	"fmt"
)


func CreateUser(c echo.Context) error {
	email      := html.EscapeString(c.FormValue("email"))
	password   := html.EscapeString(c.FormValue("password"))
	first_name := html.EscapeString(c.FormValue("first_name"))
	last_name  := html.EscapeString(c.FormValue("last_name"))
	sex,err    := strconv.Atoi(html.EscapeString(c.FormValue("sex")))
	fmt.Println( c.FormValue("email") )
	fmt.Println( c.FormValue("sex") )
	fmt.Println( html.EscapeString(c.FormValue("sex")) )

	if !checkErr(err, "fail validate sex") {
		return c.JSON(http.StatusOK, responseWrap(err, nil))
	}

	user, err := model.CreateUser(email, password, first_name, last_name, sex)
	checkErr(err, "create user fail")

	return c.JSON(http.StatusOK, responseWrap(err, user))
}

