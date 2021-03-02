package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 通道关闭后无法在发送消息
	// queue <- "three"			// panic: send on closed channel

	// 通过 range 迭代 chaneel 示例（关闭后仍然可以接收）
	for elem := range queue {
		fmt.Println(elem)
	}
}
