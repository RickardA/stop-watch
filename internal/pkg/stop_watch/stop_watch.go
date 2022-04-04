package stop_watch

import (
	"fmt"
	"time"
)

type StopWatch struct {
	timer   *time.Ticker
	counter int
	C       chan string
}

func NewStopWatch() *StopWatch {
	c := make(chan string, 1)

	return &StopWatch{
		counter: 0,
		C:       c,
	}
}

func (sw *StopWatch) Start() {
	sw.timer = time.NewTicker(time.Millisecond)

	go func() {
		prev := ""
		for range sw.timer.C {
			sw.counter++

			curr := sw.msToTime(sw.counter)

			if prev != curr {
				sw.C <- curr
			}

			prev = curr
		}
	}()
}

func (sw *StopWatch) Stop() {
	sw.timer.Stop()
	sw.counter = 0
}

func (sw *StopWatch) GetCurrentCount() int {
	return sw.counter
}

func (sw *StopWatch) msToTime(ms int) string {
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
