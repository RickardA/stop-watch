package app

import "fmt"

func msToTime(ms int) string {
	minutes := 0
	seconds := 0
	centiSeconds := 0

	centiMs := ms % 1000
	if centiMs != 0 {
		centiSeconds = centiMs / 10
	}

	ms = ms - centiMs

	tempSec := ms / 1000

	seconds = tempSec % 60

	tempSec = tempSec - seconds

	minutes = tempSec / 60

	return fmt.Sprintf("%02d:%02d:%02d", minutes, seconds, centiSeconds)
}
