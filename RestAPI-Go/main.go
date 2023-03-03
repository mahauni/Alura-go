package main

import (
	"fmt"

	"github.com/mahauni/Alura-go/RestAPI-Go/models"
	"github.com/mahauni/Alura-go/RestAPI-Go/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
		{Nome: "Nome 3", Historia: "Historia 3"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
