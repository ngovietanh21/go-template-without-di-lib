package timeutil

import (
	"promotion/pkg/failure"
	"time"
)

func GetCurrentTimeInLocation(location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Now(), failure.ErrorWithTrace(err)
	}
	return time.Now().In(loc), nil
}
