package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// os.Setenv 设置环境变量
	// 第一个参数是变量名，第二个参数是变量值
	os.Setenv("FOO", "1")

	// os.Getenv 获取环境变量的值
	// 如果变量不存在，返回空字符串
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR")) // BAR 未设置，将返回空字符串

	fmt.Println()

	// os.Environ() 返回所有环境变量的字符串切片
	// 每个字符串格式为 "key=value"
	for _, e := range os.Environ() {
		// strings.SplitN 将字符串按 "=" 分割成最多 2 部分
		// 这样可以正确处理值中包含 "=" 的情况
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0]) // 只打印环境变量名（key）
	}
}

/*
Environment Variables 环境变量示例主旨：

1. 核心功能：
   演示如何在 Go 程序中读取、设置和遍历系统环境变量。

2. 主要操作：
   - os.Setenv()：设置环境变量
   - os.Getenv()：获取环境变量值
   - os.Environ()：获取所有环境变量
   - os.Unsetenv()：删除环境变量（未在示例中使用）

3. 环境变量特点：
   - 键值对形式：NAME=VALUE
   - 字符串类型：所有值都是字符串
   - 进程继承：子进程继承父进程的环境变量
   - 系统级别：在操作系统级别管理

4. 常见用途：
   - 配置信息：数据库连接字符串、API密钥
   - 运行环境：开发/测试/生产环境标识
   - 路径设置：PATH、HOME、GOPATH等
   - 功能开关：启用/禁用特定功能
   - 敏感信息：避免硬编码密码和密钥

5. 最佳实践：
   - 使用大写字母命名环境变量
   - 提供默认值处理未设置的变量
   - 验证环境变量的格式和有效性
   - 文档化所需的环境变量
   - 敏感信息通过环境变量而非代码传递

6. 获取环境变量的模式：
   ```go
   // 基本获取
   value := os.Getenv("MY_VAR")

   // 带默认值
   value := os.Getenv("MY_VAR")
   if value == "" {
       value = "default_value"
   }

   // 检查是否存在
   value, exists := os.LookupEnv("MY_VAR")
   if !exists {
       // 处理变量不存在的情况
   }
   ```

7. 类型转换：
   环境变量都是字符串，需要转换为其他类型：
   ```go
   import "strconv"

   port := os.Getenv("PORT")
   if port == "" {
       port = "8080"
   }
   portNum, err := strconv.Atoi(port)
   ```

8. 安全考虑：
   - 避免在日志中打印敏感环境变量
   - 使用 os.Clearenv() 清理测试环境
   - 限制环境变量的访问权限
   - 验证环境变量值的安全性

9. 十二要素应用：
   环境变量是十二要素应用方法论中配置管理的推荐方式：
   - 严格分离配置和代码
   - 不同环境使用不同配置
   - 避免配置文件的复杂性

10. 实际应用示例：
    ```go
    dbURL := os.Getenv("DATABASE_URL")
    apiKey := os.Getenv("API_KEY")
    debug := os.Getenv("DEBUG") == "true"
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    ```

这个示例展示了 Go 中环境变量操作的基础模式，
是构建可配置和可移植应用程序的重要工具。
*/
