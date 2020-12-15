package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ngoprek-kubernetes/kudemo/todo/handler"
)

func main() {
	mux := handler.InitRoutes()
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
