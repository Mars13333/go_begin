package main

import "fmt"

func main() {
	messages := make(chan string) // ① 创建一个 string 类型的 channel

	go func() { // ② 启动一个 goroutine
		messages <- "ping" // ③ 向 channel 发送数据
	}()

	msg := <-messages // ④ 从 channel 接收数据
	fmt.Println(msg)  // ⑤ 打印接收到的数据
}

/*
channel 是什么？

channel（通道）是 Go 语言中用于在多个 goroutine 之间传递数据的管道。
channel 是类型安全的，只能传递指定类型的数据。
channel 的本质是线程安全的队列，用于同步和通信。


messages <- "ping"

<- 箭头表示“发送”操作，把 "ping" 放进 channel。
只能在 goroutine 里发送，否则会 阻塞主线程!!!


msg := <-messages

<- 箭头表示“接收”操作，从 channel 取出一个值。
如果 channel 里没有数据，这里会 阻塞，直到有数据为止!!!


同步特性!!!

channel 的发送和接收操作是同步的，即：
如果没有 goroutine 在接收，发送操作会阻塞。
如果没有 goroutine 在发送，接收操作会阻塞。
这保证了数据的安全传递和同步。


channel 的主要特点总结

类型安全：只能传递指定类型的数据。
线程安全：多个 goroutine 可以安全地通过 channel 通信。
同步/阻塞：发送和接收操作默认是阻塞的，实现了 goroutine 之间的同步。
一对一通信：最基础的 channel 用法是一个发送方、一个接收方。
可以用于并发控制：比如任务分发、结果收集等。
*/
