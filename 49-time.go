package main

import (
	"fmt"
	"time"
)

// golang 内置的 time 包提供对于时间处理的相关依赖
func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	// 使用 time.Date 构造自定义时间
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 打印 then 对象的年月日信息
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 星期几 ？
	p(then.Weekday())

	// 时间前后比较
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
