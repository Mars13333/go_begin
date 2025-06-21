package main

import "fmt"

// Generate 函数：生成从2开始的连续整数序列
// 参数 ch: 只写channel，用于发送整数
func Generate(ch chan<- int) {
	for i := 2; ; i++ { // 从2开始无限循环生成整数
		ch <- i // 将整数发送到channel
	}
}

// Filter 函数：过滤掉能被prime整除的数
// 参数 in: 只读channel，接收待过滤的整数
// 参数 out: 只写channel，发送过滤后的整数
// 参数 prime: 素数，用于过滤其倍数
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in         // 从输入channel接收整数
		if i%prime != 0 { // 如果该数不能被prime整除
			out <- i // 则发送到输出channel
		}
		// 如果该数能被prime整除，则丢弃（过滤掉）
	}
}

func main() {
	// 创建初始channel，用于接收Generate函数生成的整数
	ch := make(chan int)

	// 启动Generate goroutine，开始生成整数序列
	go Generate(ch)

	// 获取前10个素数
	for i := 0; i < 10; i++ {
		prime := <-ch          // 从当前channel接收一个素数
		fmt.Print(prime, "\n") // 打印素数

		// 为当前素数创建一个新的Filter goroutine
		ch1 := make(chan int)     // 创建新的channel
		go Filter(ch, ch1, prime) // 启动Filter goroutine过滤当前素数的倍数
		ch = ch1                  // 将ch指向新的channel，继续下一轮筛选
	}
}

/*
素数筛算法实现原理：

这是一个经典的并发素数筛算法，使用Go的goroutine和channel实现。

算法流程：
1. Generate函数生成从2开始的连续整数序列
2. 对于每个找到的素数p，创建一个Filter goroutine
3. Filter goroutine过滤掉所有能被p整除的数
4. 剩余的数为下一个素数，重复步骤2-4

并发设计：
- 每个素数都有自己的Filter goroutine
- 多个Filter goroutine通过channel串联形成管道
- 数据在管道中流动，每个Filter负责过滤特定素数的倍数

优势：
- 充分利用Go的并发特性
- 代码简洁优雅
- 性能高效，每个素数独立处理

示例输出：2, 3, 5, 7, 11, 13, 17, 19, 23, 29
*/
