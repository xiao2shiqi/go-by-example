package main

import (
	"bytes"
	"fmt"
	"regexp"
)

// go 对于正则表达式的示例
func main() {

	// 正则最基本的使用，首个返回值是匹配结果 bool
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 导出正则进行匹配
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))

	// 找到正则匹配的字符 true
	fmt.Println(r.FindString("peach punch"))
	// 找到正则匹配字符下标 [0, 5]
	fmt.Println(r.FindStringIndex("peach punch"))
	// 子匹配项和下标
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
