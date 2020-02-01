package lib

import (
	"sync/atomic"

	"github.com/Boom0027/event_loop/src/core"
)

var cbChan = make(chan core.CbFuncAndData)
var endChan = make(chan bool)

// Count - Current counter of events
var count int64

// Init - Initializes the library
func Init() {
	defer core.EventLoop(cbChan, endChan, &count, decreaseCount)
}

// Increases the count of the number of additional routines running
func increaseCount() {
	atomic.AddInt64(&count, 1)
}

// Decreases the number of additional routines running. If the number is 0, send an end request
func decreaseCount() {
	newVal := atomic.AddInt64(&count, -1)
	if newVal == 0 {
		endChan <- true
	}
}
