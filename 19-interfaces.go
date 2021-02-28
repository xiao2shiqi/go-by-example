// interfaces 在 go 中的使用实例
// 官方描述 interfaces 是方法的集合：Interfaces are named collections of method signatures.
// 总结：struct 实现接口的所有方法，自动被 go 视为自动实现该接口，移除接口方法，则自动解除接口实现
package main

import (
	"fmt"
	"math"
)

// define geometry interface 定义接口
type geometry interface {
	// define method signatures.  在接口声明方法
	area() float64
	perim() float64		// 方法返回类型为 float64
}


// 定义 rect struct (可以理解为 java 中的 class)
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 声明 rect 的方法，实现 geometry 接口的 area 方法
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * r.width + 2 * r.height
}

// 声明 circle 的方法，实现 geometry 接口的 area 方法
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 接口方法，符合 geometry 皆可调用
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	// 构造 r 对象
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// 因为实现了 geometry 接口，所以他们都可以直接复用 measure 函数
	measure(r)		// 78.53981633974483
	measure(c)		// 31.41592653589793
}