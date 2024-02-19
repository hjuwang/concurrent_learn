package v1

import (
	"errors"
	"io"
	"os"
	"sync"
)

// DataFile 用于表示数据文件的接口类型
type DataFile interface {
	// Read 读取一个数据块
	Read() (rsn int64, d Data, err error)
	// Write 写入一个数据块
	Write(d Data) (wsn int64, err error)
	// RSN 获取最后读取的数据块的序列号
	RSN() int64
	// 获取最后写入的数据块的序列号

	WSN() int64
	// DataLen 获取数据块的长度
	DataLen() uint32
	// Close 关闭数据文件
	Close() error
}

type Data []byte

// 用于表示数据文件的实现类型
type myDataFile struct {
	f       *os.File     // 文件
	fmutex  sync.RWMutex // 用于文件的读写锁,保护文件
	woffset int64        // 写操作需要用到的偏移量
	roffset int64        // 读操作需要用到的偏移量
	wmutex  sync.Mutex   // 写操作需要用到的互斥锁（保护偏移量）
	rmutex  sync.Mutex   // 读操作需要用到的互斥锁（保护偏移量）
	dataLen uint32       // 数据块长度
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新偏移量
	var offset int64 //offset 用于记录当前偏移量
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	// 读取一个数据块
	rsn = offset / int64(df.dataLen) //rsn 表示最后读取数据块的序列号
	bytes := make([]byte, df.dataLen)

	for {
		df.fmutex.RLock()
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.fmutex.RUnlock()
				continue
			}
			df.fmutex.RUnlock()
			return
		}
		d = bytes
		df.fmutex.RUnlock()
		return
	}

}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	//读取并更新偏移量
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()

	//写入一个数据块
	wsn = offset / int64(df.dataLen)

	var bytes []byte

	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}

	//加写锁
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	return
}

func (df *myDataFile) RSN() int64 {
	//TODO implement me
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) WSN() int64 {
	//TODO implement me
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	//TODO implement me

	return df.dataLen
}

func (df *myDataFile) Close() error {
	//TODO implement me
	return df.f.Close()
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
	df := &myDataFile{f: f, dataLen: dataLen}
	return df, nil
}
