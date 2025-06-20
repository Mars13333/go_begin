package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// 查找可执行文件的完整路径
	// exec.LookPath在系统PATH环境变量中搜索指定的可执行文件
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// 准备传递给新进程的参数列表
	// 注意：args[0]通常是程序名称本身，这是Unix约定
	args := []string{"ls", "-a", "-l", "-h"}

	// 获取当前进程的环境变量
	// os.Environ()返回当前进程的所有环境变量
	env := os.Environ()

	// 执行进程替换
	// syscall.Exec会用新程序替换当前进程
	// 这个调用如果成功，将不会返回（当前进程被完全替换）
	// 只有在出错时才会返回
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	// 注意：这行代码永远不会被执行
	// 因为syscall.Exec成功时会替换整个进程
	// fmt.Println("This will never be printed")
}

/*
Exec'ing Processes 进程替换示例主旨：

1. 核心功能：
   演示如何使用syscall.Exec进行进程替换，这是Unix系统编程的基础概念。

2. 进程替换 vs 进程生成：
   - 进程生成（spawn）：创建新的子进程，父进程继续存在
   - 进程替换（exec）：用新程序完全替换当前进程

3. syscall.Exec的特点：
   - 替换当前进程的内存映像
   - 保持相同的进程ID（PID）
   - 继承文件描述符（除非设置了FD_CLOEXEC）
   - 继承进程组和会话ID
   - 如果成功，永远不会返回

4. exec.LookPath的作用：
   ```go
   // 在PATH中查找可执行文件
   path, err := exec.LookPath("command")

   // 等价于手动指定完整路径
   err := syscall.Exec("/bin/ls", args, env)
   ```

5. 参数传递约定：
   ```go
   // args[0] 通常是程序名（约定）
   // args[1:] 是实际的命令行参数
   args := []string{"program", "arg1", "arg2"}

   // 对应命令行：program arg1 arg2
   ```

6. 环境变量处理：
   ```go
   // 使用当前环境
   env := os.Environ()

   // 自定义环境
   env := []string{
       "PATH=/bin:/usr/bin",
       "HOME=/home/user",
       "USER=myuser",
   }

   // 修改现有环境
   env := os.Environ()
   env = append(env, "NEW_VAR=value")
   ```

7. 错误处理：
   ```go
   err := syscall.Exec(binary, args, env)
   if err != nil {
       // 只有出错时才会到达这里
       switch err {
       case syscall.ENOENT:
           fmt.Println("File not found")
       case syscall.EACCES:
           fmt.Println("Permission denied")
       default:
           fmt.Printf("Exec failed: %v\n", err)
       }
   }
   ```

8. 与其他exec系列函数的关系：
   在C语言中有多个exec函数：
   - execl, execle, execlp
   - execv, execve, execvp

   Go的syscall.Exec对应C的execve：
   ```c
   // C语言等价调用
   execve(binary, args, env);
   ```

9. 实际应用场景：
   ```go
   // Shell实现中的命令执行
   func executeCommand(cmd string, args []string) {
       binary, err := exec.LookPath(cmd)
       if err != nil {
           fmt.Printf("Command not found: %s\n", cmd)
           return
       }

       fullArgs := append([]string{cmd}, args...)
       env := os.Environ()

       syscall.Exec(binary, fullArgs, env)
   }
   ```

10. 与fork结合使用（Unix模式）：
    ```go
    // 注意：Go不直接支持fork，这里是概念性示例
    // 在实际的Unix编程中：
    // 1. fork() 创建子进程
    // 2. 在子进程中调用exec()替换程序
    // 3. 父进程可以等待子进程完成

    // Go中使用os/exec包代替fork+exec模式
    cmd := exec.Command("ls", "-la")
    err := cmd.Run()
    ```

11. 文件描述符继承：
    ```go
    // 默认情况下，文件描述符会被继承
    // 如果需要关闭某些文件描述符，需要在exec前处理

    // 例如：关闭不需要的文件
    file.Close()

    // 或设置FD_CLOEXEC标志（在打开文件时）
    ```

12. 信号处理：
    ```go
    // exec后信号处理器会被重置为默认行为
    // 被忽略的信号仍然被忽略
    // 被捕获的信号恢复为默认处理
    ```

13. 工作目录和权限：
    ```go
    // 工作目录保持不变
    os.Chdir("/tmp")
    syscall.Exec(binary, args, env) // 在/tmp目录下执行

    // 用户ID和组ID保持不变（除非是setuid程序）
    ```

14. 调试和监控：
    ```go
    func debugExec(binary string, args []string, env []string) {
        fmt.Printf("Executing: %s\n", binary)
        fmt.Printf("Args: %v\n", args)
        fmt.Printf("PID: %d\n", os.Getpid())

        err := syscall.Exec(binary, args, env)
        if err != nil {
            fmt.Printf("Exec failed: %v\n", err)
        }
    }
    ```

15. 安全考虑：
    ```go
    // 验证可执行文件路径
    func safeExec(command string, args []string) error {
        // 只允许特定目录下的程序
        allowedDirs := []string{"/bin", "/usr/bin", "/usr/local/bin"}

        binary, err := exec.LookPath(command)
        if err != nil {
            return err
        }

        // 验证路径
        allowed := false
        for _, dir := range allowedDirs {
            if strings.HasPrefix(binary, dir) {
                allowed = true
                break
            }
        }

        if !allowed {
            return fmt.Errorf("command not allowed: %s", binary)
        }

        fullArgs := append([]string{command}, args...)
        return syscall.Exec(binary, fullArgs, os.Environ())
    }
    ```

16. 跨平台注意事项：
    ```go
    // syscall.Exec在Windows上不可用
    // Windows使用不同的进程模型

    // 跨平台的替代方案：
    cmd := exec.Command("ls", "-la")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    os.Exit(cmd.ProcessState.ExitCode())
    ```

17. 测试和验证：
    ```bash
    # 编译并运行
    go build -o exec_demo 79_exec\'ing_processes.go
    ./exec_demo

    # 观察进程替换：
    # 1. 程序启动时的PID
    # 2. exec后变成ls进程
    # 3. 程序不会打印"This will never be printed"
    ```

18. 常见错误：
    - 忘记args[0]应该是程序名
    - 路径查找失败（命令不在PATH中）
    - 权限不足
    - 文件不存在或不可执行

19. 最佳实践：
    - 总是检查exec.LookPath的错误
    - 正确设置args[0]
    - 谨慎处理环境变量
    - 考虑安全性（避免执行任意命令）
    - 在生产环境中添加适当的日志记录

这个示例展示了Unix系统编程的核心概念，
是理解进程管理和系统调用的重要基础。

注意：syscall.Exec是Unix特有的，在Windows上不可用。
*/
