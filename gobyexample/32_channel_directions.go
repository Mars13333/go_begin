package main

import "fmt"

// ping 只允许向 pings channel 发送数据（chan<- string）
func ping(pings chan<- string, msg string) {
	pings <- msg // 发送msg到pings
}

// pong 只允许从 pings channel 接收数据（<-chan string），只允许向 pongs channel 发送数据（chan<- string）
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // 从pings接收数据
	pongs <- msg   // 把接收到的数据发送到pongs
}

func main() {
	pings := make(chan string, 1) // 创建一个容量为1的string类型channel，pings
	pongs := make(chan string, 1) // 创建一个容量为1的string类型channel，pongs
	ping(pings, "passed message") // 调用ping，只能向pings发送数据
	pong(pings, pongs)            // 调用pong，只能从pings接收、向pongs发送
	fmt.Println(<-pongs)          // 从pongs接收数据并打印
}

/*
分析：
1. 通过 chan<- string 和 <-chan string，Go 可以限制函数参数的 channel 方向，提升代码安全性。
2. ping 只能发送，pong 只能接收/发送，main 里可以双向操作。
3. 这样可以避免误用，比如在 ping 里接收、在 pong 里发送，编译器会报错。
4. 这种写法常用于明确职责分工的并发场景。

补充解析：

1. 普通 channel 类型
chan T
- 既可以发送，也可以接收（双向 channel）。
- 例如：chan string，你可以 ch <- "hello" 也可以 msg := <-ch。

2. 只发送 channel
chan<- T
- 只能发送，不能接收。
- 例如：chan<- string，你只能 ch <- "hello"，不能 msg := <-ch。
- 用于函数参数时，表示这个函数只能往 channel 里发数据。

3. 只接收 channel
<-chan T
- 只能接收，不能发送。
- 例如：<-chan string，你只能 msg := <-ch，不能 ch <- "hello"。
- 用于函数参数时，表示这个函数只能从 channel 里收数据。

4. 例子对比
31_channel_synchronization.go
func worker(done chan bool) { ... }
- 这里 done 是双向 channel，既可以发送也可以接收。
- 但实际用法只用来发送信号。

32_channel_directions.go
func ping(pings chan<- string, msg string) { ... }
func pong(pings <-chan string, pongs chan<- string) { ... }
- ping 的参数 pings chan<- string：只能发送，不能接收。
- pong 的参数 pings <-chan string：只能接收，不能发送。
- pong 的参数 pongs chan<- string：只能发送，不能接收。

5. 为什么要这样写？
- 明确职责：让函数只做"发送"或"接收"，职责清晰。
- 防止误用：如果你在只能发送的 channel 上接收，编译器会报错。
- 提升安全性：并发编程中，方向明确可以减少 bug。

6. 总结
- <- 在 chan 左边：只能接收（<-chan T）。
- <- 在 chan 右边：只能发送（chan<- T）。
- 没有 <-：双向 channel，既能发也能收（chan T）。
*/
