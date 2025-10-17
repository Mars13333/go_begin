package main

import (
	"fmt"
	"time"
)

func doSome(number int) {
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("result %d\n", number)
}

func doSomeBack(number int) string {
	time.Sleep(time.Millisecond * 100)
	return fmt.Sprintf("result %d\n", number)
}

// 如果需要子goroutine返回的内容，需要一个chan作为形参的函数！！！
// 这是go中常用的模式
func asyncDoSomeBack(number int, resCh chan string) {
	time.Sleep(time.Millisecond * 100)
	resCh <- fmt.Sprintf("result %d\n", number)
}

func main() {
	// 未用go
	start := time.Now()
	for i := 0; i < 10; i++ {
		doSome(13)
	}
	fmt.Println("use time: ", time.Since(start))
	fmt.Println("----------------------")
	// 并发的写法
	start = time.Now()
	for i := 0; i < 10; i++ {
		go doSome(13) // 启用子goroutine
	}
	// 为了让子go程执行完毕，让main协程睡眠1秒
	// 其实正常应该用sync.WaitGroup来等待所有子go程结束
	time.Sleep(time.Millisecond * 109)
	fmt.Println("go use time: ", time.Since(start))
	fmt.Println("----------------------")
	// 带返回值
	fmt.Println("同步：", doSomeBack(13))
	// 获得子goroutine的返回： 通过channel(chan)！ 它就是解决goroutine间通信的管道
	resCh := make(chan string) // 创建一个传string的管道 unbuffer的
	// 当主goroutine读取或者写入resChan的时候，是会报错的
	// resCh <- "foo" //写入unbuffer error!!!
	// v := <-resCh   //读取unbuffer error!!!
	// fmt.Println(v)
	// 因为是unbuffer,没有缓冲的，相当于读和写的时候是空的所以会报错

	// 在同一个goroutine读和写报错，但是不再一个里面就不会报错了
	go func() {
		resCh <- "foo" //写入unbuffer 不会报错

	}()
	res := <-resCh //读取unbuffer 不会报错
	fmt.Println(res)
	fmt.Println("End main goroutine")

	// 如果需要子goroutine返回的内容，需要一个chan作为形参的函数！！！
	resCh2 := make(chan string)
	// go asyncDoSomeBack(13, resCh2) // 在子goroutine中写入
	// 匿名函数写法
	go func(number int, resCh chan string) {
		time.Sleep(time.Millisecond * 100)
		resCh <- fmt.Sprintf("result %d\n", number)
	}(14, resCh2)
	// 在主goroutine中读取
	res2 := <-resCh2
	fmt.Println("从子goroutine中读取到的值：", res2)
}
