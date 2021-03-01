package main

// error 错误处理在 go 语言中的使用
import (
	"errors"
	"fmt"
)

// 按照 go 约定：error 对象作为最后一个参数返回
func f1(arg int) (int, error) {
	if arg == 42 {
		// 创建一个错误对象返回
		return -1, errors.New("can't work with 42")
	}
	// error 返回 nil 表示没有错误
	return arg + 3, nil
}

// f2 自定义 error 类型示例
type argError struct {
	arg  int
	prob string
}

// 实现 Error 方法，自动继承异常类
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// 构造自定义的 argError 异常对象
		// arg = int
		// prob = message
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// 传入 2个数验证结果
	for _, i := range []int{7, 42} {
		// 接收 r, e 作为返回值
		if r, e := f1(i); e != nil {
			// e != nil 进入条件
			fmt.Println("f1 failed:", e)
		} else {
			// 没有 error 则继续工作
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 workded:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		// 从 argError 对象中获取 arg, prob 属性
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
