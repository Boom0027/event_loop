package core

import (
	"log"
	"runtime"
)

/*EventLoop - The event loop checks if there are any running go routines.
 *If there are any running go routines wait for it and push the callback to the call stack
 */
func EventLoop(cbChan <-chan CbFuncAndData) {
	for true {
		if runtime.NumGoroutine() == 1 {
			break
		}
		funcAndData := <-cbChan
		cb, ok := funcAndData.Function.(func())
		if ok {
			cb()
		} else {
			cb1, ok := funcAndData.Function.(func(...interface{}))
			if ok {
				data := funcAndData.Data
				cb1(data...)
			}
			if !ok {
				log.Panicln("Callback should be a function")
			}
		}
		runtime.GC()
	}
}
