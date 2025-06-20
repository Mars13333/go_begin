package main

import (
	"fmt"
	"net/http"
	"time"
)

// hello 处理器函数：演示Context的使用
func hello(w http.ResponseWriter, req *http.Request) {

	// 从HTTP请求中获取Context
	// 每个HTTP请求都有一个关联的Context，用于处理取消和超时
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// 使用select语句同时监听多个channel
	select {
	// 模拟长时间运行的操作（10秒）
	case <-time.After(10 * time.Second):
		// 如果10秒内没有被取消，正常返回响应
		fmt.Fprintf(w, "hello\n")
	// 监听Context的Done channel
	case <-ctx.Done():
		// 当客户端断开连接或请求被取消时，ctx.Done()会被触发

		// 获取Context取消的原因
		err := ctx.Err()
		fmt.Println("server:", err)
		// 返回500内部服务器错误
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// 注册路由处理器
	http.HandleFunc("/hello", hello)
	// 启动HTTP服务器，监听8090端口
	http.ListenAndServe(":8090", nil)
}

/*
Context 上下文示例主旨：

1. 核心功能：
   演示如何在HTTP服务器中使用Context处理请求取消、超时和传递请求范围的值。

2. Context的主要用途：
   - 取消信号传播：当客户端断开连接时通知服务器停止处理
   - 超时控制：设置操作的最大执行时间
   - 请求范围值传递：在调用链中传递请求ID、用户信息等
   - 截止时间管理：设置绝对的截止时间

3. HTTP请求中的Context：
   - req.Context()：获取与HTTP请求关联的Context
   - 当客户端断开连接时，Context会自动被取消
   - 服务器可以通过监听ctx.Done()来响应取消事件

4. Context接口方法：
   ```go
   type Context interface {
       Deadline() (deadline time.Time, ok bool)  // 返回截止时间
       Done() <-chan struct{}                    // 返回取消信号channel
       Err() error                               // 返回取消原因
       Value(key interface{}) interface{}        // 获取存储的值
   }
   ```

5. Context取消的原因：
   - context.Canceled：Context被显式取消
   - context.DeadlineExceeded：超过了截止时间
   - 其他自定义错误

6. 创建Context的方法：
   ```go
   // 背景Context（根Context）
   ctx := context.Background()
   ctx := context.TODO() // 当不确定使用哪个Context时

   // 带取消功能的Context
   ctx, cancel := context.WithCancel(parent)
   defer cancel() // 确保释放资源

   // 带超时的Context
   ctx, cancel := context.WithTimeout(parent, 5*time.Second)
   defer cancel()

   // 带截止时间的Context
   deadline := time.Now().Add(10 * time.Second)
   ctx, cancel := context.WithDeadline(parent, deadline)
   defer cancel()

   // 带值的Context
   ctx := context.WithValue(parent, "userID", 12345)
   ```

7. 在函数中使用Context：
   ```go
   func doWork(ctx context.Context) error {
       select {
       case <-time.After(5 * time.Second):
           // 工作完成
           return nil
       case <-ctx.Done():
           // 被取消
           return ctx.Err()
       }
   }
   ```

8. 数据库操作中的Context：
   ```go
   func queryDatabase(ctx context.Context, query string) error {
       // 使用Context控制数据库查询超时
       rows, err := db.QueryContext(ctx, query)
       if err != nil {
           return err
       }
       defer rows.Close()

       // 处理结果...
       return nil
   }
   ```

9. HTTP客户端中的Context：
   ```go
   func makeRequest(ctx context.Context, url string) error {
       req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
       if err != nil {
           return err
       }

       client := &http.Client{}
       resp, err := client.Do(req)
       if err != nil {
           return err
       }
       defer resp.Body.Close()

       return nil
   }
   ```

10. 并发操作中的Context：
    ```go
    func processItems(ctx context.Context, items []Item) error {
        for _, item := range items {
            select {
            case <-ctx.Done():
                return ctx.Err() // 提前退出
            default:
                if err := processItem(ctx, item); err != nil {
                    return err
                }
            }
        }
        return nil
    }
    ```

11. Context值的传递：
    ```go
    type contextKey string

    const userIDKey contextKey = "userID"

    // 设置值
    ctx := context.WithValue(context.Background(), userIDKey, 12345)

    // 获取值
    if userID, ok := ctx.Value(userIDKey).(int); ok {
        fmt.Printf("User ID: %d\n", userID)
    }
    ```

12. 中间件中的Context使用：
    ```go
    func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, req *http.Request) {
            // 从请求中提取用户信息
            userID := extractUserID(req)

            // 将用户ID添加到Context
            ctx := context.WithValue(req.Context(), "userID", userID)
            req = req.WithContext(ctx)

            next(w, req)
        }
    }
    ```

13. 优雅关闭中的Context：
    ```go
    func gracefulShutdown() {
        ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
        defer cancel()

        if err := server.Shutdown(ctx); err != nil {
            log.Fatal("Server forced to shutdown:", err)
        }
    }
    ```

14. Context最佳实践：
    - 总是将Context作为函数的第一个参数
    - 不要将Context存储在结构体中
    - 不要传递nil Context，使用context.TODO()
    - Context是不可变的，WithValue等方法返回新的Context
    - 及时调用cancel函数释放资源
    - 只在请求范围内传递值，不要传递可选参数

15. 测试示例：
    ```bash
    # 启动服务器
    go run context.go &

    # 正常请求（等待10秒）
    curl http://localhost:8090/hello

    # 取消请求（Ctrl+C中断curl）
    curl http://localhost:8090/hello
    # 然后按Ctrl+C取消请求
    ```

16. 常见错误：
    ```go
    // 错误：不要这样做
    type Handler struct {
        ctx context.Context // 不要在结构体中存储Context
    }

    // 正确：将Context作为参数传递
    func (h *Handler) Process(ctx context.Context) error {
        // 处理逻辑
    }
    ```

17. Context链：
    ```go
    // Context可以形成链式结构
    rootCtx := context.Background()
    timeoutCtx, cancel1 := context.WithTimeout(rootCtx, 10*time.Second)
    defer cancel1()

    cancelCtx, cancel2 := context.WithCancel(timeoutCtx)
    defer cancel2()

    // cancelCtx会在任一父Context取消时被取消
    ```

18. 实际应用场景：
    - HTTP请求处理和取消
    - 数据库查询超时控制
    - gRPC调用管理
    - 批处理任务的中断
    - 微服务调用链追踪
    - 资源清理和优雅关闭

这个示例展示了Context在HTTP服务器中的基本用法，
是Go并发编程和请求处理的重要概念。

测试方法：
1. 运行服务器：go run context.go &
2. 发送请求：curl localhost:8090/hello
3. 在10秒内按Ctrl+C取消请求，观察服务器日志输出
*/
