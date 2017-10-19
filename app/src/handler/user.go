package handler

import (
	"github.com/labstack/echo"
	"html"
	"net/http"
	"strconv"
	"github.com/ipfans/echo-session"
	"log"


	"../models"
	"../responses"
)


func CreateUser(c echo.Context) error {
	email      := html.EscapeString(c.FormValue("email"))
	password   := html.EscapeString(c.FormValue("password"))
	first_name := html.EscapeString(c.FormValue("first_name"))
	last_name  := html.EscapeString(c.FormValue("last_name"))
	sex,err    := strconv.Atoi(html.EscapeString(c.FormValue("sex")))

	if !checkErr(err, "fail validate sex") {
		return c.JSON(http.StatusOK, responses.SafeResponse(err, nil))
	}

	user, err := models.CreateUser(email, password, first_name, last_name, sex)
	checkErr(err, "create user fail")

	return c.JSON(http.StatusOK, responses.SafeResponse(err, user))
}


func Login(c echo.Context) error {
	email    := html.EscapeString(c.FormValue("email"))
	password := html.EscapeString(c.FormValue("password"))

	if ! models.IsValidPassword(email, password) {
		return c.JSON(http.StatusOK, responses.SafeResponse(nil, false))
	}

	// ログイン処理
	user_id := models.GetUserIdByEmail(email)
    session := session.Default(c)
	session.Set("user_id", user_id)
	session.Save()
	
	return c.JSON(http.StatusOK, responses.SafeResponse(nil, true))
}


func Logout(c echo.Context) error {
    session := session.Default(c)
	_user_id := session.Get("user_id")
	if _user_id == nil {
		return c.JSON(http.StatusOK, responses.SafeResponse(nil, false))
	}
	user_id := _user_id.(int)
	log.Printf("user %v logout\n", user_id)
	session.Delete("user_id")
	session.Save()
	
	return c.JSON(http.StatusOK, responses.SafeResponse(nil, true))
}

