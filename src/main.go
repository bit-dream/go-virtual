package main

import (
	"github.com/bit-dream/go-virtual/pkg/loaders"
	"github.com/bit-dream/go-virtual/pkg/virtualizer"
	"time"
)

func main() {

	vecus := loaders.FindAndLoadEcus()

	stop := make(chan bool)

	for _, vecu := range vecus {
		err := virtualizer.VirtualizeEcu(vecu, stop)
		if err != nil {
			stop <- true
			return
		}
	}

	// Create a timer that will send a signal to the stop channel after 20 seconds.
	timer := time.NewTimer(20 * time.Second)
	defer timer.Stop()

	// Wait for either the timer to expire or a stop signal to be received.
	select {
	case <-timer.C:
		// Timer expired, send stop signal to all goroutines.
		stop <- true
	case <-stop:
		// Stop signal received.
	}
}
