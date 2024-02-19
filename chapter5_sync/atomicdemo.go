package main

import (
	"fmt"
	"sync/atomic"
)

/*
*
原子操作
*/
func main() {

	//加减操作
	var i32 int32
	atomic.AddInt32(&i32, 33)
	fmt.Println(i32)
	atomic.AddInt32(&i32, -20)
	fmt.Println(i32)

	//比较并交换（比较结果得到肯定的值后再交换）,compare and swap
	swapped := atomic.CompareAndSwapInt32(&i32, 13, 54)
	fmt.Println(i32, swapped)

	//原子的读
	val := atomic.LoadInt32(&i32)
	fmt.Println(val, i32)

	//原子的写
	atomic.StoreInt32(&i32, 80)
	fmt.Println(i32)

	//原子交换，并返回旧值

	old := atomic.SwapInt32(&i32, 34)
	fmt.Println(old, i32)

}

// 载入示例（读取操作,load）
var value int32

func addValue(delta int32) {

	for {

		v := atomic.LoadInt32(&value)
		if atomic.CompareAndSwapInt32(&value, v, v+delta) {

			break //如果添加成功就跳出循环
		}
	}

}

// DefineAtomicValue 自定义原子类型的值
func DefineAtomicValue() {

}
