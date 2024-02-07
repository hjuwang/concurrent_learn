package main

import (
	"fmt"
	"time"
)

func main() {

	names := []string{"Eric", "Jim", "John", "Mark", "Tom"}

	for _, name := range names {

		go func() {
			fmt.Printf("hello%s\n", name)
		}()
	}

	time.Sleep(time.Second)
}
