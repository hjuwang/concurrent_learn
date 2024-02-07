package main

import (
	"fmt"
	"unsafe"
)

func main() {

	/*在 go 中 空结构体类型的变量是不占用内存的,g
	并且所有该类型的变量都拥有相同的内存地址（从这句话也能看出，空结构类型是不占用内存的                          ）
	*/

	var a struct{}
	var b struct{}

	fmt.Printf("%p %p\n", &a, &b)

	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
}
