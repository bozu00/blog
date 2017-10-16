package models

import (
    "database/sql"
    "gopkg.in/gorp.v1"
    "log"
    "fmt"
    "strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)


func checkErr(err error, msg string) {
    if err != nil {
        // log.Fatalln(msg, err)
        log.Println(msg, err)
    }
}

func InitDb() *gorp.DbMap {
    // connect to db using standard Go database/sql API
    // use whatever database/sql driver you wish
    db, err := sql.Open("mysql", "developer:password@tcp(localhost:3306)/media")
    checkErr(err, "sql.Open failed")

    // construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

    return dbmap
}

func MySQLConnect(host string, port int, user string, pass string, dbname string) *sqlx.DB {
	db, err := sqlx.Connect("mysql",fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",user, pass, host, strconv.Itoa(port), dbname))
    checkErr(err, "sqlx.Open failed")
	return db
}

func DBConnect() *sqlx.DB {
	return MySQLConnect("localhost", 3306, "developer", "password", "media")
}


type Article struct {
    Id                int64      `db:"id"`
    Title             string     `db:"title"`
    Content           string     `db:"content"`
    AuthorId          int64      `db:"author_id"`
    AuthorName        string     `db:"author_name"`
	ArticleCategoryId int64      `db:"article_category_id"`
	MediaId		      int64      `db:"media_id"`
	CreatedAt         string     `db:"created_at"`
	UpdatedAt         string     `db:"updated_at"`
}


type ArticleWithAuthor struct {
    Id                int64      `db:"id"`
    Title             string     `db:"title"`
    Content           string     `db:"content"`
	ArticleCategoryId int64      `db:"article_category_id"`
	MediaId		      int64      `db:"media_id"`
	CreatedAt         string     `db:"created_at"`
	UpdatedAt         string     `db:"updated_at"`
	AuthorId          int64      `db:"author_id"`
	Author                       `db:"author"`
}

type Author struct {
	Id                int64      `db:"id"`
	Email             string     `db:"email"`
	FirstName         string     `db:"first_name"`
	LastName          string     `db:"last_name"`
	CreatedAt         string     `db:"created_at"`
	UpdatedAt         string     `db:"updated_at"`
}

func GetArticles(media_id int, limit int, offset int) ([]ArticleWithAuthor, error) {
	db := DBConnect()
	defer db.Close()

	sql := `
        select 
			article.*,
			author.id AS "author.id",
			author.email AS "author.email",
			author.first_name AS "author.first_name",
			author.last_name AS "author.last_name",
			author.created_at AS "author.created_at",
			author.updated_at AS "author.updated_at"
        from articles as article
        join authors as author on article.author_id = author.id 
        where article.media_id=?
        order by article.created_at   
        `
    var articles []ArticleWithAuthor
	err := db.Select(&articles, sql, media_id)
    checkErr(err, "fetch Articles failed")

	return articles, err
}

func GetArticle(media_id int, article_id int) (ArticleWithAuthor, error)  {
	db := DBConnect()
	defer db.Close()

	sql := `
        select 
			article.*,
			author.id AS "author.id",
			author.email AS "author.email",
			author.first_name AS "author.first_name",
			author.last_name AS "author.last_name"
        from articles as article
        join authors as author on article.author_id = author.id 
        where article.media_id=?
		and article.id=?
		limit 1
        `

    var article ArticleWithAuthor
	err := db.Get(&article, sql, media_id, article_id)
    checkErr(err, "fetch Articles failed")

	return article, err
}

