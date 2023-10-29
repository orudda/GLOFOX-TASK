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
	bookingService := services.DBBookingService(classService)
	classController := controllers.NewClassController(classService)
	bookingController := controllers.NewBookingController(bookingService)

	r.Use(middleware.Logger)
	r.Route("/classes", func(r chi.Router) {
		r.Get("/", classController.GetClasses)
		r.Post("/", classController.CreateClass)

		r.Route("/{classID}", func(r chi.Router) {
			r.Get("/", classController.GetClassByID)
			r.Put("/", classController.UpdateClass)
			// r.Delete("/", classController.DeleteClass)
		})
	})

	r.Route("/bookings", func(r chi.Router) {
		r.Post("/", bookingController.CreateBooking)
		r.Get("/", bookingController.GetBookings)

		r.Route("/{bookingID}", func(r chi.Router) {
			// r.Put("/", bookingController.UpdateBooking)
			// r.Delete("/", bookingController.DeleteBooking)
		})
	})

	serverAddr := "127.0.0.1:8080"
	fmt.Println("Servidor iniciado em:", serverAddr)

	http.ListenAndServe(serverAddr, r)
}
