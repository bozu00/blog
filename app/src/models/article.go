package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id                int64  `db:"id"`
	Title             string `db:"title"`
	Content           string `db:"content"`
	AuthorId          int64  `db:"author_id"`
	AuthorName        string `db:"author_name"`
	ArticleCategoryId int64  `db:"article_category_id"`
	MediaId           int64  `db:"media_id"`
	CreatedAt         string `db:"created_at"`
	UpdatedAt         string `db:"updated_at"`
}

type ArticleWithAuthor struct {
	Id                int64  `db:"id"`
	Title             string `db:"title"`
	Content           string `db:"content"`
	ArticleCategoryId int64  `db:"article_category_id"`
	MediaId           int64  `db:"media_id"`
	CreatedAt         string `db:"created_at"`
	UpdatedAt         string `db:"updated_at"`
	AuthorId          int64  `db:"author_id"`
	Author            `db:"author"`
}

type Author struct {
	Id        int64  `db:"id"`
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
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

func GetArticle(media_id int, article_id int) (ArticleWithAuthor, error) {
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
