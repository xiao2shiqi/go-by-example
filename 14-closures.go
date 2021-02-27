package main

import "fmt"

// go 匿名函数和闭包的使用展示
func intSeq() func() int {
	i := 0
	// 这里返回的是一个匿名函数
	return func() int {
		i++
		return i
	}
}

func main() {
	// 这里 nextInt 接收的是一个函数
	nextInt := intSeq()
	// nextInt 可以作为函数被使用
	fmt.Println(nextInt())			// this is out: 1
	// 函数反复调用，但内部的变量不会被改变
	fmt.Println(nextInt())			// this is out: 2
	fmt.Println(nextInt())			// this is out: 3
	// 重新获取函数示例后，内部变量被改变
	newInts := intSeq()
	fmt.Println(newInts())			// this is out: 1
}