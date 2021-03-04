package main

import (
	"fmt"
	"strings"
)
/*
 * go 语言操作数据集合的代码示例
 * 很多方法传入函数操作，感觉像函数式编程
 */ 
 // 查找数据下标
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 元素包含
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))	
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))		// index : 2
	fmt.Println(Include(strs, "grape"))		// false
	
	fmt.Println(Any(strs, func(v string) bool {		// true
		// 查找 p 开头的字符
		return strings.HasPrefix(v, "p")
	}))
	
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	// 过滤所有包含 v 的字符
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	// 传入数组和处理函数，对字符进行处理
	fmt.Println(Map(strs, strings.ToUpper))

}