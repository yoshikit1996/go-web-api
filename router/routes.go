package router

import (
	"net/http"

	"github.com/yoshikit1996/go-webapi-bbs/controller"
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
		controller.PostIndex,
	},
	Route{
		"NewPost",
		"POST",
		"/posts/new",
		controller.NewPost,
	},
}
