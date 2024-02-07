package main

import (
	"fmt"
	"time"
)

var strChan3 = make(chan string, 3)

func main() {
	//通知chan
	notifyChan := make(chan struct{}, 1)
	waitChan := make(chan struct{}, 2)
	go receive(strChan3, notifyChan, waitChan)
	go send(strChan3, notifyChan, waitChan)

	<-waitChan
	<-waitChan

}

func receive(strChan <-chan string, noteCh <-chan struct{}, waitCh chan<- struct{}) {
	<-noteCh
	fmt.Println("接受一个同步信号，并且等待一秒钟。。。【接收者】")
	time.Sleep(time.Second)

	for elem := range strChan {
		fmt.Println("接收到消息:", elem)
	}

	fmt.Println("停止接受，【接收者】")

	waitCh <- struct{}{}
}

func send(strChan chan<- string, noteCh chan<- struct{}, waitCh chan<- struct{}) {

	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("sent", elem, "[sender]")
		if elem == "c" {
			noteCh <- struct{}{}
			fmt.Println("发送一个同步信号，【发送者】")
		}
	}
	fmt.Println("等待2秒钟。。。【发送者】")
	time.Sleep(time.Second * 2)
	close(strChan)
	waitCh <- struct{}{}

}
