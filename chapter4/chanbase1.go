package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {

	//用于发送者通知 接收者 开始运行
	syncChan1 := make(chan struct{}, 1)

	//用于阻塞主 routine 等待 子routine 运行完毕
	syncChan2 := make(chan struct{}, 2)

	go func() { // 用于演示接受操作
		<-syncChan1
		fmt.Println("接受一个同步信号，并且等待一秒钟。。。【接收者】")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("接收到消息:", elem)
			} else {
				break
			}

		}

		fmt.Println("停止接受，【接收者】")
		syncChan2 <- struct{}{}
	}()

	go func() { //用于演示发送操作
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("sent", elem, "[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("发送一个同步信号，【发送者】")
			}
		}
		fmt.Println("等待2秒钟。。。【发送者】")
		time.Sleep(time.Second * 2)
		close(strChan)
		syncChan2 <- struct{}{}

	}()

	<-syncChan2
	<-syncChan2

}
