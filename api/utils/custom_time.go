// utils/custom_time.go
package utils

import (
	"time"
)

type CustomTime struct {
	time.Time
}

const customTimeLayout = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse(`"`+customTimeLayout+`"`, string(data))
	if err != nil {
		return err
	}
	ct.Time = parsedTime
	return nil
}
