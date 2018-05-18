package main

import (
	"log"
	"net/http"

	"github.com/yoshikit1996/go-webapi-bbs/router"
)

func main() {
	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
