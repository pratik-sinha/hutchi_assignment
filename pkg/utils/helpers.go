package utils

import "time"

func GetUTCTimeStamp() time.Time {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	return now
}
