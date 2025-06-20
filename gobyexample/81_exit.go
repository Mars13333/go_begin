package main

import (
	"fmt"
	"os"
)

func main() {

	// 注册一个defer函数
	// 重要：这个defer函数不会被执行！
	// 因为os.Exit()会立即终止程序，跳过所有defer函数
	defer fmt.Println("!")

	// 立即退出程序，返回状态码3
	// os.Exit()不会执行任何defer函数，也不会进行垃圾回收
	// 程序会立即终止，状态码3会传递给操作系统
	os.Exit(3)

	// 注意：这行代码永远不会被执行
	// fmt.Println("This will never be printed")
}

/*
Exit 程序退出示例主旨：

1. 核心功能：
   演示os.Exit()的使用方式以及它与defer语句的重要区别。

2. os.Exit()的特点：
   - 立即终止程序执行
   - 跳过所有defer函数
   - 不进行垃圾回收
   - 返回指定的退出状态码给操作系统
   - 不会执行任何清理工作

3. 退出状态码约定：
   ```go
   os.Exit(0)   // 成功退出
   os.Exit(1)   // 一般错误
   os.Exit(2)   // 误用shell命令
   os.Exit(126) // 命令无法执行
   os.Exit(127) // 命令未找到
   os.Exit(128) // 无效的退出参数
   // 128+n: 被信号n终止
   ```

4. defer vs os.Exit：
   ```go
   func demonstrateDefer() {
       defer fmt.Println("This will be printed")
       return // 正常返回，defer会执行
   }

   func demonstrateExit() {
       defer fmt.Println("This will NOT be printed")
       os.Exit(1) // 立即退出，defer不会执行
   }
   ```

5. 正确的退出模式：
   ```go
   func properExit() {
       defer cleanup() // 清理资源

       if err := doSomething(); err != nil {
           fmt.Printf("Error: %v\n", err)
           return // 使用return而不是os.Exit，让defer执行
       }

       fmt.Println("Success")
   }

   func main() {
       properExit()
       // 如果需要特定退出码，在main函数最后使用os.Exit
       // 或者让程序自然结束（退出码为0）
   }
   ```

6. 错误处理和退出：
   ```go
   func handleErrorsAndExit() {
       file, err := os.Open("config.txt")
       if err != nil {
           fmt.Printf("Failed to open config file: %v\n", err)
           os.Exit(1)
       }
       defer file.Close() // 这个defer会在os.Exit前执行

       // 处理文件...

       // 如果需要退出，使用return让defer执行
       if someCondition {
           fmt.Println("Condition met, exiting normally")
           return
       }
   }
   ```

7. 在不同场景中的退出处理：
   ```go
   // CLI应用程序
   func cliApp() {
       if len(os.Args) < 2 {
           fmt.Println("Usage: program <argument>")
           os.Exit(1)
       }

       result, err := processArgument(os.Args[1])
       if err != nil {
           fmt.Printf("Error: %v\n", err)
           os.Exit(2)
       }

       fmt.Println(result)
       // 正常退出，状态码0
   }

   // 服务器应用程序
   func serverApp() {
       server := setupServer()
       defer server.Close()

       // 设置信号处理
       c := make(chan os.Signal, 1)
       signal.Notify(c, os.Interrupt, syscall.SIGTERM)

       go func() {
           <-c
           fmt.Println("Shutting down server...")
           server.Close()
           os.Exit(0) // 优雅关闭后退出
       }()

       server.Start()
   }
   ```

8. 测试中的退出处理：
   ```go
   func TestSomething(t *testing.T) {
       // 在测试中不要使用os.Exit()
       // 使用t.Fatal()或t.Error()代替

       result, err := someFunction()
       if err != nil {
           t.Fatalf("Function failed: %v", err) // 而不是os.Exit(1)
       }

       if result != expected {
           t.Errorf("Expected %v, got %v", expected, result)
       }
   }
   ```

9. 资源清理模式：
   ```go
   func resourceCleanupPattern() {
       // 打开资源
       db, err := openDatabase()
       if err != nil {
           log.Printf("Failed to open database: %v", err)
           os.Exit(1)
       }
       defer db.Close()

       cache, err := openCache()
       if err != nil {
           log.Printf("Failed to open cache: %v", err)
           return // 让defer执行，关闭数据库
       }
       defer cache.Close()

       // 业务逻辑...

       // 正常结束，所有defer会按LIFO顺序执行
   }
   ```

10. 子程序中的退出处理：
    ```go
    func subFunction() error {
        // 在子函数中返回错误而不是直接退出
        if someCondition {
            return fmt.Errorf("condition not met")
        }
        return nil
    }

    func main() {
        if err := subFunction(); err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1) // 只在main函数中使用os.Exit
        }
    }
    ```

11. 退出状态码的检查：
    ```bash
    # 运行程序并检查退出状态码
    go run exit.go
    echo $?  # 输出: 3

    # 在shell脚本中使用
    if go run exit.go; then
        echo "Program succeeded"
    else
        echo "Program failed with exit code $?"
    fi
    ```

12. 优雅退出的完整示例：
    ```go
    func gracefulExit() {
        var exitCode int

        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("Panic recovered: %v\n", r)
                exitCode = 2
            }

            // 执行清理工作
            cleanup()

            // 最后退出
            if exitCode != 0 {
                os.Exit(exitCode)
            }
        }()

        // 主要业务逻辑
        if err := doWork(); err != nil {
            fmt.Printf("Work failed: %v\n", err)
            exitCode = 1
            return
        }

        fmt.Println("Work completed successfully")
    }
    ```

13. 信号处理与退出：
    ```go
    func signalHandlingExit() {
        c := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt, syscall.SIGTERM)

        go func() {
            sig := <-c
            fmt.Printf("Received signal: %v\n", sig)

            // 执行清理工作
            cleanup()

            // 根据信号设置适当的退出码
            switch sig {
            case os.Interrupt:
                os.Exit(130) // 128 + SIGINT(2)
            case syscall.SIGTERM:
                os.Exit(143) // 128 + SIGTERM(15)
            default:
                os.Exit(1)
            }
        }()

        // 主程序逻辑...
    }
    ```

14. 最佳实践：
    - 只在main函数中使用os.Exit()
    - 子函数应该返回错误而不是直接退出
    - 使用defer进行资源清理
    - 在需要defer执行时使用return而不是os.Exit()
    - 为不同的错误情况使用不同的退出码
    - 在测试中避免使用os.Exit()

15. 常见错误：
    ```go
    // 错误：在子函数中直接退出
    func badFunction() {
        if err := doSomething(); err != nil {
            os.Exit(1) // 不好：跳过了调用者的defer
        }
    }

    // 正确：返回错误让调用者处理
    func goodFunction() error {
        if err := doSomething(); err != nil {
            return fmt.Errorf("doSomething failed: %w", err)
        }
        return nil
    }
    ```

16. 调试技巧：
    ```go
    func debugExit() {
        defer func() {
            fmt.Println("This defer will execute")
        }()

        fmt.Println("About to exit")

        // 如果需要看到defer执行，使用return
        // return

        // 如果需要立即退出，使用os.Exit
        os.Exit(42)
    }
    ```

17. 跨平台注意事项：
    - 退出状态码在不同操作系统上的含义可能略有不同
    - Windows上的退出码范围是0-4294967295
    - Unix系统通常使用0-255范围

18. 性能考虑：
    - os.Exit()是最快的退出方式，因为跳过了清理
    - 但这可能导致资源泄露或数据丢失
    - 在性能关键的场景中，权衡速度和安全性

这个示例展示了程序退出的重要概念，
特别是os.Exit()与defer的关系，这是Go编程中的重要知识点。

测试方法：
1. 运行程序：go run exit.go
2. 检查退出状态码：echo $?
3. 观察defer函数不会被执行
*/
