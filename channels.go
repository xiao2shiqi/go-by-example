package main

// channels 是连接 goroutines 的管道
// 你可以发送 values 到 channels 中，然后另一个 goroutines 可以从 channels 中获取值
import "fmt"

func main() {

	// 使用 make(chan) 创建 channels
	messages := make(chan string)

	go func() {
		// 使用 <- 语法把字符放入 channels 中
		messages <- "ping"
	}()

	// 从 message channels 中获取另一个 groutines value，代替同步锁
	msg := <-messages
	fmt.Println(msg)
}
