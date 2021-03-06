package main

import (
	"fmt"
	"time"
)

// 获取 UNIX 开始以来的毫秒数
func main() {

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()

	// 打印当前时间
	fmt.Println(now)
	millis := nanos / 1000000

	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
