package lib

import (
	"time"

	"github.com/tirthamoulib/event_loop/src/core"
)

// SetTimeout - sets a timeout for the given miliseconds
func SetTimeout(t int, cb func()) {
	// Step 1: Start function in another routine
	increaseCount()

	go func() {
		// Step A: Sleep for the duration
		time.Sleep(time.Duration(t) * time.Millisecond)

		// Step B: Send the callback back to the channel
		cbChan <- core.CbFuncAndData{
			Function: cb,
		}
	}()
}
