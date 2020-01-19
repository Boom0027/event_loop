package main

import (
	"github.com/tirthamoulib/event_loop/src"
	"github.com/tirthamoulib/event_loop/src/lib"
)

func main() {
	// Step 1: Defer listening to the event loop
	defer lib.Init()

	// Step 2: Start the src
	src.Start()
}
