package main

import (
	"fmt"
	"time"
)

/**
ticker 断续器
*/

func main() {

	intChann := make(chan int, 1)

	ticker := time.NewTicker(time.Second)

	go func() {
		for _ = range ticker.C { //每隔一秒向通道中发送一个数值

			select {
			case intChann <- 1:
			case intChann <- 2:
			case intChann <- 3:
			}
		}
		fmt.Println("End .[sender]")
	}()

	var sum int

	for e := range intChann {
		fmt.Printf("Receive %v\n", e)
		sum += e

		if sum > 10 {
			fmt.Printf("Got: %v\n", sum)
			break
		}
	}

	//ticker.Stop() //用完不要忘记停止

	fmt.Println("End .[receiver]")
}
