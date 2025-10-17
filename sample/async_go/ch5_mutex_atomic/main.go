package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	// sync.Mutex // rwmutex也可以
	number int32
}

func (c *Counter) Increment(i int) {
	// c.Mutex.Lock()
	// defer c.Mutex.Unlock()
	// c.number += i
	atomic.AddInt32(&c.number, int32(i))

}

// data race检查，go run --race main.go
// 推荐使用test文件 main_test.go来做单元测试
// 10+9+8+...+1
func main() { 
	counter := &Counter{}
	wg := &sync.WaitGroup{}
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			counter.Increment(n)
		}(i)
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter.number)
}
