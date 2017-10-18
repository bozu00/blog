package models

import (
	_ "github.com/go-sql-driver/mysql"

	"../services"
)

type User struct {
	Id        int64  `db:"id"`
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func CreateUser(email string, password string, first_name string, last_name string, sex int) (User, error) {
	db := DBConnect()
	defer db.Close()

	sql := `
	insert into users 
	(email, password, first_name, last_name, sex, salt) Values 
	(?, ?, ?, ?, ?, ?); 
	`

	password_hash, salt := services.MakePasswordHashAndSalt(password) 
	var user User
	res,err := db.Exec(sql, email, password_hash, first_name, last_name, sex, salt)

	user_id, err := res.LastInsertId()
	if !checkErr(err, "create user failed") {
		return user, err
	}

	err = db.Get(&user, `select id, email, first_name, last_name, created_at, updated_at from users where id=?`, user_id)
	checkErr(err, "select user created fail")

	return user, err
}


func GetUserIdByEmail(email string) (user_id int) {
	db := DBConnect()
	defer db.Close()
	_ = db.Get(&user_id, `select id from users where email=?`, user_id)
	return
}


func IsValidPassword(email string, password string) bool {
	pass, salt, err := checkUserExists(email)
	if !checkErr(err, "select salt password failed") {
		return false
	}
	 
	password_hash := services.GetPasswordHash(salt, password)
	if pass != password_hash {
		return false
	}

	return true
}

func checkUserExists(email string) (password string, salt string, err error) {
	db := DBConnect()
	defer db.Close()

	type PassSalt struct {
		Password string `db:"password"`
		Salt     string `db:"salt"`
	}
	sql := `
	select password, salt from users where email=? limit 1;
	`
	var res PassSalt
	err = db.Get(&res, sql, email)

	if !checkErr(err, "user does not exits") {
		return "", "", err
	}

	return res.Password, res.Salt, err
}


