// services/booking_service.go
package services

import (
	"api/models"
	"api/utils"
	"errors"
)

type BookingService struct {
	Bookings     []models.Booking
	ClassService *ClassService
}

func DBBookingService() *BookingService {
	return &BookingService{
		Bookings: []models.Booking{},
	}
}

func (s *BookingService) GetBookings() *[]models.Booking {
	return &s.Bookings
}

func (s *BookingService) CreateBooking(booking models.Booking) error {
	classAvailable := s.IsClassAvailable(booking.Date)
	if !classAvailable {
		return errors.New("Class not available on the requested date")
	}
	s.Bookings = append(s.Bookings, booking)
	return nil
}

func (c *BookingService) IsClassAvailable(date utils.CustomTime) bool {
	classes := c.ClassService.GetClasses()
	for _, class := range classes {
		if date.Equal(class.StartDate.Time) || (date.After(class.StartDate.Time) && date.Before(class.EndDate.Time)) {
			// The class exists for the requested date
			return true
		}
	}
	return false
}
