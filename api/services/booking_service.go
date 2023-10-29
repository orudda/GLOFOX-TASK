// services/booking_service.go
package services

import (
	"api/models"
	"errors"
)

type BookingService struct {
	Bookings []models.Booking
	Classes  *ClassService
	nextID   int
}

func DBBookingService(classService *ClassService) *BookingService {
	return &BookingService{
		Bookings: []models.Booking{},
		Classes:  classService,
		nextID:   1,
	}
}

func (s *BookingService) GetBookings() *[]models.Booking {
	return &s.Bookings
}

func (s *BookingService) CreateBooking(booking models.Booking) error {
	classAvailableId := s.Classes.IsClassAvailable(booking.Date)

	if classAvailableId == nil {
		return errors.New("Class not available on the requested date")
	}

	classAvailable, err := s.Classes.GetClassByID(*classAvailableId)
	if err != nil {
		return errors.New("Error when trying get class by Id")
	}

	err = s.Classes.DecrementClassCapacity(classAvailable.ID)
	if err != nil {
		return err
	}

	booking.ClassID = *classAvailableId
	booking.ID = s.nextID
	s.Bookings = append(s.Bookings, booking)
	s.nextID++
	return nil
}

func (s *BookingService) UpdateBooking(updatedBooking models.Booking) error {
	for i, booking := range s.Bookings {
		if booking.ID == updatedBooking.ID {
			s.Bookings[i] = updatedBooking
			return nil
		}
	}
	return errors.New("Booking not found")
}

func (s *BookingService) DeleteBooking(id int) error {
	for i, booking := range s.Bookings {
		if booking.ID == id {
			s.Bookings = append(s.Bookings[:i], s.Bookings[i+1:]...)
			err := s.Classes.IncrementClassCapacity(booking.ClassID)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("Booking not found")
}
