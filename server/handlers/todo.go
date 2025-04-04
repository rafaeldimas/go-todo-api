package handlers

import (
	"net/http"

	"github.com/rafaeldimas/go-todo-api/domain/todo"
)

type todoHandler struct {
	controller todo.TodoController
}

func NewTodoHandler(controller todo.TodoController) todoHandler {
	return todoHandler{
		controller: controller,
	}
}

func (handlers todoHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlers.controller.Create(w, r)
	case http.MethodGet:
		handlers.controller.List(w, r)
	case http.MethodPut:
		handlers.controller.Update(w, r)
	case http.MethodDelete:
		handlers.controller.Delete(w, r)
	default:
		http.NotFound(w, r)
	}
}
