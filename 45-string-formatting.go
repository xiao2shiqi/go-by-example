package main

import "fmt"

// go 为字符串格式化提供很好的支持，以下是示例
type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	// # 对象格式化输出，通常使用 Printf

	// 输出 value
	fmt.Printf("%v\n", p)
	// 输出 k:v
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 14)

	// ... 还有很多格式输出，这里就不深入研究了，有需要直接查字典就行

}
