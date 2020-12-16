package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/ngoprek-kubernetes/kudemo/todo/api/db"
	"github.com/ngoprek-kubernetes/kudemo/todo/api/handler"
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

	mux := handler.InitRoutes(postgres)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
