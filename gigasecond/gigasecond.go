package gigasecond

import (
	"math"
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	timestamp := t.Unix() + int64(math.Pow(10, 9))
	return time.Unix(timestamp, 0)
}
