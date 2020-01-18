package core

import (
	"log"
	"runtime"
)

/*EventLoop - The event loop checks if there are any running go routines.
 *If there are any running go routines wait for it and push the callback to the call stack
 */
func EventLoop(cbChan <-chan CbFuncAndData) {
	// Step 1: Start the loop
	for {
		// Step 2: Check for any running go routines. If there are no routines, break
		if runtime.NumGoroutine() == 1 {
			break
		}

		// Step 3: Wait for data from the channel
		funcAndData := <-cbChan

		// Step 4: Assert function signature 1
		cb, ok := funcAndData.Function.(func())
		if ok {
			cb()
			continue
		}

		// Step 5: Assert function signature 2
		cb1, ok := funcAndData.Function.(func(...interface{}))
		if ok {
			data := funcAndData.Data
			cb1(data...)
		}

		// Step 6: Panin if function doesn't match any
		if !ok {
			log.Panicln("Callback should be a function")
		}

		// Step 7: Run the garbage collector
		runtime.GC()
	}
}
