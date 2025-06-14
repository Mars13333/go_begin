package main

import (
	"fmt"
	"time"
)

/*并行方案*/

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCheckTMCost = 180
)

func idCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	fmt.Println("\tgoroutine-", id, ":idCheck ok\n")
	return idCheckTmCost
}

func bodyCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	fmt.Println("\tgoroutine-", id, ":bodyCheck ok\n")
	return bodyCheckTmCost
}

func xRayCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTMCost))
	fmt.Println("\tgoroutine-", id, ":xRayCheck ok\n")
	return xRayCheckTMCost
}

func airportSecurityCheck(id int) int {
	fmt.Println("goroutine-", id, ":airportSecurityCheck\n")
	total := 0
	total += idCheck(id)
	total += bodyCheck(id)
	total += xRayCheck(id)
	fmt.Println("goroutine-", id, ":airportSecurityCheck ok")
	return total
}

func start(id int, f func(int) int, queue <-chan struct{}) <-chan int {
	c := make(chan int)
	go func() {
		total := 0
		for {
			_, ok := <-queue
			if !ok {
				c <- total
				return
			}
			total += f(id)
		}
	}()
	return c
}

func max(args ...int) int {
	n := 0
	for _, v := range args {
		if v > n {
			n = v
		}
	}
	return n
}

func main() {
	total := 0
	passengers := 30

	c := make(chan struct{})
	c1 := start(1, airportSecurityCheck, c)
	c2 := start(2, airportSecurityCheck, c)
	c3 := start(3, airportSecurityCheck, c)

	for i := 0; i < passengers; i++ {
		c <- struct{}{}
	}
	close(c)
	total = max(<-c1, <-c2, <-c3)
	fmt.Println("total time cost:", total)
}

/*
...

total time cost: 3600
*/
