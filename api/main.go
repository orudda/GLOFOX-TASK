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
	bookingService := services.DBBookingService()
	classController := controllers.NewClassController(classService)
	bookingController := controllers.NewBookingController(bookingService)

	r.Use(middleware.Logger)
	r.Route("/classes", func(r chi.Router) {
		r.Get("/", classController.GetClasses)
		r.Post("/", classController.CreateClass)
	})

	r.Route("/bookings", func(r chi.Router) {
		r.Post("/", bookingController.CreateBooking)
	})

	serverAddr := "127.0.0.1:8081"
	fmt.Println("Servidor iniciado em:", serverAddr)

	http.ListenAndServe(serverAddr, r)
}
