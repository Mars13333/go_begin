package main

import (
	"fmt"
	"time"
)

// worker函数，代表一个工人，不断从jobs通道取任务，处理后将结果写入results通道
// jobs <-chan int 表示 jobs 是只读通道，只能接收任务
// results chan<- int 表示 results 是只写通道，只能发送结果
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { // 只能从 jobs 读取任务
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second) // 模拟任务耗时
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2 // 只能向 results 发送结果
	}
}

func main() {

	const numJobs = 5                  // 总共要处理的任务数
	jobs := make(chan int, numJobs)    // 任务通道，带缓冲区
	results := make(chan int, numJobs) // 结果通道，带缓冲区

	// 启动3个worker协程，每个worker都能从jobs通道取任务
	for w := 1; w <= 3; w++ { 
		go worker(w, jobs, results)
	}

	// 向jobs通道发送5个任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// 关闭jobs通道，通知所有worker没有新任务了
	close(jobs)

	// 主协程等待所有任务结果（5个），防止主协程提前退出
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

/*
生产环境建议加上超时、错误处理、监控等机制。


应用场景
worker pool（工作池）模式在实际开发中应用非常广泛，尤其适合高并发、批量任务处理、资源受限的场景。
下面举几个常见的业务场景

1. 批量数据处理
比如你要处理一批图片、视频、文件等，每个任务都可以独立完成，但总量很大。如果你开太多协程会耗尽资源，开太少又效率低。
worker pool 可以让你用固定数量的 worker 并发处理，既高效又不会压垮系统。
例子：
批量图片压缩/水印
批量导入/导出数据
批量发送邮件、短信

2. 网络爬虫/采集
爬虫需要抓取大量网页，但不能无限制并发，否则会被目标网站封禁或本地资源耗尽。
用 worker pool 控制并发数量，既能高效抓取，又能保护自己和目标网站。

3. 并发请求第三方API
比如你要给很多用户推送消息、同步数据到第三方服务，但第三方API有QPS限制。
worker pool 可以帮你平滑地并发请求，避免超限。

4. 日志/消息异步处理
日志、消息、事件等经常需要异步写入数据库、消息队列等。
worker pool 可以让你高效地异步处理这些任务，防止单个写入阻塞主流程。

5. 任务调度系统
比如定时任务、分布式任务调度，worker pool 可以让你灵活地分配和执行任务。
*/
