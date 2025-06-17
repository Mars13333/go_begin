package main

import "fmt"

func main() {

	messages := make(chan string, 2) // ① 创建一个带缓冲区的 channel，容量为2

	messages <- "buffered" // ② 发送第1个数据
	messages <- "channel"  // ③ 发送第2个数据

	// messages <- "xxxxxx"
	/*
		如果尝试，发送第三个数据，则会报错：
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send]:
		main.main()
		E:/code/go_begin/gobyexample/30_channel_buffering.go:11 +0x58
		exit status 2

		你的 channel 容量是2，已经存了2个数据。
		你又尝试发送第3个数据 "xxxxxx"，此时缓冲区已满。
		由于没有其他 goroutine 在接收数据，发送操作会一直阻塞，主 goroutine 就卡在这里。
		Go 运行时检测到所有 goroutine 都在“等”，没有人能继续执行，于是报出：
		all goroutines are asleep - deadlock!
	*/

	fmt.Println(<-messages) // ④ 接收第1个数据
	fmt.Println(<-messages) // ⑤ 接收第2个数据
}

/*
什么是带缓冲的 channel？

普通（无缓冲）channel：make(chan T)，发送和接收必须同步，否则会阻塞。
带缓冲的 channel：make(chan T, N)，可以存储 N 个元素，发送时只要缓冲区没满就不会阻塞。
*/

/*
代码执行流程
messages := make(chan string, 2)
创建一个容量为2的 string 类型 channel。
这个 channel 最多可以存放 2 个数据，不需要立刻被接收。
messages <- "buffered"
向 channel 发送第1个数据，缓冲区还有空间，不会阻塞。
messages <- "channel"
发送第2个数据，缓冲区刚好满，也不会阻塞。
fmt.Println(<-messages)
从 channel 取出第1个数据 "buffered"，缓冲区腾出一个位置。
fmt.Println(<-messages)
取出第2个数据 "channel"，缓冲区再次变空。
*/

/*
带缓冲 channel 的特点
异步发送：只要缓冲区没满，发送操作不会阻塞。
异步接收：只要缓冲区有数据，接收操作不会阻塞。
同步机制：当缓冲区满时，发送会阻塞；当缓冲区空时，接收会阻塞。
适合生产者-消费者模型：生产者可以先把数据放进缓冲区，消费者慢慢取。
*/

/*
你可以这样理解
带缓冲的 channel 就像一个“带格子的信箱”，可以临时存放几封信。
只要信箱没满，投信的人（发送方）可以随时投递，不用等收信的人（接收方）来取。
信箱满了，投信的人就只能等收信的人来取走一封，才能继续投递。
*/
