// models/booking.go
package models

import "api/utils"

type Booking struct {
	ID         int
	MemberName string
	Date       utils.CustomTime
}
