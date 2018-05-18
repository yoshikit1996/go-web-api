package main

import (
	"github.com/yoshikit1996/go-webapi-bbs/router"
	"log"
	"net/http"
)

func main() {
	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
