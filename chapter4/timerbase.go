package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("Present time: %v\n", time.Now())
	expirationTime := <-timer.C
	fmt.Printf("Expiration time: %v\n", expirationTime)

	fmt.Printf("Timer stopped?: %v\n", timer.Stop()) //false 说明定时器已经停止
}
