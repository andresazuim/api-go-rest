package main

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {

	models.Personalidades = []models.Personalidade{}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
