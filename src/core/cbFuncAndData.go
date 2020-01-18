package core

// CbFuncAndData - Defines the structure of the message through the channel
type CbFuncAndData struct {
	Function interface{}   // The call back function
	Data     []interface{} // The data that needs to be passed to the function
}
