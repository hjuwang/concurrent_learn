package main

import (
	"fmt"
	"time"
)

/**
复用定时器
*/

func main() {

	intChan := make(chan int, 1)

	go func() {

		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			intChan <- 1
		}
		close(intChan)
	}()

	timeout := time.Millisecond * 500
	var timer *time.Timer

	for {
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout)
		}

		select {
		case e, ok := <-intChan:
			if !ok {
				fmt.Println("End")
				return
			}
			fmt.Println("Receive", e)
		case <-timer.C:
			fmt.Println("Timeout")

		}
	}
}
