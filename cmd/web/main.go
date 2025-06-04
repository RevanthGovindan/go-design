package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
}

func main() {
	var app = application{}
	server := &http.Server{
		Handler:           app.routes(),
		Addr:              port,
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}

	fmt.Println("starting web application on port", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
