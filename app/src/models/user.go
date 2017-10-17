package models

import (
	_ "github.com/go-sql-driver/mysql"
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
	(email, password, first_name, last_name, sex) Values 
	(?, ?, ?, ?, ?); 
	`

	var user User
	res,err := db.Exec(sql, email, password, first_name, last_name, sex)

	user_id, err := res.LastInsertId()
	if !checkErr(err, "create user failed") {
		return user, err
	}

	err = db.Get(&user, `select id, email, first_name, last_name, created_at, updated_at from users where id=?`, user_id)
	checkErr(err, "select user created fail")

	return user, err
}
