// main.go
package main

import (
	"fmt"
	"net/http"

	"api/controllers"
	"api/services"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	classService := services.NewClassService()
	classController := controllers.NewClassController(classService)

	r.Post("/classes", classController.CreateClass)

	serverAddr := "127.0.0.1:8000"
	fmt.Println("Servidor iniciado em:", serverAddr)

	http.ListenAndServe(serverAddr, r)
}
