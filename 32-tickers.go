package main

import (
	"fmt"
	"time"
)

// ticker 代表想要重复的做某一件事情
func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// 这里停顿 1600ms
	time.Sleep(2000 * time.Millisecond)
	ticker.Stop() // 关闭 ticker
	done <- true  // 关闭 select
	fmt.Println("Ticker stopped")
}
