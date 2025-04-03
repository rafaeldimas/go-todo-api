package handlers

import (
	"net/http"

	"github.com/rafaeldimas/go-todo-api/todo"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	controller := todo.NewController()

	switch r.Method {
	case http.MethodPost:
		controller.Create(w, r)
	case http.MethodGet:
		controller.List(w, r)
	case http.MethodPut:
		controller.Update(w, r)
	case http.MethodDelete:
		controller.Delete(w, r)
	}
}
