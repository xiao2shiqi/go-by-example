package main

import (
	"fmt"
	"sort"
)

// 自定义数组排序示例程序
type byLength []string

func (s byLength) Len() int {
	return len(s)
}

// 交换元素
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	// return len(s[i]) < len{s[j]}
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
