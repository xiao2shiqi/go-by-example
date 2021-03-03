package main

import (
	"fmt"
	"sort"
)

// go 内置排序程序处理示例
func main() {

	// 数组字符串排序
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 数组整型排序
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	// 判断数组是否有序
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
