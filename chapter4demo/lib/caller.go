package lib

import "time"

// Caller 表示调用器接口
type Caller interface {
	BuildReq() RawReq
	Call(req []byte, timeoutNS time.Duration) ([]byte, error)
	CheckResp(rawReq RawReq, resp RawResp) *CallResult
}
