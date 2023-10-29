// services/booking_service.go
package services

import (
	"api/models"
	"errors"
)

type BookingService struct {
	Bookings []models.Booking
	Classes  *ClassService
}

func DBBookingService(classService *ClassService) *BookingService {
	return &BookingService{
		Bookings: []models.Booking{},
		Classes:  classService,
	}
}

func (s *BookingService) GetBookings() *[]models.Booking {
	return &s.Bookings
}

func (s *BookingService) CreateBooking(booking models.Booking) error {
	classAvailable := s.Classes.IsClassAvailable(booking.Date)

	if !classAvailable {
		return errors.New("Class not available on the requested date")
	}
	s.Bookings = append(s.Bookings, booking)
	return nil
}
