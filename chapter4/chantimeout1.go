package main

import (
	"fmt"
	"time"
)

func main() {

	chanInt := make(chan int, 1)

	go func() {
		time.Sleep(time.Second)
		chanInt <- 1
	}()

	select {
	case v := <-chanInt:
		fmt.Printf("Receive %v\n", v)
	case <-time.NewTimer(time.Millisecond * 1500).C: //设置一个1500毫秒的定时器
		fmt.Printf("Timeout\n")
	}

	time.NewTimer(time.Second * 2).Stop()
}
