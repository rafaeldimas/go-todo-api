package main

import (
	"log"

	"github.com/rafaeldimas/go-todo-api/server"
)

func main() {
	log.Println("Aplicação iniciada")
	server.Init(8080)
	log.Println("Aplicação finalizada")
}
