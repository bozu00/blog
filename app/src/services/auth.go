package services


import (
	"log"
	"fmt"
	"time"
	"net/http"
	"strconv"
	"encoding/binary"
	"crypto/rand"
	"crypto/sha256"
	"github.com/labstack/echo"
	"github.com/ipfans/echo-session"

	"../responses"
	"../context"
)

var (
	STRETCHCOUNT int = 500
)

func checkErr(err error, msg string) bool {
	if err != nil {
		// log.Fatalln(msg, err)
		log.Println(msg, err)
		return false
	}
	return true
}


func MakePasswordHashAndSalt(password string) (password_hash string, salt string) {
	salt = makeSalt()
	password_hash = GetPasswordHash(salt, password)
	return
}

func GetPasswordHash(_salt string, raw_password string) string {
	salt := []byte(_salt)
	password := []byte(raw_password)

	hash := []byte("")
	for i := 0; i < STRETCHCOUNT; i++ {
		hasher := sha256.New()
		hasher.Write(salt)
		hasher.Write(password[:])
		hasher.Write(hash[:])
		hash = hasher.Sum(nil)
	}
	password_hash := fmt.Sprintf("%x", hash)
	return password_hash
}

func makeSalt() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}


func AuthFailResponse(err error, res interface{}) responses.Response {
	return responses.Response{
		401,
		"OK",
		res,
	}
}

func Login(c echo.Context) {
  session := session.Default(c)
	cc := c.(*context.CustomContext)
	user_id := cc.Get("UserId")

	session.Set("auth_token", user_id)
	session.Save()
}

func MustAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session := session.Default(c)

	  token, err := c.Cookie("auth_token")
		if err != nil {
		  return c.JSON(http.StatusOK, responses.AuthFailResponse(nil, nil))
		}

		user_id := session.Get(token)
		if user_id == nil {
		  return c.JSON(http.StatusOK, responses.AuthFailResponse(nil, nil))
		} 

		cc := c.(*context.CustomContext)

		cc.Set("UserId", user_id)

		/*
		token_lifetime := decodeTime(v.(string))
		now  := time.Now()
		duration := token_lifetime.Sub(now)

		if duration > 0 { 
			// ここは期限切れ
		  return c.JSON(http.StatusOK, responses.AuthFailResponse(nil, nil))
		}
		*/

		// session.Set("auth_token", count)
		// session.Save()


			// 一致してなければUnauthorizedを返す（ステータスは適当）
			//return echo.NewHTTPError(http.StatusUnauthorized)
		return next(cc)
	}
}

func decodeTime(string_time string) time.Time {	
	format := "Mon Jan 2 15:04:05 MST 2006" 
	time, err := time.Parse(format, string_time)
	if !checkErr(err, "invalid time string in session") {
		// ここを非常に大きな日付にする 2500年くらい
	}

	return time
}


func encodeTime(time time.Time) string {	
	format := "Mon Jan 2 15:04:05 MST 2006" 
	return time.Format(format)
}



