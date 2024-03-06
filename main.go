package main

import (
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome1", Historia: "historia 1"},
		{Id: 2, Nome: "nome2", Historia: "historia 2"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
