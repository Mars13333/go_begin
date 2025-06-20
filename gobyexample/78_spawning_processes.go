package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// 创建一个执行"date"命令的Command对象
	// exec.Command返回一个*exec.Cmd，表示要执行的外部命令
	dateCmd := exec.Command("date")

	// Output()方法执行命令并返回标准输出
	// 这是执行简单命令并获取输出的最便捷方法
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// 演示错误处理：执行一个会失败的命令
	// "date -x" 是一个无效的参数组合
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		// 使用类型断言处理不同类型的错误
		switch e := err.(type) {
		case *exec.Error:
			// exec.Error：命令无法启动（如命令不存在）
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			// exec.ExitError：命令执行但以非零状态退出
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// 演示与命令的标准输入/输出交互
	// 创建grep命令，用于搜索包含"hello"的行
	grepCmd := exec.Command("grep", "hello")

	// 获取命令的标准输入管道，用于向命令发送数据
	grepIn, _ := grepCmd.StdinPipe()
	// 获取命令的标准输出管道，用于读取命令输出
	grepOut, _ := grepCmd.StdoutPipe()
	// 启动命令但不等待其完成
	grepCmd.Start()
	// 向命令的标准输入写入数据
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	// 关闭标准输入，告诉命令没有更多输入
	grepIn.Close()
	// 读取命令的所有输出
	grepBytes, _ := io.ReadAll(grepOut)
	// 等待命令完成
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 演示执行shell命令
	// 使用bash -c可以执行复杂的shell命令字符串
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

/*
Spawning Processes 进程生成示例主旨：

1. 核心功能：
   演示如何在Go程序中执行外部命令和进程，包括命令执行、输入输出处理和错误处理。

2. exec.Command基础用法：
   ```go
   // 创建命令
   cmd := exec.Command("command", "arg1", "arg2")

   // 执行并获取输出
   output, err := cmd.Output()

   // 执行并获取输出和错误输出
   output, err := cmd.CombinedOutput()

   // 简单执行（不获取输出）
   err := cmd.Run()
   ```

3. 命令执行的三种方式：
   - Run()：执行命令并等待完成，不返回输出
   - Output()：执行命令并返回标准输出
   - CombinedOutput()：执行命令并返回标准输出和错误输出的组合

4. 错误类型处理：
   - exec.Error：命令无法启动（命令不存在、权限问题等）
   - exec.ExitError：命令启动但以非零状态退出
   - 其他系统错误

5. 管道操作：
   ```go
   cmd := exec.Command("grep", "pattern")

   // 获取标准输入管道
   stdin, err := cmd.StdinPipe()

   // 获取标准输出管道
   stdout, err := cmd.StdoutPipe()

   // 获取标准错误管道
   stderr, err := cmd.StderrPipe()

   // 启动命令
   cmd.Start()

   // 与命令交互...

   // 等待命令完成
   cmd.Wait()
   ```

6. 设置命令环境：
   ```go
   cmd := exec.Command("env")
   cmd.Env = append(os.Environ(), "FOO=bar")
   cmd.Dir = "/tmp" // 设置工作目录
   output, err := cmd.Output()
   ```

7. 超时控制：
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()

   cmd := exec.CommandContext(ctx, "sleep", "10")
   err := cmd.Run()
   if err != nil {
       // 处理超时或其他错误
   }
   ```

8. 实时输出处理：
   ```go
   cmd := exec.Command("ping", "google.com")
   stdout, _ := cmd.StdoutPipe()
   cmd.Start()

   scanner := bufio.NewScanner(stdout)
   for scanner.Scan() {
       fmt.Println("Output:", scanner.Text())
   }

   cmd.Wait()
   ```

9. 命令链和管道：
   ```go
   // 模拟: ls | grep .go
   lsCmd := exec.Command("ls")
   grepCmd := exec.Command("grep", ".go")

   // 将ls的输出连接到grep的输入
   grepCmd.Stdin, _ = lsCmd.StdoutPipe()
   grepOut, _ := grepCmd.StdoutPipe()

   grepCmd.Start()
   lsCmd.Run()

   output, _ := io.ReadAll(grepOut)
   grepCmd.Wait()
   ```

10. 异步执行：
    ```go
    cmd := exec.Command("long-running-command")

    // 异步启动
    err := cmd.Start()
    if err != nil {
        log.Fatal(err)
    }

    // 在后台运行，继续其他工作...

    // 稍后等待完成
    err = cmd.Wait()
    ```

11. 进程信息获取：
    ```go
    cmd := exec.Command("sleep", "60")
    cmd.Start()

    // 获取进程ID
    pid := cmd.Process.Pid
    fmt.Printf("Process ID: %d\n", pid)

    // 获取进程状态
    state := cmd.ProcessState
    if state != nil && state.Exited() {
        fmt.Printf("Exit code: %d\n", state.ExitCode())
    }
    ```

12. 信号处理：
    ```go
    cmd := exec.Command("sleep", "60")
    cmd.Start()

    // 发送中断信号
    cmd.Process.Signal(os.Interrupt)

    // 或强制终止
    cmd.Process.Kill()
    ```

13. 标准流重定向：
    ```go
    cmd := exec.Command("command")

    // 重定向到文件
    outFile, _ := os.Create("output.txt")
    defer outFile.Close()
    cmd.Stdout = outFile

    // 重定向到当前进程的标准输出
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    cmd.Run()
    ```

14. 跨平台命令执行：
    ```go
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/C", "dir")
    } else {
        cmd = exec.Command("ls", "-la")
    }

    output, err := cmd.Output()
    ```

15. 安全考虑：
    ```go
    // 避免命令注入，不要这样做：
    // cmd := exec.Command("sh", "-c", userInput)

    // 正确的做法：验证和清理输入
    func safeCommand(userFile string) error {
        // 验证文件名
        if !isValidFilename(userFile) {
            return errors.New("invalid filename")
        }

        cmd := exec.Command("cat", userFile)
        return cmd.Run()
    }
    ```

16. 性能优化：
    ```go
    // 重用Command对象（注意：只能执行一次）
    cmd := exec.Command("echo", "hello")

    // 设置合适的缓冲区大小
    var buf bytes.Buffer
    cmd.Stdout = &buf

    // 使用context控制超时
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    cmd = exec.CommandContext(ctx, "command")
    ```

17. 错误处理最佳实践：
    ```go
    cmd := exec.Command("command")
    output, err := cmd.Output()

    if err != nil {
        if exitError, ok := err.(*exec.ExitError); ok {
            // 获取stderr输出
            stderr := exitError.Stderr
            log.Printf("Command failed with exit code %d: %s",
                exitError.ExitCode(), stderr)
        } else {
            log.Printf("Failed to execute command: %v", err)
        }
        return err
    }
    ```

18. 常见用例：
    - 系统管理和自动化脚本
    - 构建和部署工具
    - 数据处理管道
    - 外部工具集成
    - 系统监控和诊断
    - 文件处理和转换

19. 注意事项：
    - 始终处理错误和异常情况
    - 注意命令注入安全问题
    - 合理设置超时时间
    - 正确处理标准输入输出
    - 考虑跨平台兼容性
    - 避免阻塞主程序执行

这个示例展示了Go中执行外部命令的基本模式，
是系统编程和工具开发的重要技能。
*/
