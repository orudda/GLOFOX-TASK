// tests/booking_test.go
package tests

import (
	"api/models"
	"api/services"
	"api/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var date, _ = time.Parse("2006-01-02", "2023-12-01")
var booking = models.Booking{
	ID:         1,
	ClassID:    0,
	MemberName: "Test Member",
	Date:       utils.CustomTime{Time: date},
}
var class = models.Class{
	ID:        0,
	Name:      "Test Name",
	StartDate: utils.CustomTime{Time: date},
	EndDate:   utils.CustomTime{Time: date},
	Capacity:  10,
}

func TestBooking(t *testing.T) {
	booking := models.Booking{
		MemberName: "Test Member",
		Date:       utils.CustomTime{Time: time.Now()},
	}

	assert.Equal(t, "Test Member", booking.MemberName)
	assert.NotNil(t, booking.Date)
}

func TestCreateBooking(t *testing.T) {
	classService := services.DBClassService()
	bookingService := services.DBBookingService(classService)

	_ = classService.CreateClass(class)
	_ = bookingService.CreateBooking(booking)
	assert.Len(t, bookingService.Bookings, 1)
}

func TestUpdateBooking(t *testing.T) {
	classService := services.DBClassService()
	bookingService := services.DBBookingService(classService)
	_ = classService.CreateClass(class)
	_ = bookingService.CreateBooking(booking)

	booking.MemberName = "Updated Test Mamber"

	_ = bookingService.UpdateBooking(booking)

	actualBooking := bookingService.GetBookings()

	// Ensure that no bookings were created
	assert.Equal(t, "Updated Test Mamber", (*actualBooking)[0].MemberName)
}

func TestDeleteBooking(t *testing.T) {
	classService := services.DBClassService()
	bookingService := services.DBBookingService(classService)

	_ = classService.CreateClass(class)
	_ = bookingService.CreateBooking(booking)
	_ = bookingService.DeleteBooking(booking.ID)
	assert.Len(t, bookingService.Bookings, 0)
}
