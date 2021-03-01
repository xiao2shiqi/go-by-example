package main

import "fmt"

// channels 默认无缓冲
func main() {

	// 创建 2个字符的缓冲区的 channels
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
