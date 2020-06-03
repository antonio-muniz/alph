package clock

import "time"

type workingClock struct{}

func NewWorkingClock() Clock {
	return workingClock{}
}

func (c workingClock) Now() time.Time {
	return time.Now().UTC()
}
