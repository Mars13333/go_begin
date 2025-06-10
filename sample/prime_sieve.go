package main

import "fmt"

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}

/*
素数筛：

Go版本程序实现了一个并发素数筛，它采用的是goroutine的并发组合。
程序从素数2开始，依次为每个素数建立一个goroutine，用于作为筛除该素数的倍数。
ch指向当前最新输出素数所位于的筛子goroutine的源channel。
*/
