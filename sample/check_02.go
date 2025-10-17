package main

import (
	"time"
)

func idCheck2(id int) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	println("\tgoroutine-", id, ":idCheck ok\n")
	return idCheckTmCost
}

func bodyCheck2(id int) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	println("\tgoroutine-", id, ":bodyCheck ok\n")
	return bodyCheckTmCost
}

func xRayCheck2(id int) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCHeckTmCost))
	println("\tgoroutine-", id, ":xRayCheck ok\n")
	return xRayCHeckTmCost
}

// airportSecurityCheck2 并行版本的机场安检总流程
func airportSecurityCheck2(id int) int {
	println("goroutine-", id, ":airportSecurityCheck ...\n")
	total := 0
	// 依次执行三个检查步骤（在当前goroutine内顺序执行）
	total += idCheck2(id)
	total += bodyCheck2(id)
	total += xRayCheck2(id)
	println("goroutine-", id, ":airportSecurityCheck ok\n")
	return total
}

// start 创建并启动一个goroutine工作池
// 参数说明：
//
//	id: goroutine编号，用于标识不同的工作goroutine
//	f: 要执行的函数（这里是airportSecurityCheck2）
//	queue: 接收任务的通道，<-chan struct{}表示只读通道
//
// 返回值：<-chan int 返回结果的通道
func start(id int, f func(int) int, queue <-chan struct{}) <-chan int {
	// 创建结果通道，用于goroutine向主线程返回数据
	c := make(chan int)

	// 启动一个新的goroutine（轻量级线程）
	go func() {
		total := 0 // 当前goroutine处理的总耗时
		for {
			// 从queue通道接收任务
			// queue通道关闭时，ok会变为false
			_, ok := <-queue
			if !ok {
				// 通道已关闭，向结果通道发送总耗时，然后退出goroutine
				c <- total
				return
			}
			// 收到任务，执行安检流程
			total += f(id)
		}
	}()
	return c
}

// max 返回多个整数中的最大值
// 用于计算三个并行通道中最长的处理时间
func max(args ...int) int {
	n := 0
	for _, v := range args {
		if v > n {
			n = v
		}
	}
	return n
}

// 方案2：并行方案
// 核心思想：增加安检通道，创建3个goroutine（轻量级线程），分别代表三个并行安检通道
// 每个通道可以独立处理乘客，实现真正的并行处理
func runCheckExample2() {
	passengers := 30 // 总乘客数

	// 创建任务分发通道
	// struct{}是最小的数据类型，这里只用作信号，不传输实际数据
	c := make(chan struct{})

	// 启动3个并行工作goroutine
	// 每个goroutine代表一个独立的安检通道
	c1 := start(1, airportSecurityCheck2, c) // 通道1
	c2 := start(2, airportSecurityCheck2, c) // 通道2
	c3 := start(3, airportSecurityCheck2, c) // 通道3

	// 向任务通道发送乘客任务
	// 每个struct{}{}代表一个需要安检的乘客
	for i := 0; i < passengers; i++ {
		c <- struct{}{} // 发送任务信号
	}
	close(c) // 关闭通道，通知所有goroutine任务完成

	// 从三个结果通道接收数据
	// <-c1会阻塞，直到对应的goroutine完成任务并发送结果
	total := max(<-c1, <-c2, <-c3) // 取三个通道的最大值（最慢的通道）
	println("total time cost:", total)
	// 并行方案结果：3600毫秒（比顺序方案的10800毫秒快3倍！）
}
