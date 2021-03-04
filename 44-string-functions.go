package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

// 所有对 strings 函数处理的方法示例
func main() {

	p("Contains:	", s.Contains("test", "es"))	// 包含 es 字符示例
	p("Count:		", s.Count("test", "t"))		// 字符 Count 统计
	p("HasPrefix:	", s.HasPrefix("test", "te"))	// 字符前缀
	p("HasSuffix:	", s.HasSuffix("test", "st"))	// 字符后缀
	p("Index:		", s.Index("test", "e"))		// 查找 Index
	p("Join:		", s.Join([]string{"a", "b"}, "-"))	// Join 通过字符将数组连接
	p("Repeat:		", s.Repeat("a", 5))				// 重复字符
	p("Replace:   ", s.Replace("foo", "o", "0", -1))	// 指定下标替换字符
	p("Replace:   ", s.Replace("foo", "o", "0", 1))		// 指定下标替换字符
	p("Split:     ", s.Split("a-b-c-d-e", "-"))			// 字符分割
	p("ToLower:   ", s.ToLower("TEST"))					// 大小写转换
    p("ToUpper:   ", s.ToUpper("test"))
    p()

    p("Len: ", len("hello"))							// 字符长度
    p("Char:", "hello"[1])								// 字符集编码
}