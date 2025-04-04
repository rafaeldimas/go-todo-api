package routes

import (
	"net/http"

	"github.com/rafaeldimas/go-todo-api/domain/todo"
	"github.com/rafaeldimas/go-todo-api/server/handlers"

	infra "github.com/rafaeldimas/go-todo-api/infra/repositories/todo"
)

func RegisterTodoRoute() {
	repository := infra.NewInMemoryRepository()
	service := todo.NewService(repository)
	controller := todo.NewController(service)

	http.HandleFunc("/todos", handlers.NewTodoHandler(controller).Handler)
}
