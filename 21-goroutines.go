package main

// goroutine 是一个轻量级线程执行器
import (
	"fmt"
	"time"
)

// 一个输出函数，输出 0 1 2
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	// 通过关键字 go 即可启动线程指定 f() 函数
	go f("goroutine")

	go func(msg string) {
		// out print: golang
		fmt.Println(msg)
	}("going")

	// 阻塞主线程 1S 观察 go 线程执行情况
	time.Sleep(time.Second)
	fmt.Println("done")
}
