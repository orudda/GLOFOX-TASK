// models/booking.go
package models

import "api/utils"

type Booking struct {
	MemberName string
	Date       utils.CustomTime
}
