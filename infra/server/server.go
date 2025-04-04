package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rafaeldimas/go-todo-api/infra/server/routes"
)

func Init(port int64) {
	router := http.NewServeMux()

	routes.RegisterTodoRoute(router)

	log.Printf("Servidor inicializado na port %d", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Fatalf("Error ao inicializar o servidor: %v", err)
	}
}
