package main

import "fmt"

// for 是 go 唯一的循环
func main() {

	i := 1
	// 最简单的条件循环
	for i <= 3 {
		i = i + 1
	}
	fmt.Println("i result:", i)		// i == 4

	// 初始化条件j的循环
	for j := 7; j <= 9; j++ {
		// j = 7, 8, 9
		fmt.Println(j)
	}

	// fmt.Println("j result:", j)	// undefined j	局部变量作用域问题

	// 无限循环，条件退出
	for {
		fmt.Println("loop")
		break
	}

	// continue 在 go 中的应用
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}