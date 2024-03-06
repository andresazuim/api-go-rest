package routes

import (
	"api-go-rest/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasAsPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
