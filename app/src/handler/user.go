package handler

import (
	"github.com/labstack/echo"
	"html"
	"net/http"
	"strconv"
	"github.com/ipfans/echo-session"
	"time"
	"fmt"

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
    session := session.Default(c)
	user_id := models.GetUserIdByEmail(email)
	token_cookie, err := c.Cookie("auth_token")
    token := token_cookie.Value
	if err != nil {
		return err
	}
	fmt.Println(token)
	session.Set(token, user_id)
	session.Save()
	
	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = token 
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)


	return c.JSON(http.StatusOK, responses.SafeResponse(nil, true))
}

