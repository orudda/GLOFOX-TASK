// controllers/booking_controller.go
package controllers

import (
	"encoding/json"
	"net/http"

	"api/models"
	"api/services"
	"api/utils"
)

type BookingController struct {
	BookingService *services.BookingService
}

func NewBookingController(bookingService *services.BookingService) *BookingController {
	return &BookingController{
		BookingService: bookingService,
	}
}

func (c *BookingController) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var newBooking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&newBooking); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Check if a class is available on the requested date
	classAvailable := c.BookingService.IsClassAvailable(newBooking.Date)
	if !classAvailable {
		utils.RespondWithError(w, http.StatusBadRequest, "Class not available on the requested date")
		return
	}

	c.BookingService.CreateBooking(newBooking)
	utils.RespondWithJSON(w, http.StatusCreated, newBooking)
}

func (c *BookingController) GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings := c.BookingService.GetBookings()
	utils.RespondWithJSON(w, http.StatusOK, bookings)
}
