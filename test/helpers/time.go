package helpers

import "time"

var now = time.Date(2020, time.May, 24, 20, 35, 37, 0, time.UTC)

func Now() time.Time {
	return now
}
