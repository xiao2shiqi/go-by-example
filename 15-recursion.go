package main

import "fmt"

// go 对于递归的使用
func fact(n int) int {
	// syntax error: unexpected return, expecting
	if n == 0 {
		return 1
	} 
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}