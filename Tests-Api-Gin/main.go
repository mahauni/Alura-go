package main

import (
	"github.com/mahauni/Alura-go/Tests-Api-Gin/database"
	"github.com/mahauni/Alura-go/Tests-Api-Gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
