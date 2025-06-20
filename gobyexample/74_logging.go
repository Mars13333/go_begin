package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog" // Go 1.21+ 的结构化日志包
)

func main() {

	// 使用标准日志记录器输出基本日志
	log.Println("standard logger")

	// 设置日志标志：标准标志 + 微秒级时间戳
	// LstdFlags 包含日期和时间，Lmicroseconds 添加微秒精度
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// 设置日志标志：标准标志 + 文件名和行号
	// Lshortfile 显示调用日志的文件名和行号
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// 创建自定义日志记录器
	// 参数：输出目标、前缀、标志
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// 动态修改日志前缀
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// 创建写入到缓冲区的日志记录器
	// 这对于测试或收集日志内容很有用
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("hello")

	// 从缓冲区读取并打印日志内容
	fmt.Print("from buflog:", buf.String())

	// 使用 slog（结构化日志）- Go 1.21+ 特性
	// 创建 JSON 格式的日志处理器，输出到标准错误
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// 结构化日志：支持键值对参数
	// 这样的日志更容易被日志分析工具处理
	myslog.Info("hello again", "key", "val", "age", 25)
}

/*
Logging 日志记录示例主旨：

1. 核心功能：
   演示 Go 语言中的两种日志记录方式：传统的 log 包和现代的 slog 包。

2. 传统日志 (log 包)：
   - log.Println()：基本日志输出
   - log.SetFlags()：设置日志格式标志
   - log.New()：创建自定义日志记录器
   - 支持不同的输出目标和格式

3. 日志标志选项：
   - log.LstdFlags：标准标志（日期 + 时间）
   - log.Lmicroseconds：微秒级时间戳
   - log.Lshortfile：文件名和行号
   - log.Llongfile：完整文件路径和行号
   - log.LUTC：使用 UTC 时间

4. 自定义日志记录器：
   - 可以指定输出目标（文件、缓冲区等）
   - 可以设置自定义前缀
   - 可以配置不同的格式标志
   - 支持动态修改配置

5. 结构化日志 (slog 包 - Go 1.21+)：
   - 支持结构化数据（键值对）
   - 多种输出格式（JSON、文本等）
   - 更好的性能和可扩展性
   - 支持日志级别（Debug、Info、Warn、Error）

6. 日志级别和方法：
   ```go
   slog.Debug("debug message")
   slog.Info("info message")
   slog.Warn("warning message")
   slog.Error("error message")
   ```

7. 结构化日志的优势：
   - 机器可读：JSON 格式便于解析
   - 上下文丰富：键值对提供更多信息
   - 查询友好：便于日志搜索和分析
   - 性能更好：减少字符串拼接

8. 实际应用模式：
   ```go
   // 应用启动日志
   slog.Info("server starting", "port", 8080, "env", "production")

   // 错误日志
   slog.Error("database connection failed",
       "error", err.Error(),
       "host", dbHost,
       "retries", retryCount)

   // 业务日志
   slog.Info("user login",
       "user_id", userID,
       "ip", clientIP,
       "duration_ms", duration.Milliseconds())
   ```

9. 日志最佳实践：
   - 使用适当的日志级别
   - 包含足够的上下文信息
   - 避免记录敏感信息（密码、令牌等）
   - 使用结构化格式便于分析
   - 考虑日志的性能影响

10. 配置选项：
    ```go
    // 文本格式处理器
    textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelDebug,
        AddSource: true,
    })

    // JSON 格式处理器
    jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
        ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
            // 自定义属性处理
            return a
        },
    })
    ```

11. 日志轮转和管理：
    - 使用第三方库（如 lumberjack）进行日志轮转
    - 配置日志文件大小和保留策略
    - 考虑日志的存储和备份

12. 生产环境考虑：
    - 合理的日志级别设置
    - 避免过度日志记录影响性能
    - 集中化日志收集和分析
    - 监控和告警配置

这个示例展示了 Go 中日志记录的演进和最佳实践，
从简单的文本日志到现代的结构化日志系统。
*/
