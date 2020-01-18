package lib

import "github.com/tirthamoulib/event_loop/src/core"

// CbChan - Global callback channel that is being used
var CbChan = make(chan core.CbFuncAndData)
