package main

import (
	b64 "encoding/base64"
	"fmt"
)

// golang 对于常用的 base64 的使用示例（encoding/base64 提供支持）
func main() {

	data := "abc123!?$*&()'-=@~"
	// 通过 data 生成出来的 sEnc
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 反向解析 sEnc 得到 data 数据
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
