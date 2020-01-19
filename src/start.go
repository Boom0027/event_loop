package src

import (
	"fmt"

	"github.com/tirthamoulib/event_loop/src/lib"
)

func createTimeLoop(num int) {
	fmt.Println(num)
	if num == 0 {
		return
	}
	lib.SetTimeout(1000, func() {
		createTimeLoop(num - 1)
	})
}

// Start - start of the program
func Start() {
	fmt.Println("Hello")
	createTimeLoop(10)
	fmt.Println("World")
}
