package time_day

import "time"

var now = time.Now

func IsWorkMorning() bool {
	t := now()
	return t.Weekday() >= time.Monday && t.Weekday() <= time.Friday && t.Hour() < 12
}
