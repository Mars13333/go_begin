package main

import (
	"log"
	"sync"
	"time"
)

func doSomething(id int) {
	log.Printf("before do job:(%d) \n", id)
	time.Sleep(3 * time.Second)
	log.Printf("after do job:(%d) \n", id)
}
func main1() {
	go doSomething(1)
	go doSomething(2)
	go doSomething(3)
	//当运行代码的时候，会发现没有任何输出。
	//因为我们的 main() 函数其实也是在一个 goroutine 中执行，但是 main() 执行完毕后，其他三个 goroutine 还没开始执行，所以就无法看到输出结果。
	//为了看到输出结果，我们可以使用 time.Sleep() 方法让 main() 函数延迟结束，例如：
	time.Sleep(4 * time.Second)
	/*
		可以看到，执行完所有任务从原本的 9 秒下降到 3 秒，大大提高了我们的效率，根据打印输出结果还可以看出：
			多个 goroutine 的执行是随机。
			对于 IO 密集型任务特别有效，比如文件，网络读写。
	*/
}

// 上面例子中，其实我们还可以使用 sync.WaitGroup 来等待所有的 goroutine 结束，从而实现并发的同步，这比使用 time.Sleep() 更加优雅，例如：
func doSomething2(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("before do job:(%d) \n", id)
	time.Sleep(3 * time.Second)
	log.Printf("after do job:(%d) \n", id)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go doSomething2(1, &wg)
	go doSomething2(2, &wg)
	go doSomething2(3, &wg)

	wg.Wait()
	log.Printf("finish all jobs\n")
}
