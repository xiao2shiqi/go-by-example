package main

import (
	"fmt"
	"time"
)

// 使用 goroutines and channels 实现工作线程池
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		// 将工作结果接入 results channel
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	// 创建 2个缓存区 channel
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w < 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a < numJobs; a++ {
		<-results
	}
}
