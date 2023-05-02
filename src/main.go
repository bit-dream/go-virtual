package main

import (
	"fmt"
	"github.com/bit-dream/go-virtual/pkg/loaders"
	"github.com/bit-dream/go-virtual/pkg/virtualizer"
	"time"
)

func main() {

	ecus := loaders.FindAndLoadEcus()
	stop, err := virtualizer.StartAllVirtualEcus(ecus)
	if err != nil {
		fmt.Println("Application terminated prematurely due to an error: ", err)
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
