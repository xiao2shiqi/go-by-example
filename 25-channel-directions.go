package main

import "fmt"

// 当使用 channel 作为函数的时候，可以限制通道的作用，增加代码安全性
// pings chan<- 示例：
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// 创建通道 ?
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
