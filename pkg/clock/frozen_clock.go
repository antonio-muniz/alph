package clock

import "time"

type frozenClock struct {
	frozenAt time.Time
}

func NewFrozenClock(frozenAt time.Time) Clock {
	return frozenClock{frozenAt: frozenAt}
}

func (c frozenClock) Now() time.Time {
	return c.frozenAt.UTC()
}
