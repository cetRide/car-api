package util

import (
	"time"
)

func GenerateTimeNow() time.Time {
	timeStamp, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		panic(err)
	}
	return timeStamp
}

func CheckTimeDifference(created_at time.Time) float64 {
	diff := GenerateTimeNow().Sub(created_at).Hours()
	return diff
}
