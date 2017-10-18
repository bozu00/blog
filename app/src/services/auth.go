package services


import (
	"log"
	"fmt"
	"net/http"
	"strconv"
	"encoding/binary"
	"crypto/rand"
	"crypto/sha256"
	"github.com/labstack/echo"
	"github.com/ipfans/echo-session"

	"../responses"
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

func MustAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session := session.Default(c)
		user_id := session.Get("user_id")
		fmt.Println(user_id)
		if user_id == nil {
			// 未ログイン
		  return c.JSON(http.StatusOK, responses.AuthFailResponse(nil, nil))
		} 
		return next(c)
	}
}
