package model

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DbWrapper struct {
	db          *(sql.DB)
	time_format string
}

func NewDbWrapper() DbWrapper {
	db, err := sql.Open("sqlite3", "./model/comments.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "COMMENTS" ("NAME" VARCHAR(255), "COMMENT" VARCHAR(255), "TS" default CURRENT_TIMESTAMP)`,
	)
	if err != nil {
		panic(err)
	}

	return DbWrapper{db, "2006-01-02 15:04:05"}
}

func (dw DbWrapper) insertPost(post Post) {
	_, err := dw.db.Exec(
		`INSERT INTO COMMENTS (NAME, COMMENT, TS) VALUES (?, ?, ?)`,
		post.Name,
		post.Comment,
		time.Now().Format(dw.time_format),
	)
	if err != nil {
		panic(err)
	}
}

func (dw DbWrapper) select_all() Posts {
	var posts = Posts{}

	rows, err := dw.db.Query(
		`SELECT NAME, COMMENT, TS FROM COMMENTS`,
	)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var name, comment, ts string

		if err := rows.Scan(&name, &comment, &ts); err != nil {
			log.Fatal("rows.Scan()", err)
			return nil
		}

		_ts, _ := time.Parse(dw.time_format, ts)
		if err != nil {
			log.Fatal("time.Parse()", err)
		}

		post := Post{Name: name, Comment: comment, Due: _ts}
		posts = append(posts, post)
	}
	return posts
}
