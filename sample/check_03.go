package main

import (
	"time"
)

func idCheck3(id string) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	println("\tgoroutine-", id, ":idCheck ok\n")
	return idCheckTmCost
}

func bodyCheck3(id string) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	println("\tgoroutine-", id, ":bodyCheck ok\n")
	return bodyCheckTmCost
}

func xRayCheck3(id string) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCHeckTmCost))
	println("\tgoroutine-", id, ":xRayCheck ok\n")
	return xRayCHeckTmCost
}

func newAirportSecurityCheckChannel(id string, queue <-chan struct{}) {
	go func(id string) {
		println("goroutine-", id, ":airportSecurityCheckChannel is ready...\n")
		// 启动x光检查
		queue3, quit3, result3 := start3(id, xRayCheck3, nil)
		// 启动人身检查
		queue2, quit2, result2 := start3(id, bodyCheck3, queue3)
		// 启动身份检查
		queue1, quit1, result1 := start3(id, idCheck3, queue2)

		// 优化：使用for range代替for { select {} }
		// for range会自动处理通道关闭，代码更简洁
		for v := range queue {
			queue1 <- v
		}

		// 通道关闭后的清理工作
		close(quit1)
		close(quit2)
		close(quit3)
		total := max3(<-result1, <-result2, <-result3)
		println("goroutine-", id, ":airportSecurityCheckChannel time cost:", total, "\n")
		println("goroutine-", id, ":airportSecurityCheckChannel closed\n")
	}(id)
}

func start3(id string, f func(string) int, next chan<- struct{}) (chan<- struct{}, chan<- struct{}, <-chan int) {
	queue := make(chan struct{}, 10)
	quit := make(chan struct{})
	result := make(chan int)

	go func() {
		total := 0
		for {
			select {
			case <-quit:
				result <- total
				return
			case v := <-queue:
				total += f(id)
				if next != nil {
					next <- v
				}
			}
		}
	}()
	return queue, quit, result
}

// max 返回多个整数中的最大值
// 用于计算三个并行通道中最长的处理时间
func max3(args ...int) int {
	n := 0
	for _, v := range args {
		if v > n {
			n = v
		}
	}
	return n
}

// 方案3：并发方案
// 模拟开启了3条通道(newAirportSecurityCheckChannel)，每条通道创建3个goroutine
// 分别处理idCheck,bodyCheck,xRayCheck,3个goroutine之间通过channel相连
func runCheckExample3() {
	passengers := 30 // 总乘客数
	queue := make(chan struct{}, 30)
	newAirportSecurityCheckChannel("channel1", queue)
	newAirportSecurityCheckChannel("channel2", queue)
	newAirportSecurityCheckChannel("channel3", queue)
	time.Sleep(5 * time.Second) // 保证上述三个goroutine都已经处于ready状态
	for i := 0; i < passengers; i++ {
		queue <- struct{}{}
	}
	time.Sleep(5 * time.Second)
	close(queue)                  // 为了打印各通道的处理时长
	time.Sleep(100 * time.Second) // 防止main goroutine退出
}
