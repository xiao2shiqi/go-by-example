package main

import "fmt"

// go 内置对多参数返回的支持
func vals() (int, int) {
    return 3, 7
}

func main() {

	// 接收多参数返回
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 通过 _, 获取第2个参数
	_, c := vals()
	fmt.Println(c)

}