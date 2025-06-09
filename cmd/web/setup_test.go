package main

import (
	"go-breeders/configuration"
	"log"
	"testing"
)

var testApp application

func TestMain(m *testing.T) {
	dsn := "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s"
	db, err := initMysql(dsn)
	if err != nil {
		log.Panic(err)
	}
	testApp = application{
		App: configuration.New(db),
	}
}
