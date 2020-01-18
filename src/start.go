package src

import (
	"fmt"

	"github.com/tirthamoulib/event_loop/src/lib"
)

// Start - start of the program
func Start() {
	fmt.Println("Hello")
	lib.SetTimeout(1000, func() {
		fmt.Println("Anything")
		lib.SetTimeout(1000, func() {
			fmt.Println("Something else")
		})
	})
	fmt.Println("World")
}
