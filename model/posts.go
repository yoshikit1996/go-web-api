package model

import "time"

type Post struct {
	Name    string    `json:"name"`
	Comment string    `json:"comment"`
	Due     time.Time `json:"time"`
}

type Posts []Post

var dbWrapper = NewDbWrapper()

func NewPost(post Post) Post {
	dbWrapper.insertPost(post)
	return post
}

func GetPosts() Posts {
	return dbWrapper.select_all()
}
