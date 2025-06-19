package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// 创建一个原子计数器，类型为 atomic.Uint64
	// 原子操作保证在并发环境下的线程安全，无需额外的锁机制
	var ops atomic.Uint64

	// 创建 WaitGroup 用于等待所有 goroutine 完成
	var wg sync.WaitGroup

	// 启动 50 个 goroutine，每个 goroutine 执行 1000 次计数操作
	// 总共会执行 50 * 1000 = 50000 次原子加法操作
	for range 50 {
		wg.Add(1) // 每启动一个 goroutine，计数器加 1

		go func() {
			// 每个 goroutine 内部执行 1000 次原子加法
			for range 1000 {
				// ops.Add(1) 是原子操作，线程安全
				// 相当于 ops++ 但在并发环境下安全
				ops.Add(1)
			}

			// 当前 goroutine 完成，通知 WaitGroup
			wg.Done()
		}()
	}

	// 等待所有 50 个 goroutine 完成
	wg.Wait()

	// ops.Load() 原子地读取计数器的值
	// 最终输出应该是 50000（50个goroutine * 1000次操作）
	fmt.Println("ops:", ops.Load())
}

/*
原子计数器的优势：

1. 线程安全：多个 goroutine 同时操作同一个变量时不会出现竞态条件
2. 性能优异：比使用 mutex 锁的方式更高效
3. 简单易用：不需要手动加锁解锁

如果不使用原子操作，而是使用普通的 int 变量：
var ops int
ops++  // 在并发环境下不安全，可能导致数据竞争

或者使用 mutex：
var ops int
var mu sync.Mutex
mu.Lock()
ops++
mu.Unlock()  // 安全但性能较差

原子操作是处理简单并发计数的最佳选择。
*/
