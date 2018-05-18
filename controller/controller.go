package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/yoshikit1996/go-webapi-bbs/model"
)

func PostIndex(w http.ResponseWriter, r *http.Request) {
	posts := model.GetPosts()
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}
}

func NewPost(w http.ResponseWriter, r *http.Request) {

	var post model.Post

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &post); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := model.NewPost(post)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
