package main

import "fmt"

// if/else 使用示例
func main() {
	
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
        fmt.Println("7 is odd")			// print this line
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")		// println 8 is divisible by 4
    }

	// 初始化，并且进入 if/else 条件
	if num := 9; num < 0 {
		fmt.Println(num, "is nefative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")		// 输出这一行
	} else {
		fmt.Println(num, "has multiple digits")
	}
}