package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ngoprek-kubernetes/kudemo/todo/db"
	"github.com/ngoprek-kubernetes/kudemo/todo/service"
)

type todoHandler struct {
	static *db.Static
}

func (h *todoHandler) GetStatic(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepo(r.Context(), h.static)

	todoList, err := service.GetAll(ctx)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseOK(w, todoList)
}

func responseOK(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func responseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}
