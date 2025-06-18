package main

import (
	"fmt"
	"sync"
	"time"
)

// worker函数，模拟一个耗时任务
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup // 创建WaitGroup，用于等待所有协程完成

	for i := 1; i <= 5; i++ {
		wg.Add(1) // 启动一个worker前，计数加1

		// 启动一个匿名协程
		go func() {
			// defer会在当前协程return前最后执行，保证无论正常结束还是异常退出都能调用Done
			// wg.Done() 并不关心“是谁”结束的，它只是把 WaitGroup 的计数减一。
			// WaitGroup 只关心还有多少个未完成的任务，而不关心具体是哪个任务完成。
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait() // 阻塞等待所有wg.Done()被调用，即所有worker完成

}
