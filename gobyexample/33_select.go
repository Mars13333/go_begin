package main

import "fmt"
import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	// 这里的 () 是立即调用匿名函数的语法。
	// 如果没有最后的 ()，那么 func() { ... } 只是定义了一个函数，但没有执行它。

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
