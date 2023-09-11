package main

import (
	"bytes"
	"log"
	"mortage-calculator/api"
	"net/http"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func main() {

	router := api.Routes(*logger)

	srv := &http.Server{
		Addr:    ":8700",
		Handler: router,
	}
	srv.ListenAndServe()
}
