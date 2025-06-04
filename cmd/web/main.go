package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
}

type appConfig struct {
	useCache bool
}

func main() {
	var app = application{templateMap: make(map[string]*template.Template)}

	flag.BoolVar(&app.config.useCache, "cache", false, "use template cache")
	flag.Parse()

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
