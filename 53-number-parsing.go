package main

import (
	"fmt"
	"strconv"
)

// golang 从字符中解析数字的示例程序，strconv 包的应用和使用
func main() {

	// 64 为解析的精度
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// ParseInt 表示结果要适合 64位
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// 识别十六进制格式
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
