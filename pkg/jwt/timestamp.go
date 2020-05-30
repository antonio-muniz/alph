package jwt

import (
	"strconv"
	"time"
)

type Timestamp time.Time

func (t Timestamp) MarshalJSON() ([]byte, error) {
	secondsSinceEpoch := time.Time(t).Unix()
	serializedTimestamp := []byte(strconv.FormatInt(secondsSinceEpoch, 10))
	return serializedTimestamp, nil
}

func (t *Timestamp) UnmarshalJSON(serializedTimestamp []byte) (err error) {
	secondsSinceEpoch, err := strconv.ParseInt(string(serializedTimestamp), 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(secondsSinceEpoch, 0).UTC()
	return nil
}
