// services/booking_service.go
package services

import (
	"api/models"
)

type BookingService struct {
	Bookings []models.Booking
}

func DBBookingService() *BookingService {
	return &BookingService{
		Bookings: []models.Booking{},
	}
}

func (s *BookingService) CreateBooking(booking models.Booking) {
	s.Bookings = append(s.Bookings, booking)
}
