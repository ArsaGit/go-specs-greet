package main

import (
	"log"
	"net/http"

	"github.com/ArsaGit/go-specs-greet/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
	// TODO Add new endpoint(http.NewServeMux)
}
