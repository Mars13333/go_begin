package main

import (
	"fmt"
	"time"
)

func main() {
	messageCh := make(chan int, 100)
	for i := 0; i < 10; i++ {
		messageCh <- i // buff 100, 写入10个，安全
	}
	//增加close,避免deadlock
	close(messageCh)
	// 关闭后，不能再写入，否则panic
	// messageCh <- 100 // panic: send on closed channel
	// range 进行读取
	// for item := range messageCh {
	// 	fmt.Println("message number:", item)
	// }
	// 现在执行，看没有close的效果
	// 可以看到，读取了所有内容0到9后就deadlock了。
	// 如果想要避免deadlock,可以直接在发送方的goroutine中关闭即可

	// range模拟
	for {
		time.Sleep(time.Second)
		item, ok := <-messageCh
		if !ok {
			break // 到9了就退出
		}
		fmt.Println(item, ok)
	}
	// 现在执行，可以看到，读取了所有内容0到9都是true后，ok变成false,然后就退出了
	fmt.Println("End main goroutine")
}
