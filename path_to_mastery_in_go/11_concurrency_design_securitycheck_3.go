package main

import (
	"fmt"
	"time"
)

/*并发方案*/

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCheckTMCost = 180
)

func idCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	fmt.Println("\tgoroutine-", id, ":idCheck ok\n")
	return idCheckTmCost
}

func bodyCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	fmt.Println("\tgoroutine-", id, ":bodyCheck ok\n")
	return bodyCheckTmCost
}

func xRayCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTMCost))
	fmt.Println("\tgoroutine-", id, ":xRayCheck ok\n")
	return xRayCheckTMCost
}

func start(id string, f func(string) int, next chan<- struct{}) (chan<- struct{}, chan<- struct{}, <-chan int) {
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

func newAirportSecurityCheckChannel(id string, queue <-chan struct{}) {
	go func(id string) {
		print("goroutine-", id, ": airportSecurityCheckChannel is ready...\n")
		// 启动X光检查
		queue3, quit3, result3 := start(id, xRayCheck, nil)

		// 启动人身检查
		queue2, quit2, result2 := start(id, bodyCheck, queue3)

		// 启动身份检查
		queue1, quit1, result1 := start(id, idCheck, queue2)

		for {
			select {
			case v, ok := <-queue:
				if !ok {
					close(quit1)
					close(quit2)
					close(quit3)
					total := max(<-result1, <-result2, <-result3)
					print("goroutine-", id, ": airportSecurityCheckChannel time cost:", total, "\n")
					print("goroutine-", id, ": airportSecurityCheckChannel closed\n")
					return
				}
				queue1 <- v
			}
		}
	}(id)
}

func max(args ...int) int {
	n := 0
	for _, v := range args {
		if v > n {
			n = v
		}
	}
	return n
}

func main() {
	passengers := 30
	queue := make(chan struct{}, 30)
	newAirportSecurityCheckChannel("channel1", queue)
	newAirportSecurityCheckChannel("channel2", queue)
	newAirportSecurityCheckChannel("channel3", queue)

	time.Sleep(5 * time.Second) // 保证上述三个goroutine都已经处于ready状态
	for i := 0; i < passengers; i++ {
		queue <- struct{}{}
	}
	time.Sleep(5 * time.Second)
	close(queue)                   // 为了打印各通道的处理时长
	time.Sleep(1000 * time.Second) // 防止main goroutine退出
}

/*
...
goroutine-channel2: airportSecurityCheckChannel time cost:2160
goroutine-channel2: airportSecurityCheckChannel closed
goroutine-channel3: airportSecurityCheckChannel time cost:2160
goroutine-channel3: airportSecurityCheckChannel closed
goroutine-channel1: airportSecurityCheckChannel time cost:1080
goroutine-channel1: airportSecurityCheckChannel closed
*/
