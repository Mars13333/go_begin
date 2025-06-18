package main

import "fmt"

func main() {
	// 创建一个容量为2的string型缓冲通道
	queue := make(chan string, 2)
	// 向通道中发送两个元素，由于有缓冲区，不会阻塞
	queue <- "one"
	queue <- "two"
	// 关闭通道，表示不再发送数据
	close(queue)

	// 使用range遍历通道，直到通道被关闭且数据取完为止
	// 这里接收不会阻塞，因为通道中已有数据
	for elem := range queue {
		fmt.Println(elem)
	}
	// 如果没有接收方，且通道缓冲区满，再发送会阻塞，最终导致死锁
	// 如果没有发送方，接收方会在通道为空时阻塞，直到有数据或通道关闭
}
