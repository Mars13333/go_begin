package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个2秒后触发的定时器
	timer1 := time.NewTimer(2 * time.Second)

	// 阻塞等待timer1的C通道收到信号，表示定时器到期
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// 创建一个1秒后触发的定时器
	timer2 := time.NewTimer(time.Second)
	// 启动一个goroutine等待timer2到期
	go func() {
		// 如果timer2未被提前Stop，这里会在1秒后收到信号并打印
		// 但本例中timer2很快被Stop，所以这里永远不会收到信号
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// time.Sleep(2 * time.Second) // 如果加入这一行，timer2会正常触发，打印Timer 2 fired

	// 在定时器到期前尝试停止timer2
	stop2 := timer2.Stop()
	if stop2 {
		// Stop成功，说明timer2还没到期，定时器被成功取消
		fmt.Println("Timer 2 stopped")
	}

	// 主协程等待2秒，确保所有输出都能显示
	time.Sleep(2 * time.Second)
}

/*
go func() 确实是立即启动一个新的协程（goroutine），但“立即”只是指协程被调度到可运行队列，
并不保证它会马上执行到 <-timer2.C 这一行。Go 的调度器会在主协程和新协程之间分配 CPU 时间，
但具体哪个先执行、执行到哪一步，是不确定的。在你的代码里，主协程执行到 timer2.Stop() 的速度非常快，
通常会在新协程还没来得及执行 <-timer2.C 之前就已经把 timer2 停掉了。
*/
