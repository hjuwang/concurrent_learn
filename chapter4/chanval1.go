package main

import (
	"fmt"
	"time"
)

// 注意 map 是引用类型
var mapChan = make(chan map[string]int, 1)

func main() {

	//同步子 routine
	syncChan := make(chan struct{}, 2)

	go func() { //用于显示接受操作
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			} else {
				break
			}

		}

		fmt.Println("Stopped Receiving. [receiver]")

		syncChan <- struct{}{}
	}()

	go func() { //用于演示发送操作
		countMap := make(map[string]int, 5)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}

		close(mapChan)
		syncChan <- struct{}{}

	}()
	<-syncChan
	<-syncChan
}
