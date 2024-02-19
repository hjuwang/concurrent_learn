package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var rwm sync.RWMutex

	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Try lock for reading... [%d]\n", i)
			rwm.RLock()
			fmt.Printf("Lock for reading... [%d]\n", i)
			time.Sleep(time.Second * 2)
			fmt.Printf("Try unlock for reading... [%d]\n", i)
			rwm.RUnlock()
			fmt.Printf("Unlock for reading... [%d]\n", i)
		}(i)

	}

	time.Sleep(time.Millisecond * 100)
	fmt.Println("Try lock for writing")
	rwm.Lock()
	fmt.Println("locked for writing")

}
