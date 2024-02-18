package lib

import "time"

// CallResult 用于表示调用结果的结构
type CallResult struct {
	ID     int64         // ID
	Req    RawReq        // 原生请求
	Resp   RawResp       // 原生响应
	Code   RetCode       // 响应代码
	Msg    string        // 结果成因的简述
	Elapse time.Duration // 耗时
}

// RawReq 用于表示原生请求的结构
type RawReq struct {
	ID  int64
	Req []byte
}

// RawResp 用于表示原生响应的结构
type RawResp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}

type RetCode int
