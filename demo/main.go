package main

import "fmt"

func main() {

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for i1 := range ch {
		fmt.Println(i1)
	}

}
