package main

import "fmt"

// range iterates over elements in a variety of data structures
// range 用于遍历各种数据结构，让遍历代码更加简短 ?
func main() {

	// 创建一个数组
	nums := []int{2, 3, 4}
	sum := 0
	// range 获取 2个返回值：index, element
	// 但是如果 index 没有使用，编译器会报错： index declared but not used，故而使用 _, 替代
	// for index, num := range nums {
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 通过 range 获取 index
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range 如何遍历 map ?
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range 遍历 map 只获取 key
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range 遍历字符串，会按照 index: unicode 来迭代
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	
}