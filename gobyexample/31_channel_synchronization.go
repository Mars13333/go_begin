package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true //工作完成后，向 channel 发送信号
}

func main() {
	done := make(chan bool, 1) //创建一个带缓冲的 channel，容量为1

	go worker(done) //启动 worker 协程

	<-done //（接收）主协程等待 worker 完成

}

//输出：working...done

/*
这个示例展示了哪些特点？

① channel 用于同步而不是传递数据
这里的 channel 类型是 chan bool，但实际上传递的 true 只是一个“信号”，表示“我干完了”。
这种用法常见于通知和同步，而不是数据交换。

② 主协程等待子协程完成
主协程执行到 <-done 时会阻塞，直到 worker 协程向 done 发送了 true。
这样可以保证主协程不会提前退出，确保 worker 的工作完成。

③ channel 的同步特性
即使 channel 是带缓冲的（make(chan bool, 1)），只要主协程还没接收，worker 协程发送后就能立刻返回。
如果 channel 是无缓冲的，worker 发送时会阻塞，直到主协程接收。

④ 代码输出顺序可控
由于主协程会等待 <-done，所以输出一定是：
working...done
之后主协程才会退出。



可以这样理解
这个 channel 就像一个“完成信号器”。
worker 干完活后按一下按钮（done <- true），主协程等着这个信号（<-done），收到信号后才继续。


总结
channel 不仅能传递数据，还能用来做同步和通知。
这种用法可以让主协程等待其他协程完成任务，常用于并发编程中的“等待所有任务完成”场景。
这种同步方式比 time.Sleep 更安全、可靠。
*/
