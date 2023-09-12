package util

import "time"

func Totimestamp(timestamp string) time.Time {
    t, _ := time.Parse("2006-01-02T15:04:05Z",timestamp)
	return t
}