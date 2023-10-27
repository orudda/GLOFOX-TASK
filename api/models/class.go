// models/class.go
package models

import "time"

type Class struct {
	ID        int
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Capacity  int
}
