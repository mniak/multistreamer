package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/on_publish", onpublish)
	log.Fatal(http.ListenAndServe(":8080", router))
}
