package router

import (
	"go-webapi-bbs/handlers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"PostIndex",
		"GET",
		"/posts/",
		handlers.PostIndex,
	},
	Route{
		"NewPost",
		"POST",
		"/posts/new",
		handlers.NewPost,
	},
}
