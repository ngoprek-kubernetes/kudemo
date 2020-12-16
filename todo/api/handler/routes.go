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
	mux.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoHandler.getAllTodo(w, r)
		case http.MethodPost:
			todoHandler.insertTodo(w, r)
		case http.MethodDelete:
			todoHandler.deleteTodo(w, r)
		default:
			responseError(w, http.StatusNotFound, "")
		}
	})

	return mux
}
