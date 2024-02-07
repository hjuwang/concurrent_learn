package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	names := []string{"Eric", "Jim", "John", "Mark", "Tom"}

	for _, name := range names {

		go func(name string) {
			fmt.Printf("hello%s\n", name)
		}(name)
	}

	time.Sleep(time.Second)

	runtime.GC()
}
