package lib

import (
	"time"

	"github.com/tirthamoulib/event_loop/src/core"
)

// SetTimeout - sets a timeout for the given miliseconds
func SetTimeout(t int, cb func()) {
	go func() {
		time.Sleep(time.Duration(t) * time.Millisecond)
		CbChan <- core.CbFuncAndData{
			Function: cb,
		}
	}()
}
