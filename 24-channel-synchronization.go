// 通道同步
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	// 创建 done 通道，用于同步状态
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
