package timex

import "time"

func TimestampNow() int64 {
	return time.Now().Unix()
}
