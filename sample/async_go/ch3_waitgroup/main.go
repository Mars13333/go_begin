package main

import (
	"fmt"
	"sync"
	"time"
)

func doSome(number int) {
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("result %d\n", number)
}

func asyncDoSome(number int, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("result %d\n", number)
	wg.Done() // 等待-1
}

func main() {
	start := time.Now()
	wg := &sync.WaitGroup{} // 一般都用取地址符号&，因为后面要传到子goroutine中使用
	// 三个方法：add done wait
	// add 增加等待的goroutine数量
	// done 某个goroutine完成时调用，等待-1
	// wait 阻塞等待，直到等待数量为0
	for i := 0; i < 10; i++ {
		wg.Add(1)              // 增加等待的goroutine数量
		go asyncDoSome(13, wg) // 启用子goroutine
	}
	// 为了让子go程执行完毕，让main协程睡眠1秒
	// 这种代码是错误的，应该用sync.WaitGroup来等待所有子go程结束
	// time.Sleep(time.Millisecond * 109)

	wg.Wait() // 等待所有子goroutine完成
	fmt.Println("use time: ", time.Since(start))
	fmt.Println("end game")
}
