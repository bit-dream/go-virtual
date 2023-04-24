package virtualizer

import (
	"time"
)

func PeriodicTimer(cycleTime int, stop chan bool, callback func()) {

	ticker := time.NewTicker(time.Duration(cycleTime) * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			callback()
		case <-stop:
			return
		}
	}
}

func Trigger() {

}
