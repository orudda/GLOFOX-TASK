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
	return &BookingController{BookingService: bookingService}
}

func (c *BookingController) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var newBooking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&newBooking); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	c.BookingService.CreateBooking(newBooking)
	utils.RespondWithJSON(w, http.StatusCreated, newBooking)
}
