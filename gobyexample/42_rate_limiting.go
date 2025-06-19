package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建一个带缓冲的channel，模拟5个请求
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 基础限流器：每200毫秒产生一个时间信号
	// time.Tick返回一个channel，定期发送当前时间
	limiter := time.Tick(200 * time.Millisecond)

	fmt.Println("=== 基础限流：每200ms处理一个请求 ===")
	// 处理每个请求时，都要先从limiter接收一个信号
	// 这样就限制了处理速度：每200ms只能处理一个请求
	for req := range requests {
		<-limiter // 等待限流器允许（阻塞200ms）
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("\n=== 突发限流：允许突发3个请求，然后限流 ===")
	
	// 突发限流器：带缓冲的channel，容量为3
	burstyLimiter := make(chan time.Time, 3)

	// 预先填充3个时间信号，允许前3个请求立即执行（突发处理）
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// 启动goroutine，每200ms向突发限流器添加一个信号
	// 这样在突发处理完后，后续请求仍然受到限流控制
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 创建另一组请求用于演示突发限流
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	
	// 处理突发请求
	// 前3个请求会立即执行（因为burstyLimiter预填充了3个信号）
	// 后2个请求会按200ms间隔执行（受限流控制）
	for req := range burstyRequests {
		<-burstyLimiter // 从突发限流器获取许可
		fmt.Println("request", req, time.Now())
	}
}
