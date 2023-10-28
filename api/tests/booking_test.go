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

func TestBooking(t *testing.T) {
	booking := models.Booking{
		MemberName: "Test Member",
		Date:       utils.CustomTime{Time: time.Now()},
	}

	assert.Equal(t, "Test Member", booking.MemberName)
	assert.NotNil(t, booking.Date)
}

func TestCreateBooking(t *testing.T) {
	bookingService := services.DBBookingService()
	booking := models.Booking{
		MemberName: "Test Member",
		Date:       utils.CustomTime{Time: time.Now()},
	}

	bookingService.CreateBooking(booking)
	assert.Len(t, bookingService.Bookings, 1)
}
