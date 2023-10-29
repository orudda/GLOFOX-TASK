// controllers/booking_controller.go
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/models"
	"api/services"
	"api/utils"

	"github.com/go-chi/chi"
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

	err := c.BookingService.CreateBooking(newBooking)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, newBooking)
}

func (c *BookingController) GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings := c.BookingService.GetBookings()
	utils.RespondWithJSON(w, http.StatusOK, bookings)
}

func (c *BookingController) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	bookingID, _ := strconv.Atoi(chi.URLParam(r, "bookingID"))
	var updatedBooking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&updatedBooking); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	updatedBooking.ID = bookingID
	if err := c.BookingService.UpdateBooking(updatedBooking); err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, updatedBooking)
}

func (c *BookingController) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	bookingID, _ := strconv.Atoi(chi.URLParam(r, "bookingID"))
	if err := c.BookingService.DeleteBooking(bookingID); err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}
