package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ngoprek-kubernetes/kudemo/todo/db"
	"github.com/ngoprek-kubernetes/kudemo/todo/handler"
)

func main() {
	var postgres *db.Postgres
	var err error
	postgres, err = db.ConnectPostgres()
	if err != nil {
		panic(err)
	}
	if postgres == nil {
		panic("postgres is unreachable")
	}

	mux := handler.InitRoutes()
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
