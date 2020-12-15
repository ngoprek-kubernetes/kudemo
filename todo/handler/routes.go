package handler

import (
	"net/http"

	"github.com/ngoprek-kubernetes/kudemo/todo/db"
)

func InitRoutes() *http.ServeMux {
	todoHandler := &todoHandler{
		static: &db.Static{},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/static", todoHandler.GetStatic)

	return mux
}
