package main

import (
	"fmt"
	"time"		// 时间库
)

// swtich 多条件判断语句
func main() {

	// 初始化 i
	i := 2
	fmt.Println("Write ", i , " as ")

	// switch 语法为 ： switch value { case: }
	// 多个 switch 匹配符
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")			// 输出
	case 3:
		fmt.Println("three")
	}

	// 根据当前时间输出
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		// 这是一个周末
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")		
	}

	// 根据当前时间输出
	t := time.Now()
	switch {		// switch 可以读外部变量
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")		
	}

	// 定义函数闭包，switch 通过传参判断
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
			// %T\n 获取类型的方式有些诡异.... 暂不深究
            fmt.Printf("Don't know type %T\n", t)
        }
    }

	whatAmI(true)		// I'm a bool
	whatAmI(1)			// I'm a int
	whatAmI("hey")		// Don't know type

}