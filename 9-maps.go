package main

import "fmt"

func main() {
	// 初始化 map 格式：map[key-type]val-type
	m := make(map[string] int)

	// 设置 kv 到 map
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)		// out： map 输出格式也很好 map[k1:7 k2:13]

	// 获取 map 指定 key: map["key"]
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	// 通过内置函数来删除 map 指定 key ?
	delete(m, "k2")
	fmt.Println("map:", m)		//  map[k1: 7]

	_, prs := m["k2"]			// out false, 使用 false 来表示不存在的值
	fmt.Println("prs:", prs)

	// 一行代码声明并且初始化 map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}