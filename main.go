package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rafaeldimas/go-todo-api/infra/server"
)

func main() {
	log.Println("Aplicação iniciada")
	server.Init(8080)
}
