// models/booking.go
package models

import "api/utils"

type Booking struct {
	ID         int
	ClassID    int
	MemberName string
	Date       utils.CustomTime
}
