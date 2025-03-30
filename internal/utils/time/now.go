package timeutil

import "time"

func UTCNow() time.Time {
	return time.Now().UTC()
}
