package util

import (
	"time"
)

func Totimestamp(timestamp string) time.Time {
	const layout = "2006-01-02 15:04:05"
    t, _ := time.Parse(layout,timestamp)

	return t
}