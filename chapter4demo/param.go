package chapter4demo

import (
	"example.com/m/lib"
	"time"
)

var timeoutNS time.Duration //响应超时时间，单位：ns
var Ips uint32              //每秒载荷量
var duration time.Duration  //负载持续时间

var resultCh chan *lib.CallResult
