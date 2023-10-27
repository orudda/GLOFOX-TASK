// main.go
package main

import (
	"fmt"
	"net/http"

	"api/controllers"
	"api/services"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	classService := services.DBClassService()
	classController := controllers.NewClassController(classService)

	r.Use(middleware.Logger)
	r.Route("/classes", func(r chi.Router) {
		r.Get("/", classController.GetClasses)
		r.Post("/", classController.CreateClass)
	})

	serverAddr := "127.0.0.1:8081"
	fmt.Println("Servidor iniciado em:", serverAddr)

	http.ListenAndServe(serverAddr, r)
}
