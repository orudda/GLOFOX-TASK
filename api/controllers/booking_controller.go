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
	ClassService   *services.ClassService
	BookingService *services.BookingService
}

func NewBookingController(classService *services.ClassService, bookingService *services.BookingService) *BookingController {
	return &BookingController{
		ClassService:   classService,
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
	classAvailable := c.isClassAvailable(newBooking.Date)
	if !classAvailable {
		utils.RespondWithError(w, http.StatusBadRequest, "Class not available on the requested date")
		return
	}

	c.BookingService.CreateBooking(newBooking)
	utils.RespondWithJSON(w, http.StatusCreated, newBooking)
}

func (c *BookingController) isClassAvailable(date utils.CustomTime) bool {
	classes := c.ClassService.GetClasses()
	for _, class := range classes {
		if date.Equal(class.StartDate.Time) || (date.After(class.StartDate.Time) && date.Before(class.EndDate.Time)) {
			// The class exists for the requested date
			return true
		}
	}
	return false
}
