// go 构建是字段的集合体，对于数据的记录非常有用
package main

import "fmt"

// 通过 type，strcut 声明 person 实体
type person struct {
	// 声明 name, age 属性
	name string
	age int
}

// 传入参数 name 获取 person 实体
func newPerson(name string) *person {
	// 构造并且初始化 person
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// 打印构造的对象 (相比 java 无需声明 toString 即可打印对象)
	fmt.Println(person{"Bob", 20})
	// 也可以指定属性初始化对象
	fmt.Println(person{name: "Alice", age: 30})
	// 未指定初始化的属性，默认值为 0
	fmt.Println(person{name: "Fred"})
	// 暂时还不明白 &符号的意义在哪里
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	// 将创建对象赋值给变量 s
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	// 将 s 的引用传递给 sp
	sp := &s
	fmt.Println(sp.age)			// out: 50
	sp.age = 51
	fmt.Println(sp.age)			// out: 51
	fmt.Println(s.age)			// out: 51		s 的引用也被改变了...
	
}