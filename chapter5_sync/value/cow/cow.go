package cow

import (
	"errors"
	"fmt"
	"sync/atomic"
)

type ConcurrentArray interface {

	// Set 设定值
	Set(index uint32, elem int) (err error)
	// Get 获取值
	Get(index uint32) (elem int, err error)
	// Len 获取长度
	Len() uint32
}

type concurrentArray struct {
	length uint32
	val    atomic.Value
}

func NewConcurrentArray(length uint32) ConcurrentArray {

	array := concurrentArray{}
	array.length = length
	array.val.Store(make([]int, length))

	//注意这里是返回指针类型
	return &array
}

func (array *concurrentArray) Set(index uint32, elem int) (err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}

	if err = array.checkValue(); err != nil {
		return
	}

	//先用新的new_array 存储旧数据，再将更新过的数组存储进 array 中
	newArray := make([]int, array.length)

	copy(newArray, array.val.Load().([]int))

	//newArray[index] = elem+ newArray[index]

	i := int32(newArray[index])

	atomic.AddInt32(&i, int32(elem))
	array.val.Store(newArray)

	return
}

func (array *concurrentArray) Get(index uint32) (elem int, err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}

	if err = array.checkValue(); err != nil {
		return
	}

	elem = array.val.Load().([]int)[index]
	return
}

func (array *concurrentArray) Len() uint32 {

	return array.length
}

// 添加检查索引的方法
// checkIndex 用于检查索引的有效性。
func (array *concurrentArray) checkIndex(index uint32) error {
	if index >= array.length {
		return fmt.Errorf("Index out of range [0, %d)!", array.length)
	}
	return nil
}

// checkValue 用于检查原子值中是否已存有值。
func (array *concurrentArray) checkValue() error {
	v := array.val.Load()
	if v == nil {
		return errors.New("Invalid int array!")
	}
	return nil
}
