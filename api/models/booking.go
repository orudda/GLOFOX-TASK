// models/booking.go
package models

import "time"

type Booking struct {
	MemberName string
	Date       time.Time
}
