package main

// 引入多个包
import (
    "fmt"
    "math"
)

// 使用 const 声明一个常量
const s string = "constant"

func main() {
	fmt.Println(s)

	// const 可以出现在任何地方...
	const n = 500000000
	const d = 3e20 / n

	// 无法改变的常量
	// n = 50001  // cannot assign to n (declared const) !!! 

	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}