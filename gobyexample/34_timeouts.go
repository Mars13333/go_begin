package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	/*
		goroutine 需要睡眠 2 秒才能向 c1 发送数据
		但是 time.After(time.Second) 只等待 1 秒就会发送超时信号
		所以 1 秒后，超时条件先满足，打印 "timeout 1"
	*/

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	/*
		goroutine 需要睡眠 2 秒才能向 c2 发送数据
		time.After(time.Second * 3) 等待 3 秒才超时
		所以 2 秒后，c2 先收到数据，打印 "result 2"，不会打印 "timeout 2"
	*/

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

	/*
		time.After(duration) 返回一个 channel，这个 channel 会在指定时间后接收到一个时间值
		// time.After 的简化实现原理
		func After(d time.Duration) <-chan time.Time {
				c := make(chan time.Time, 1)
				go func() {
						time.Sleep(d)
						c <- time.Now()  // 时间到了，发送当前时间到channel
				}()
				return c
		}

		所以 <-time.After(time.Second) 中：
		time.After(time.Second) 返回一个 channel
		<- 从这个 channel 接收数据
		1秒后，Go 运行时会向这个 channel 发送时间值
		select 语句接收到这个值，执行对应的 case

		所以 time.After 的作用就是为 select 提供一个"兜底"的超时机制，确保程序不会无限期地等待下去。
		或者可以使用default避免阻塞。
		select {
			case res := <-c1:
					fmt.Println(res)
			default:
					fmt.Println("no data available")  // 立即执行，不阻塞
		}
	*/
}
