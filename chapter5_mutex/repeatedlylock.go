package main

import (
	"fmt"
	"sync"
	"time"
)

/**
锁定已经锁定的 mutex 会阻塞
解锁为未锁定的锁 会发生panic
*/

func main() {

	var mutex sync.Mutex
	fmt.Println("Lock the lock.(main)")

	mutex.Lock()
	fmt.Println("The lock is locked.(main)")

	for i := 0; i < 3; i++ {

		go func(i int) {
			fmt.Printf("Lock the lock.(g%d)\n", i)
			mutex.Lock() //锁定已经锁定的 mutex 会阻塞当前的 goroutine
			fmt.Printf("The lock is locked .(g%d)\n", i)

		}(i)

	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock.(main)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked.(main)")
	time.Sleep(time.Second)
}
