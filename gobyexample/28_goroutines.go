package main

/*

goroutine 是 Go 语言中的轻量级线程，由 Go 运行时（runtime）调度和管理。

启动方式：在函数调用前加 go 关键字即可。
*/

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 同步执行 f("direct")
	f("direct")

	// 启动一个新的 goroutine，异步执行 f("goroutine")
	go f("goroutine")

	// 启动一个匿名 goroutine，异步打印 going。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 主 goroutine 等待 1 秒，保证其他 goroutine 有机会执行完毕。
	time.Sleep(time.Second)
	// 主 goroutine 打印 done。
	fmt.Println("done")

}

/*

goroutine 是并发执行的，调度顺序由 Go 运行时决定。
你看到的三种结果，都是可能的。
如果主 goroutine 不等待（比如去掉 time.Sleep），有些 goroutine 可能还没执行完，程序就退出了

goroutine 就像“分身”，你让它去做一件事，不等它做完，自己可以继续做别的事。
但如果你想等所有分身都做完再走，需要用 time.Sleep 或更好的同步方法（如 sync.WaitGroup）。
*/

/*
结果1：
direct : 0
direct : 1
direct : 2
going
goroutine : 0
goroutine : 1
goroutine : 2
done
*/
/*
结果2：
direct : 0
direct : 1
direct : 2
going
goroutine : 0
goroutine : 1
goroutine : 2
done
*/
/*
结果3：
direct : 0
direct : 1
direct : 2
goroutine : 0
goroutine : 1
goroutine : 2
going
done
*/
