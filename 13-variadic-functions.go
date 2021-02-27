package main

import "fmt"

// go 可变参数示例
func sum(numbs ...int) {
	fmt.Println(numbs, " ")
	total := 0
	// 通过案例可以看到，可变参数会被作为数组类型来解析
	for _, num := range numbs {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1)
	sum(1, 2)
	sum(1, 2, 3)

	numbs := []int{1,2,3,4}
	sum(numbs...)
}