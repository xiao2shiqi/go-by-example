package main

import (
	"crypto/sha1"
	"fmt"
)

// golang 对于 sha1 散列计算示例("crypto/sha1" 包提供支持)
func main() {

	s := "sha1 this string"
	// 生成哈希的模式
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
