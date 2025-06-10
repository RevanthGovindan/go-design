package main

import (
	"flag"
	"fmt"
	"go-breeders/adapters"
	"go-breeders/configuration"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *configuration.Application
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {
	var app = application{templateMap: make(map[string]*template.Template)}

	flag.BoolVar(&app.config.useCache, "cache", false, "use template cache")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "dsn")
	flag.Parse()

	// get database

	db, err := initMysql(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}

	// jsonBackend := &JSONBackend{}
	// jsonAdapter := &RemoteService{
	// 	Remote: jsonBackend,
	// }
	// adapter pattern
	xmlBackend := &adapters.XmlBackend{}
	xmlAdapter := &adapters.RemoteService{
		Remote: xmlBackend,
	}

	app.App = configuration.New(db, xmlAdapter)

	server := &http.Server{
		Handler:           app.routes(),
		Addr:              port,
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}

	fmt.Println("starting web application on port", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
