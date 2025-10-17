package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()
	// 创建可取消的上下文
	// ctx := context.Background()
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		ip, err := GetIp(ctx)
		fmt.Println("IP:", ip, "Error:", err)
	}()

	// 模拟 2 秒后手动取消
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	wg.Wait()
	fmt.Println("执行完成...", time.Since(start))
}

func GetIp(ctx context.Context) (ip string, err error) {
	select {
	case <-ctx.Done(): // 等待取消
		fmt.Println("协程取消:", ctx.Err())
		return "", ctx.Err()
	case <-time.After(4 * time.Second): // 模拟任务耗时
		ip = "192.168.1.1"
		return ip, nil
	}
}
