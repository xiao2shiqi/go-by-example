package main

import (
	"fmt"
	"math/rand"
	"time"
)

// golang 对于随机数提供的支持和示例，随机数所在包：math/rand
func main() {

	// 使用 rand.Intn 随机数貌似是确定性的
	fmt.Print(rand.Intn(100), ",") // 使用 , 号结尾
	fmt.Print(rand.Intn(100))

	fmt.Println()
	// 生成随机浮点数
	fmt.Println(rand.Float64())
	// 格式化浮点数
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 真正意义的随机数....需要一个 seed 种子数据
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 固定种子数据，生成固定的随机数
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
