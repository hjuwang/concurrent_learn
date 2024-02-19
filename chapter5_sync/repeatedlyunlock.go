package main

import (
	"fmt"
	"sync"
)

/*
解锁未锁定的mutex 会引发panic，且不能被recover
*/

func main() {

	defer func() {

		fmt.Println("Try recover the panic")
		if p := recover(); p != nil {
			fmt.Printf("Recovered the panic (%#v).\n", p)
		}
	}()

	var mutex sync.Mutex
	fmt.Println("Lock the lock")

	mutex.Lock()
	fmt.Println("The lock is locked")
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	fmt.Println("The lock is unlocked")
	fmt.Println("Unlock the lock again")
	mutex.Unlock()
}
