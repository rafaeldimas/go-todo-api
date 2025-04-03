package routes

import (
	"net/http"

	"github.com/rafaeldimas/go-todo-api/server/handlers"
)

func RegisterTodoRoute() {
	http.HandleFunc("/todos", handlers.TodoHandler)
}
