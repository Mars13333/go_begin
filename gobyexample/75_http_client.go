package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// 发送GET请求到指定URL
	// http.Get是一个便捷方法，相当于http.DefaultClient.Get()
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err) // 处理网络错误或其他请求错误
	}
	// 确保在函数结束时关闭响应体
	// 这是防止资源泄露的重要步骤
	defer resp.Body.Close()

	// 打印HTTP响应状态码和状态文本
	// 例如："200 OK", "404 Not Found" 等
	fmt.Println("Response status:", resp.Status)

	// 使用bufio.Scanner逐行读取响应体内容
	// Scanner提供了便捷的文本读取方法
	scanner := bufio.NewScanner(resp.Body)
	// 读取前5行内容并打印
	// scanner.Scan()返回false表示没有更多内容或遇到错误
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text()) // 获取当前行的文本内容
	}

	// 检查Scanner是否遇到了读取错误
	// 这与到达文件末尾不同，需要单独检查
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

/*
HTTP Client HTTP客户端示例主旨：

1. 核心功能：
   演示如何使用Go的net/http包创建HTTP客户端，发送请求并处理响应。

2. HTTP请求方法：
   - http.Get()：发送GET请求的便捷方法
   - http.Post()：发送POST请求
   - http.PostForm()：发送表单数据
   - http.Head()：发送HEAD请求

3. 响应处理：
   - resp.Status：HTTP状态码和文本
   - resp.StatusCode：数字状态码
   - resp.Header：响应头信息
   - resp.Body：响应体内容（io.ReadCloser）

4. 资源管理：
   - defer resp.Body.Close()：确保响应体被正确关闭
   - 防止内存泄露和连接泄露
   - 即使不读取响应体也要关闭

5. 内容读取方式：
   ```go
   // 方式1：使用bufio.Scanner逐行读取
   scanner := bufio.NewScanner(resp.Body)
   for scanner.Scan() {
       fmt.Println(scanner.Text())
   }

   // 方式2：读取全部内容
   body, err := io.ReadAll(resp.Body)
   if err != nil {
       // 处理错误
   }
   fmt.Println(string(body))

   // 方式3：直接复制到输出
   io.Copy(os.Stdout, resp.Body)
   ```

6. 自定义HTTP客户端：
   ```go
   client := &http.Client{
       Timeout: 30 * time.Second,
       Transport: &http.Transport{
           MaxIdleConns:        100,
           MaxIdleConnsPerHost: 100,
           IdleConnTimeout:     90 * time.Second,
       },
   }

   resp, err := client.Get("https://example.com")
   ```

7. 创建自定义请求：
   ```go
   req, err := http.NewRequest("GET", "https://example.com", nil)
   if err != nil {
       // 处理错误
   }

   // 设置请求头
   req.Header.Set("User-Agent", "MyApp/1.0")
   req.Header.Set("Authorization", "Bearer token")

   client := &http.Client{}
   resp, err := client.Do(req)
   ```

8. POST请求示例：
   ```go
   // 发送JSON数据
   jsonData := `{"name":"John","age":30}`
   resp, err := http.Post("https://api.example.com/users",
       "application/json",
       strings.NewReader(jsonData))

   // 发送表单数据
   formData := url.Values{}
   formData.Set("username", "john")
   formData.Set("password", "secret")
   resp, err := http.PostForm("https://example.com/login", formData)
   ```

9. 错误处理：
   ```go
   resp, err := http.Get("https://example.com")
   if err != nil {
       // 网络错误、DNS错误、超时等
       log.Fatal(err)
   }
   defer resp.Body.Close()

   // 检查HTTP状态码
   if resp.StatusCode != http.StatusOK {
       log.Fatalf("HTTP error: %s", resp.Status)
   }
   ```

10. 超时设置：
    ```go
    client := &http.Client{
        Timeout: 10 * time.Second, // 总超时时间
    }

    // 或者使用context控制超时
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    req, _ := http.NewRequestWithContext(ctx, "GET", "https://example.com", nil)
    resp, err := client.Do(req)
    ```

11. 请求头操作：
    ```go
    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
    req.Header.Del("User-Agent")

    // 获取响应头
    contentType := resp.Header.Get("Content-Type")
    allHeaders := resp.Header["Set-Cookie"] // 获取多个值
    ```

12. Cookie处理：
    ```go
    jar, _ := cookiejar.New(nil)
    client := &http.Client{
        Jar: jar, // 自动处理cookie
    }

    // 手动设置cookie
    cookie := &http.Cookie{
        Name:  "session",
        Value: "abc123",
    }
    req.AddCookie(cookie)
    ```

13. 文件上传：
    ```go
    file, _ := os.Open("file.txt")
    defer file.Close()

    resp, err := http.Post("https://example.com/upload",
        "text/plain",
        file)
    ```

14. 下载文件：
    ```go
    resp, err := http.Get("https://example.com/file.zip")
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    out, err := os.Create("file.zip")
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, resp.Body)
    ```

15. 最佳实践：
    - 总是关闭响应体
    - 设置合理的超时时间
    - 检查HTTP状态码
    - 使用自定义客户端而不是默认客户端
    - 处理重定向和认证
    - 考虑连接池和重用
    - 实现重试机制

16. 常见用例：
    - API调用和集成
    - 网页抓取和爬虫
    - 文件下载和上传
    - 健康检查和监控
    - 微服务间通信

这个示例展示了Go HTTP客户端的基础用法，
是构建网络应用和API客户端的基础。
*/
