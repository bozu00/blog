package models

import (
	"database/sql"
	"fmt"
	"gopkg.in/gorp.v1"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func checkErr(err error, msg string) bool {
	if err != nil {
		// log.Fatalln(msg, err)
		log.Println(msg, err)
		return false
	}
	return true
}

func InitDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("mysql", "developer:password@tcp(db:3306)/media")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	return dbmap
}

func MySQLConnect(host string, port int, user string, pass string, dbname string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", user, pass, host, strconv.Itoa(port), dbname))
	checkErr(err, "sqlx.Open failed")
	return db
}

func DBConnect() *sqlx.DB {
	return MySQLConnect("db", 3306, "developer", "password", "media")
}

