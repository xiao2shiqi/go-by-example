package main

// 引入多个包
import (
    "fmt"
    "math"
)

// 使用 const 声明一个常量
const s string = "constant"

func main() {
	// 输出常量
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n

	// cannot assign to n (declared const) !!!
	// n = 50001
	// fmt.Println(n)

	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}