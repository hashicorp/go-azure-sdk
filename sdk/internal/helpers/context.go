package helpers

import "time"

// SleepUntil sleeps until the specified interval or the context is cancelled/times out
func SleepUntil(nextInterval time.Duration, doneFunc <-chan struct{}) {
	select {
	case <-time.After(nextInterval):
		return
	case <-doneFunc:
		return
	}
}
