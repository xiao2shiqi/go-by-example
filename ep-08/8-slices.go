package main

import "fmt"

// slices 是 go 关键的数据类型，提供比 arrays 强大的接口
func main() {

	// 使用 make 创建长度的 3的空切片
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 像数组一样设置 slices
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)			// out: [a b c]
	fmt.Println("get:", s[2])		// out: [c]		... 下标从 0 开始
	fmt.Println("len:", len(s))		// out: 3

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)			// out: [a b c d e f]

	c := make([]string, len(s))		// create a len = 5 slices
	// 拷贝 slices
	copy(c, s)
	fmt.Println("cpy:", c)			// out: [a b c d e f]

	// 获取范围 range[2..5]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 获取范围 range[0..5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// 获取范围 range[2..5]
	l = s[2:]
	fmt.Println("fl3:", l)

	// 一行代码声明和初始化 slices  := []string{}
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// slices 多维实现
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}