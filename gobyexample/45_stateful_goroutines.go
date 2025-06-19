package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// readOp 表示读操作的请求结构
type readOp struct {
	key  int      // 要读取的键
	resp chan int // 响应通道，用于返回读取的值
}

// writeOp 表示写操作的请求结构
type writeOp struct {
	key  int       // 要写入的键
	val  int       // 要写入的值
	resp chan bool // 响应通道，用于确认写入完成
}

func main() {

	// 使用原子计数器统计操作次数
	var readOps uint64  // 读操作计数器
	var writeOps uint64 // 写操作计数器

	// 创建通道用于传递读写操作请求
	reads := make(chan readOp)   // 读操作请求通道
	writes := make(chan writeOp) // 写操作请求通道

	// 启动状态管理 goroutine - 这是整个程序的核心
	go func() {
		var state = make(map[int]int) // 共享状态：存储键值对的 map
		for {
			select {
			case read := <-reads: // 处理读操作请求
				read.resp <- state[read.key] // 从 state 读取值并通过响应通道返回
			case write := <-writes: // 处理写操作请求
				state[write.key] = write.val // 将值写入 state
				write.resp <- true           // 通过响应通道确认写入完成
			}
		}
	}()

	// 启动 100 个读操作 goroutine
	for range 100 {
		go func() {
			for {
				// 创建读操作请求
				read := readOp{
					key:  rand.Intn(5),   // 随机选择 0-4 的键
					resp: make(chan int)} // 创建响应通道
				reads <- read                 // 发送读请求到状态管理 goroutine
				<-read.resp                   // 等待读操作完成并接收结果
				atomic.AddUint64(&readOps, 1) // 原子地增加读操作计数
				time.Sleep(time.Millisecond)  // 短暂休眠，模拟实际工作
			}
		}()
	}

	// 启动 10 个写操作 goroutine
	for range 10 {
		go func() {
			for {
				// 创建写操作请求
				write := writeOp{
					key:  rand.Intn(5),    // 随机选择 0-4 的键
					val:  rand.Intn(100),  // 随机选择 0-99 的值
					resp: make(chan bool)} // 创建响应通道
				writes <- write                // 发送写请求到状态管理 goroutine
				<-write.resp                   // 等待写操作完成确认
				atomic.AddUint64(&writeOps, 1) // 原子地增加写操作计数
				time.Sleep(time.Millisecond)   // 短暂休眠，模拟实际工作
			}
		}()
	}

	// 让程序运行 1 秒，收集操作统计
	time.Sleep(time.Second)

	// 原子地读取最终的操作计数并输出
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}

/*
Stateful Goroutines 有状态 Goroutine 示例主旨：

1. 核心思想：
   使用单个 goroutine 来管理共享状态，而不是使用互斥锁。
   这种模式通过消息传递来实现并发安全，符合 Go 的设计哲学：
   "不要通过共享内存来通信，而要通过通信来共享内存"

2. 架构模式：
   - 状态管理器：单个 goroutine 拥有并管理所有共享状态
   - 客户端：多个 goroutine 通过通道发送请求来访问状态
   - 请求-响应模式：每个操作都有对应的响应通道

3. 关键组件：
   - readOp/writeOp 结构体：封装操作请求和响应通道
   - 状态管理 goroutine：通过 select 处理读写请求
   - 客户端 goroutine：发送请求并等待响应

4. 优势：
   - 无需显式锁：避免了互斥锁的复杂性和潜在死锁
   - 数据竞争免疫：只有一个 goroutine 访问状态
   - 清晰的所有权：状态的所有权明确属于管理 goroutine
   - 易于理解：请求-响应模式直观易懂

5. 性能特点：
   - 所有状态访问都串行化：可能成为性能瓶颈
   - 适合状态访问不是主要瓶颈的场景
   - 通道通信有一定开销：比直接内存访问慢

6. 适用场景：
   - 状态相对简单且访问频率适中
   - 需要避免锁的复杂性
   - 状态访问逻辑复杂，难以用简单锁保护
   - 需要对状态访问进行额外控制（如限流、日志等）

7. 与互斥锁方案的对比：
   - 互斥锁：允许并发读，但需要小心处理锁的获取和释放
   - 有状态 goroutine：完全串行化，但逻辑更清晰，无死锁风险

这个示例展示了 Go 语言中处理并发状态管理的另一种重要模式。
*/
