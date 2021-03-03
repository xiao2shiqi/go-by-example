package main

import (
	"fmt"
	"os"
)

// defer 在函数执后执行，通常用于确保资源关闭，类似 java 的 finally
func main() {

	f := createFile("/tmp/defer.txt")
	// 延迟关闭文件
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	// 关闭文件检查错误很重要，即使在 defer 中
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
