package main

import "fmt"

func main() {

	dataChan := make(chan int, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() { //用于演示接受操作
		<-syncChan1
		for {
			if elem, ok := <-dataChan; ok {
				fmt.Printf("Received: %d\n", elem)
			} else {
				break
			}
		}
		fmt.Println("Done receiving.")
		syncChan2 <- struct{}{}
	}()

	go func() { //用于演示发送擦欧总

		for i := 0; i < 5; i++ {
			dataChan <- i
			fmt.Println("Sent: ", i)
		}
		close(dataChan)
		syncChan1 <- struct{}{}
		fmt.Println("Done sending.")
		syncChan2 <- struct{}{}

	}()

	<-syncChan2
	<-syncChan2
}
