package main

import "fmt"

/**
* Go by Example: Variables
*/
func main() {

	// 声明字符（类型推断）
	var a = "initial"
	fmt.Println(a)

	// 声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 类型推断
	var d = true
	fmt.Println(d)

	// 整型初始化默认值为 0
	var e int
	fmt.Println(e)

	// 初始化 + 赋值缩写
	f := "apple"
	fmt.Println(f)


}