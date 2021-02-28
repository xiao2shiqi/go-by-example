// go 对指针使用的实例
package main

import "fmt"

func zeroval(ival int) {
	// 尝试改变变量 ival 的值
	ival = 0
}

// *int 意味着指针传递，使用引用传递
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)		// 如果是复制传递，那么函数的 i = 0 行为是不会改变 i 的值，所以 out: 1

	zeroptr(&i)
	fmt.Println("zeroptr:", i)		// *int 代表传递引用，那么函数的 i = 0 的行为是会改变 i 的值，所以 out: 0

	fmt.Println("pointer:", &i)		// 使用 &i 语法直接打印引用的内存地址
}