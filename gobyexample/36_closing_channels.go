package main

import "fmt"

func main() {
	// 创建一个容量为5的int型通道，用于传递任务
	jobs := make(chan int, 5)
	// 创建一个bool型通道，用于通知主协程所有任务已处理完毕
	done := make(chan bool)

	// 启动一个goroutine用于接收并处理jobs通道中的任务
	go func() {
		for {
			// 从jobs通道接收数据，more表示通道是否已关闭
			j, more := <-jobs
			if more {
				// 如果通道未关闭，打印接收到的任务
				fmt.Println("received job", j)
			} else {
				// 如果通道已关闭且数据已取完，打印提示并通过done通道通知主协程
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 向jobs通道发送3个任务
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// 关闭jobs通道，表示没有更多任务发送
	close(jobs)
	fmt.Println("sent all jobs")

	// 等待goroutine处理完所有任务并通知主协程
	<-done

	// 再次尝试从已关闭的jobs通道接收数据，ok为false
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
