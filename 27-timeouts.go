package main

import (
	"fmt"
	"time"
)

// goroutine 下的超时机制 timeouts 使用
func main() {

	c1 := make(chan string, 1)

	go func() {
		// c1 goroutines 等待 2s
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
		// 等待 1s 超时，c1 的 2S 处理将会超时
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1") // out: timeout 1
	}

	c2 := make(chan string, 1)
	go func() {
		// c2 等待 2s 输出结果
		time.Sleep(2 * time.Second)
		c2 <- "result 2" // out: result 2
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
		// 等待 3秒超时，c2 将会处理成功
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}
