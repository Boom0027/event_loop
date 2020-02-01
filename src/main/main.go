package main

import (
	"github.com/tirthamouli/event_loop/src"
	"github.com/tirthamouli/event_loop/src/lib"
)

func main() {
	// Step 1: Defer listening to the event loop
	defer lib.Init()

	// Step 2: Start the src
	src.Start()
}
