package util

import "time"

func DurationInMicroSeconds(start, end time.Time) int64 {
	duration := end.Sub(start)
	return duration.Microseconds()
}
