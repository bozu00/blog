package services


import (
	"crypto/sha256"
	"log"
	"encoding/binary"
	"crypto/rand"
	"strconv"
	"../models"
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

func IsValidUser(email string, password string) bool {
	pass, salt, err := CheckUserExists(email)
	if !checkErr(err, "select salt password failed") {
		return false
	}
	 
	password_hash := getPasswordHash(salt, password)
	if pass != password_hash {
		return false
	}

	return true
}


func getPasswordHash(_salt string, raw_password string) string {
	salt := []byte(_salt)
	password := []byte(raw_password)

	hash := []byte("")
	for i := 0; i < STRETCHCOUNT; i++ {
		hasher := sha256.New()
		hasher.Write(salt[:])
		hasher.Write(password[:])
		hasher.Write(hash[:])
		hash = hasher.Sum(nil)
	}

	password_hash := string(hash)
	return password_hash
}

func makeSalt() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func CheckUserExists(email string) (password string, salt string, err error) {
	db := models.DBConnect()
	defer db.Close()

	type pass_salt struct {
		password string
		salt     string
	}
	sql := `
	select password, salt from users where email=? limit 1;
	`
	var res pass_salt
	err = db.Get(&res, sql, email)

	if !checkErr(err, "user does not exits") {
		return "", "", err
	}

	return res.password, res.salt, err
}

