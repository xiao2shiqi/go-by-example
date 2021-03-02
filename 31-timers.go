package main

import (
	"fmt"
	"time"
)

// timer 组件在 go 里面的应用
// timer 代表你想再未来某一个时间做的事情
func main() {

	// 创建 time 任务(创建任务不会阻塞线程)
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C                   // 程序将在这里阻塞 2S
	fmt.Println("Timer 1 fired") // out: timer 1 fired

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C                   // timer 在这里阻塞 1S
		fmt.Println("Timer 2 fired") // 因为 timer2 已关闭，在触发前已经停止了，所以这里不会输出
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
