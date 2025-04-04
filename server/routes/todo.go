package routes

import (
	"net/http"

	"github.com/rafaeldimas/go-todo-api/domain/todo"

	infra "github.com/rafaeldimas/go-todo-api/infra/repositories/todo"
)

func RegisterTodoRoute(router *http.ServeMux) {
	repository := infra.NewMysqlRepository()
	service := todo.NewService(repository)
	controller := todo.NewController(service)

	router.HandleFunc("POST /todos", controller.Create)
	router.HandleFunc("GET /todos", controller.List)
	router.HandleFunc("PUT /todos/{id}", controller.Update)
	router.HandleFunc("DELETE /todos/{id}", controller.Delete)
}
