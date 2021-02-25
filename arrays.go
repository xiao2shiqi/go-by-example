package main

import "fmt"

// go by example array
func main() {

	// 声明 长度5 空数组
	var a [5]int
	fmt.Println("emp:", a)

	// 数组 a[4] 设置值为 100
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])		// 获取指定下标的数组

	fmt.Println("len:", len(a))		// 打印数组长度

	b := [5]int{1, 2, 3, 4, 5}		// 声明指定值的数组
	fmt.Println("dcl:", b)

	// 初始化数组和嵌套循环
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}