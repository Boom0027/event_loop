package core

import (
	"log"
)

/*EventLoop - The event loop checks if there are any running go routines.
 *If there are any running go routines wait for it and push the callback to the call stack
 */
func EventLoop(cbChan chan CbFuncAndData, endChan <-chan bool, c *int64, dc func()) {
	go checkIfFinished(c, cbChan, endChan)

	// Step 1: Start the loop
	for funcAndData := range cbChan {
		// Step A: Assert function signature 1
		cb, ok := funcAndData.Function.(func())
		if ok {
			// Step i: Callback
			cb()

			// Step ii: Decrease count
			dc()
			continue
		}

		// Step B: Assert function signature 2
		cb1, ok := funcAndData.Function.(func(...interface{}))
		if ok {
			data := funcAndData.Data
			cb1(data...)
		}

		// Step C: Panic if function doesn't match any
		if !ok {
			log.Panicln("Callback should be a function")
		}

		// Step D: Decrease count
		dc()
	}
}

func checkIfFinished(c *int64, cbChan chan CbFuncAndData, endChan <-chan bool) {
	res := <-endChan
	if res {
		close(cbChan)
	}
}
