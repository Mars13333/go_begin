package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个每500毫秒触发一次的Ticker
	ticker := time.NewTicker(500 * time.Millisecond)
	// 创建一个用于通知协程停止的通道
	done := make(chan bool)

	// 启动一个goroutine，循环监听ticker和done信号
	go func() {
		for {
			select {
			// 如果收到done信号，退出循环，结束协程
			case <-done:
				return
			// 每当ticker定时到达时，打印当前时间
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// 主协程等待1.6秒，期间ticker会触发多次
	time.Sleep(1600 * time.Millisecond)
	// 停止ticker，后续不会再有tick事件
	ticker.Stop()
	// 如果不写 done <- true，协程不会收到退出信号，会一直阻塞在 select
	// done <- true
	fmt.Println("Ticker stopped")
	// 程序主协程结束，整个程序会退出，但协程其实还在阻塞（如果主协程不结束会资源泄漏）
}
