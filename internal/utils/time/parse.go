package timeutil

import (
	"strings"
	"time"
)

func Parse(timeString string) (time.Time, error) {
	if strings.Contains(timeString, " ") {
		timeString = strings.Replace(timeString, " ", "+", 1)
	}
	return time.Parse(time.RFC3339, timeString)
}
