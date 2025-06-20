package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// 创建一个用于接收信号的channel
	// 缓冲大小为1，确保不会丢失信号
	sigs := make(chan os.Signal, 1)

	// 注册要监听的信号
	// SIGINT: 中断信号（通常是Ctrl+C）
	// SIGTERM: 终止信号（通常用于优雅关闭）
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 创建一个用于通知程序完成的channel
	done := make(chan bool, 1)

	// 启动一个goroutine来处理信号
	go func() {

		// 等待接收信号
		// 这个操作会阻塞，直到收到注册的信号
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig) // 打印接收到的信号名称
		done <- true     // 通知主goroutine程序应该退出
	}()

	fmt.Println("awaiting signal")
	// 主goroutine等待信号处理完成
	// 程序会一直运行，直到收到信号
	<-done
	fmt.Println("exiting")
}

/*
Signals 信号处理示例主旨：

1. 核心功能：
   演示如何在Go程序中捕获和处理操作系统信号，实现优雅关闭和信号响应。

2. 信号基础概念：
   - 信号是Unix系统中进程间通信的一种方式
   - 操作系统可以向进程发送信号来通知特定事件
   - 进程可以选择忽略、捕获或使用默认行为处理信号

3. 常见信号类型：
   ```go
   syscall.SIGINT   // 中断信号（Ctrl+C）
   syscall.SIGTERM  // 终止信号（kill命令默认信号）
   syscall.SIGKILL  // 强制终止信号（不能被捕获）
   syscall.SIGHUP   // 挂起信号（终端断开）
   syscall.SIGUSR1  // 用户自定义信号1
   syscall.SIGUSR2  // 用户自定义信号2
   syscall.SIGQUIT  // 退出信号（Ctrl+\）
   ```

4. signal.Notify的使用：
   ```go
   // 监听特定信号
   signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

   // 监听所有信号（不推荐）
   signal.Notify(sigs)

   // 停止监听信号
   signal.Stop(sigs)

   // 重置信号处理为默认行为
   signal.Reset(syscall.SIGINT)
   ```

5. 优雅关闭模式：
   ```go
   func gracefulShutdown() {
       c := make(chan os.Signal, 1)
       signal.Notify(c, os.Interrupt, syscall.SIGTERM)

       // 启动服务器或其他资源
       server := startServer()

       // 等待信号
       <-c
       fmt.Println("Shutting down gracefully...")

       // 清理资源
       server.Close()
       // 关闭数据库连接
       // 保存状态等

       fmt.Println("Shutdown complete")
   }
   ```

6. HTTP服务器优雅关闭：
   ```go
   func httpServerWithGracefulShutdown() {
       server := &http.Server{Addr: ":8080"}

       go func() {
           if err := server.ListenAndServe(); err != http.ErrServerClosed {
               log.Fatalf("Server failed: %v", err)
           }
       }()

       // 等待中断信号
       c := make(chan os.Signal, 1)
       signal.Notify(c, os.Interrupt)
       <-c

       // 优雅关闭
       ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
       defer cancel()

       if err := server.Shutdown(ctx); err != nil {
           log.Fatalf("Server shutdown failed: %v", err)
       }

       log.Println("Server shutdown complete")
   }
   ```

7. 多信号处理：
   ```go
   func handleMultipleSignals() {
       sigs := make(chan os.Signal, 1)
       signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

       for {
           sig := <-sigs
           switch sig {
           case syscall.SIGINT:
               fmt.Println("Received SIGINT, shutting down...")
               return
           case syscall.SIGTERM:
               fmt.Println("Received SIGTERM, shutting down...")
               return
           case syscall.SIGHUP:
               fmt.Println("Received SIGHUP, reloading config...")
               // 重新加载配置
               reloadConfig()
           }
       }
   }
   ```

8. 信号处理的最佳实践：
   ```go
   func signalHandlingBestPractices() {
       // 使用带缓冲的channel避免丢失信号
       sigs := make(chan os.Signal, 1)

       // 只监听需要的信号
       signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

       // 使用context控制超时
       ctx, cancel := context.WithCancel(context.Background())
       defer cancel()

       go func() {
           <-sigs
           fmt.Println("Signal received, initiating shutdown...")
           cancel()
       }()

       // 主程序逻辑
       select {
       case <-ctx.Done():
           fmt.Println("Shutting down...")
       case <-time.After(60 * time.Second):
           fmt.Println("Timeout, shutting down...")
       }
   }
   ```

9. 工作池的优雅关闭：
   ```go
   func workerPoolWithSignals() {
       jobs := make(chan int, 100)
       results := make(chan int, 100)

       // 启动工作器
       for i := 0; i < 3; i++ {
           go worker(jobs, results)
       }

       // 信号处理
       sigs := make(chan os.Signal, 1)
       signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

       go func() {
           <-sigs
           fmt.Println("Stopping workers...")
           close(jobs) // 关闭作业channel，工作器会退出
       }()

       // 发送作业
       for i := 1; i <= 100; i++ {
           select {
           case jobs <- i:
           case <-sigs:
               close(jobs)
               return
           }
       }

       close(jobs)
   }
   ```

10. 日志轮转信号处理：
    ```go
    func logRotationSignal() {
        sigs := make(chan os.Signal, 1)
        signal.Notify(sigs, syscall.SIGUSR1)

        logFile := openLogFile("app.log")
        defer logFile.Close()

        go func() {
            for range sigs {
                fmt.Println("Rotating log file...")
                logFile.Close()
                // 重命名当前日志文件
                os.Rename("app.log", "app.log.old")
                // 打开新的日志文件
                logFile = openLogFile("app.log")
            }
        }()

        // 应用程序逻辑...
    }
    ```

11. 信号忽略：
    ```go
    func ignoreSignals() {
        // 忽略特定信号
        signal.Ignore(syscall.SIGHUP)

        // 或者注册一个空的处理器
        signal.Notify(make(chan os.Signal, 1), syscall.SIGHUP)
    }
    ```

12. 子进程信号处理：
    ```go
    func handleChildProcessSignals() {
        // 处理子进程退出信号
        sigs := make(chan os.Signal, 1)
        signal.Notify(sigs, syscall.SIGCHLD)

        go func() {
            for range sigs {
                // 回收僵尸进程
                for {
                    pid, err := syscall.Wait4(-1, nil, syscall.WNOHANG, nil)
                    if err != nil || pid == 0 {
                        break
                    }
                    fmt.Printf("Child process %d exited\n", pid)
                }
            }
        }()
    }
    ```

13. 测试信号处理：
    ```bash
    # 运行程序
    go run signals.go &
    PID=$!

    # 发送SIGINT信号
    kill -INT $PID

    # 或者使用Ctrl+C
    # 发送SIGTERM信号
    kill -TERM $PID

    # 发送自定义信号
    kill -USR1 $PID
    ```

14. 跨平台注意事项：
    ```go
    // Windows上的信号支持有限
    // 只支持os.Interrupt (Ctrl+C)
    func crossPlatformSignals() {
        sigs := make(chan os.Signal, 1)

        if runtime.GOOS == "windows" {
            signal.Notify(sigs, os.Interrupt)
        } else {
            signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
        }

        <-sigs
        fmt.Println("Shutting down...")
    }
    ```

15. 信号处理的注意事项：
    - 信号处理器应该尽快完成，避免阻塞
    - 不要在信号处理器中执行复杂操作
    - 使用channel在goroutine间传递信号
    - SIGKILL和SIGSTOP不能被捕获或忽略
    - 信号是异步的，处理时要考虑并发安全

16. 调试技巧：
    ```go
    func debugSignals() {
        sigs := make(chan os.Signal, 1)
        signal.Notify(sigs)

        go func() {
            for sig := range sigs {
                fmt.Printf("Received signal: %v (%d)\n", sig, sig)
                // 可以选择是否退出
                if sig == syscall.SIGINT || sig == syscall.SIGTERM {
                    os.Exit(0)
                }
            }
        }()
    }
    ```

17. 实际应用场景：
    - Web服务器优雅关闭
    - 后台服务和守护进程
    - 批处理任务的中断处理
    - 配置热重载
    - 日志轮转
    - 资源清理和状态保存

18. 性能考虑：
    - 使用适当大小的缓冲channel
    - 避免在信号处理器中执行耗时操作
    - 合理设置超时时间
    - 考虑信号处理的优先级

这个示例展示了Go中信号处理的基础模式，
是构建健壮系统服务的重要技能。

测试方法：
1. 运行程序：go run signals.go
2. 按Ctrl+C发送SIGINT信号
3. 观察程序如何优雅地响应和退出
*/
