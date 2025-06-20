package main

import (
	"fmt"
	"net/http"
)

// hello 处理器函数：处理 /hello 路径的请求
// w: ResponseWriter 用于写入HTTP响应
// req: Request 包含客户端请求的所有信息
func hello(w http.ResponseWriter, req *http.Request) {

	// 向响应写入器写入内容，客户端将收到这个响应
	fmt.Fprintf(w, "hello\n")
}

// headers 处理器函数：显示客户端发送的所有HTTP头
func headers(w http.ResponseWriter, req *http.Request) {

	// 遍历请求中的所有HTTP头
	// req.Header 是 map[string][]string 类型
	for name, headers := range req.Header {
		// 每个头可能有多个值，所以需要遍历值的切片
		for _, h := range headers {
			// 将头名称和值写入响应
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// 注册路由处理器
	// 将 /hello 路径绑定到 hello 函数
	http.HandleFunc("/hello", hello)
	// 将 /headers 路径绑定到 headers 函数
	http.HandleFunc("/headers", headers)

	// 启动HTTP服务器
	// 监听8090端口，使用默认的ServeMux路由器（nil表示使用默认）
	// 这个调用会阻塞程序，服务器开始处理请求
	http.ListenAndServe(":8090", nil)
}

/*
HTTP Server HTTP服务器示例主旨：

1. 核心功能：
   演示如何使用Go的net/http包创建HTTP服务器，处理路由和响应客户端请求。

2. 处理器函数签名：
   ```go
   func handler(w http.ResponseWriter, req *http.Request) {
       // 处理逻辑
   }
   ```
   - w: 用于写入HTTP响应的接口
   - req: 包含客户端请求信息的结构体

3. 路由注册：
   - http.HandleFunc()：注册路径和处理器函数的映射
   - 支持路径模式匹配
   - 使用默认的ServeMux路由器

4. 服务器启动：
   - http.ListenAndServe()：启动HTTP服务器
   - 第一个参数是监听地址（:8090表示所有接口的8090端口）
   - 第二个参数是路由器（nil使用默认ServeMux）

5. HTTP请求信息访问：
   ```go
   // 请求方法
   method := req.Method

   // 请求URL
   url := req.URL.Path
   query := req.URL.Query()

   // 请求头
   userAgent := req.Header.Get("User-Agent")
   allHeaders := req.Header

   // 请求体
   body, _ := io.ReadAll(req.Body)
   defer req.Body.Close()

   // 表单数据
   req.ParseForm()
   value := req.FormValue("key")
   ```

6. HTTP响应写入：
   ```go
   // 写入响应体
   fmt.Fprintf(w, "Hello, %s!", name)
   w.Write([]byte("Hello World"))

   // 设置响应头
   w.Header().Set("Content-Type", "application/json")
   w.Header().Add("Set-Cookie", "session=abc123")

   // 设置状态码
   w.WriteHeader(http.StatusNotFound) // 必须在写入内容前调用
   ```

7. 更复杂的路由示例：
   ```go
   // 静态文件服务
   http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

   // 处理所有路径
   http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
       if req.URL.Path != "/" {
           http.NotFound(w, req)
           return
       }
       fmt.Fprintf(w, "Welcome to the home page!")
   })
   ```

8. 自定义ServeMux：
   ```go
   mux := http.NewServeMux()
   mux.HandleFunc("/api/users", usersHandler)
   mux.HandleFunc("/api/posts", postsHandler)

   server := &http.Server{
       Addr:    ":8080",
       Handler: mux,
   }

   log.Fatal(server.ListenAndServe())
   ```

9. 中间件模式：
   ```go
   func logging(next http.HandlerFunc) http.HandlerFunc {
       return func(w http.ResponseWriter, req *http.Request) {
           log.Printf("%s %s", req.Method, req.URL.Path)
           next(w, req)
       }
   }

   http.HandleFunc("/hello", logging(hello))
   ```

10. JSON API示例：
    ```go
    func apiHandler(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        response := map[string]interface{}{
            "message": "Hello API",
            "status":  "success",
        }

        json.NewEncoder(w).Encode(response)
    }
    ```

11. 处理不同HTTP方法：
    ```go
    func userHandler(w http.ResponseWriter, req *http.Request) {
        switch req.Method {
        case http.MethodGet:
            // 获取用户信息
        case http.MethodPost:
            // 创建用户
        case http.MethodPut:
            // 更新用户
        case http.MethodDelete:
            // 删除用户
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }
    ```

12. 错误处理：
    ```go
    func errorHandler(w http.ResponseWriter, req *http.Request) {
        // 返回404错误
        http.NotFound(w, req)

        // 返回自定义错误
        http.Error(w, "Something went wrong", http.StatusInternalServerError)

        // 重定向
        http.Redirect(w, req, "/login", http.StatusFound)
    }
    ```

13. 服务器配置：
    ```go
    server := &http.Server{
        Addr:         ":8080",
        Handler:      mux,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }

    log.Fatal(server.ListenAndServe())
    ```

14. HTTPS支持：
    ```go
    // 使用证书文件启动HTTPS服务器
    log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil))
    ```

15. 优雅关闭：
    ```go
    server := &http.Server{Addr: ":8080"}

    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
    server.Shutdown(ctx)
    ```

16. 测试服务器：
    ```bash
    # 测试hello端点
    curl http://localhost:8090/hello

    # 测试headers端点
    curl -H "Custom-Header: test" http://localhost:8090/headers
    ```

17. 最佳实践：
    - 使用自定义ServeMux而不是默认的
    - 设置合理的超时时间
    - 实现适当的错误处理
    - 使用中间件处理通用逻辑
    - 支持优雅关闭
    - 记录访问日志
    - 实现健康检查端点

18. 常见用例：
    - REST API服务
    - Web应用后端
    - 微服务
    - 静态文件服务
    - 反向代理
    - 健康检查服务

这个示例展示了Go HTTP服务器的基础用法，
是构建Web应用和API服务的起点。
*/
