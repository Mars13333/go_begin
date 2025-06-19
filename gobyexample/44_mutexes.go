package main

import (
	"fmt"
	"sync"
)

// Container 结构体包含一个互斥锁和一个计数器映射
type Container struct {
	mu       sync.Mutex     // 互斥锁，保护共享资源的并发访问
	counters map[string]int // 共享资源：存储不同名称的计数器
}

// inc 方法安全地增加指定名称的计数器
func (c *Container) inc(name string) {
	c.mu.Lock()         // 获取锁，确保独占访问
	defer c.mu.Unlock() // 函数结束时自动释放锁，即使发生 panic 也能释放
	c.counters[name]++  // 在锁保护下安全地修改共享数据
}

func main() {

	// 创建 Container 实例
	c := Container{
		// mu不需要初始化，因为sync.Mutex是零值可用的
		counters: map[string]int{"a": 0, "b": 0}, // 初始化计数器映射
	}

	var wg sync.WaitGroup // 用于等待所有 goroutine 完成

	// doIncrement 函数：对指定名称的计数器执行 n 次增加操作
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name) // 每次调用都会获取锁，保证线程安全
		}
		wg.Done() // 通知 WaitGroup 当前 goroutine 完成
	}

	// 启动 3 个 goroutine 并发执行
	wg.Add(3)                  // 设置等待 3 个 goroutine
	go doIncrement("a", 10000) // goroutine 1: 对 "a" 增加 10000 次
	go doIncrement("a", 10000) // goroutine 2: 对 "a" 增加 10000 次
	go doIncrement("b", 10000) // goroutine 3: 对 "b" 增加 10000 次
	wg.Wait()                  // 等待所有 goroutine 完成
	fmt.Println(c.counters)    // 输出结果：map[a:20000 b:10000]

}

/*
Mutex 互斥锁示例主旨：

1. 问题场景：
   多个 goroutine 并发访问共享资源（map）时，会发生数据竞争，
   导致结果不可预测或程序崩溃。

2. 解决方案：
   使用 sync.Mutex 互斥锁来保护共享资源，确保同一时刻只有
   一个 goroutine 能够访问和修改共享数据。

3. 关键概念：
   - Lock()：获取锁，如果锁已被占用则阻塞等待
   - Unlock()：释放锁，允许其他 goroutine 获取锁
   - defer Unlock()：确保函数结束时释放锁，防止死锁

4. 运行结果：
   - "a" 的最终值是 20000（两个 goroutine 各增加 10000）
   - "b" 的最终值是 10000（一个 goroutine 增加 10000）
   - 结果是确定的，不会因为并发而出现数据竞争

5. 性能考虑：
   - 互斥锁会降低并发性能（串行化访问）
   - 但保证了数据的正确性和一致性
   - 适用于复杂的共享数据结构（如 map、slice 等）

6. 与原子操作的对比：
   - 原子操作：适用于简单的数值计算（如计数器）
   - 互斥锁：适用于复杂的数据结构和多步操作
*/
