package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rafaeldimas/go-todo-api/server/routes"
)

func Init(port int64) {
	log.Printf("Servidor inicializado na port %d", port)

	routes.RegisterTodoRoute()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Error ao inicializar o servidor: %v", err)
	}
}
