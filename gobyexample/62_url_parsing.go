package main

import (
	"fmt"
	"net"     // 导入网络包，用于主机和端口分离
	"net/url" // 导入URL解析包
)

func main() {

	// 定义一个完整的URL字符串，包含所有组件
	// 格式：scheme://user:password@host:port/path?query#fragment
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析URL字符串为url.URL结构体
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// 获取URL的协议部分（scheme）
	fmt.Println(u.Scheme)

	// 获取用户信息部分（user:password）
	fmt.Println(u.User)
	// 获取用户名
	fmt.Println(u.User.Username())
	// 获取密码
	p, _ := u.User.Password()
	fmt.Println(p)

	// 获取主机部分（host:port）
	fmt.Println(u.Host)
	// 分离主机名和端口号
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// 获取路径部分
	fmt.Println(u.Path)
	// 获取片段部分（#后面的内容）
	fmt.Println(u.Fragment)

	// 获取原始查询字符串
	fmt.Println(u.RawQuery)
	// 解析查询参数为map
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	// 获取特定查询参数的值
	fmt.Println(m["k"][0])
}

/*
Go语言URL解析示例 - net/url包使用

主旨：
1. 演示net/url包中URL解析和操作的方法
2. 展示URL各个组成部分的提取
3. 理解URL结构体和查询参数处理
4. 学习网络地址的解析和操作

关键特性：
- net/url：Go标准库的URL处理包
- 支持完整的URL解析和构建
- 提供查询参数的解析和编码
- 支持用户认证信息的处理

URL结构说明：
URL格式：scheme://user:password@host:port/path?query#fragment

各部分含义：
- scheme：协议（如http、https、postgres、ftp等）
- user:password：用户认证信息
- host:port：主机名和端口号
- path：路径
- query：查询参数（key=value格式）
- fragment：片段标识符（#后面的内容）

URL解析函数：
- url.Parse(s)：解析URL字符串为url.URL结构体
- url.ParseRequestURI(s)：解析请求URI
- url.ParseQuery(s)：解析查询字符串为map

url.URL结构体字段：
- Scheme：协议
- User：用户信息（包含用户名和密码）
- Host：主机（包含端口）
- Path：路径
- RawQuery：原始查询字符串
- Fragment：片段
- Query：解析后的查询参数map

输出示例：
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v

查询参数处理：
```go
// 解析查询字符串
query := "name=john&age=25&city=beijing"
params, _ := url.ParseQuery(query)
fmt.Println(params["name"]) // [john]
fmt.Println(params["age"])  // [25]

// 构建查询字符串
values := url.Values{}
values.Set("name", "john")
values.Set("age", "25")
queryString := values.Encode() // "age=25&name=john"
```

URL构建：
```go
// 构建URL
u := &url.URL{
    Scheme:   "https",
    Host:     "example.com",
    Path:     "/api/users",
    RawQuery: "id=123",
}
fmt.Println(u.String()) // "https://example.com/api/users?id=123"
```

实际应用场景：
- Web应用URL处理
- API接口参数解析
- 数据库连接字符串解析
- 代理服务器配置
- 重定向URL处理
- 爬虫URL解析

URL编码和解码：
```go
// URL编码
encoded := url.QueryEscape("hello world") // "hello+world"
decoded, _ := url.QueryUnescape(encoded)  // "hello world"

// 路径编码
pathEncoded := url.PathEscape("/path with spaces")
pathDecoded, _ := url.PathUnescape(pathEncoded)
```

安全注意事项：
- 验证URL来源，避免恶意URL
- 处理URL编码，防止注入攻击
- 检查协议类型，确保安全
- 验证主机名，避免DNS攻击

错误处理：
- url.Parse()可能返回解析错误
- 检查URL格式是否正确
- 处理无效的查询参数
- 验证端口号范围

性能考虑：
- URL解析是相对轻量级的操作
- 频繁解析考虑缓存结果
- 大量URL处理考虑并发
- 查询参数解析注意内存使用

相关包：
- net：网络地址处理
- net/http：HTTP客户端和服务器
- crypto/tls：TLS/SSL支持
- context：请求上下文管理
*/
