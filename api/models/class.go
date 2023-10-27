// models/class.go
package models

import "api/utils"

type Class struct {
	ID        int
	Name      string
	StartDate utils.CustomTime
	EndDate   utils.CustomTime
	Capacity  int
}
