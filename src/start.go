package src

import (
	"fmt"

	"github.com/tirthamoulib/event_loop/src/lib"
)

// Start - start of the program
func Start() {
	fmt.Println("Hello")
	for i := 0; i < 1000000; i++ {
		j := i
		lib.SetTimeout(0, func() {
			fmt.Println(j, "Sup")
		})
	}

	fmt.Println("World")
}
