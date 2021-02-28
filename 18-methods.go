// 使用 go 在 type 上定义 method 实例
package main

import "fmt"

type rect struct {
	width, height int
}

// 定义 rect type 的 methdos 
// go 自动处理参数的传递类型：复制传递，引用传递
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2 * r.width + 2 * r.height
}

func main() {
	// 构造 rect 实体
	r := rect{width: 10, height: 5}

	// 调用 rect methods 
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// 传递 r 引用给 rp
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}