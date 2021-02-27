package main

import "fmt"

// function are central in go
// func 是 go 的核心功能
func plus(a int, b int) int {
	
	return a + b
}

// a, b, c 多个同类型参数可以直接声明！
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	// 在 main 里面调用 plus 函数
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	// 在 main 里面调用 plusPlus 函数，并且传入 3个参数
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}