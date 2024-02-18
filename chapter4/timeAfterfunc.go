package main

import (
	"fmt"
	"time"
)

func main() {

	//一秒之后会执行函数中的代码
	_ = time.AfterFunc(time.Second, func() {
		fmt.Println("After 1 second")
	})

	time.Sleep(time.Second * 3)

}
