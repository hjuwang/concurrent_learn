package main

import (
	"errors"
	"os"
	"sync"
)

type DataFile interface {
	Read() (rsn int64, d Data, err error)
	Write(d Data) (wsn int64, err error)
	RSN() int64
	// 获取最后写入的数据块的序列号

	WSN() int64
	DataLen() uint32
	Close() error
}

type Data []byte

// 用于表示数据文件的实现类型
type myDataFile struct {
	f       *os.File     // 文件
	fmutex  sync.RWMutex // 用于文件的读写锁
	woffset int64        // 写操作需要用到的偏移量
	roffset int64        // 读操作需要用到的偏移量
	wmutex  sync.Mutex   // 写操作需要用到的互斥锁
	rmutex  sync.Mutex   // 读操作需要用到的互斥锁
	dataLen uint32       // 数据块长度
}

func (m myDataFile) Read() (rsn int64, d Data, err error) {
	//TODO implement me
	panic("implement me")
}

func (m myDataFile) Write(d Data) (wsn int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (m myDataFile) RSN() int64 {
	//TODO implement me
	panic("implement me")
}

func (m myDataFile) WSN() int64 {
	//TODO implement me
	panic("implement me")
}

func (m myDataFile) DataLen() uint32 {
	//TODO implement me
	panic("implement me")
}

func (m myDataFile) Close() error {
	//TODO implement me
	panic("implement me")
}

// NewDataFile 新建一个数据文件的实例
func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}
	df := &myDataFile{f: f, dataLen: dataLen} //只对其中的两个字段进行初始化
	return df, nil
}
