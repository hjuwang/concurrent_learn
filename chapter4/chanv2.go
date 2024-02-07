package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

var mapChan1 = make(chan map[string]Counter, 1)

func main() {

	syncChan := make(chan struct{}, 1)
	go func() { //接收方
		for {
			if elem, ok := <-mapChan1; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}

		}

		fmt.Println("stopped receiving.")
		syncChan <- struct{}{}
	}()

	go func() { //发送且展示方

		countMap := map[string]Counter{
			"count": Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan1 <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)

		}

		close(mapChan1)
		syncChan <- struct{}{}

	}()

	<-syncChan
	<-syncChan
}
