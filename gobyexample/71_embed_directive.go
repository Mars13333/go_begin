package main

import (
	"embed"
)

// //go:embed 指令将文件内容嵌入到字符串变量中
// 编译时文件内容会被直接包含在二进制文件中
//
//go:embed folder/single_file.txt
var fileString string

// 同一个文件可以嵌入到不同类型的变量中
// 这里嵌入为字节切片
//
//go:embed folder/single_file.txt
var fileByte []byte

// embed.FS 类型可以嵌入整个文件系统
// 支持通配符模式匹配多个文件
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// 直接使用嵌入的字符串
	// 内容在编译时就确定，运行时无需文件系统访问
	print(fileString)

	// 字节切片需要转换为字符串才能打印
	print(string(fileByte))

	// 使用 embed.FS 读取嵌入的文件
	// ReadFile 方法返回文件内容和错误
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	// 读取另一个嵌入的文件
	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}

/*
Embed Directive 嵌入指令示例主旨：

1. 核心功能：
   Go 1.16+ 的 //go:embed 指令允许在编译时将文件内容
   直接嵌入到二进制文件中，无需运行时文件系统访问。

2. 嵌入类型：
   - string：文件内容作为字符串嵌入
   - []byte：文件内容作为字节切片嵌入
   - embed.FS：文件系统接口，支持多文件和目录结构

3. 指令语法：
   - //go:embed path/to/file：嵌入单个文件
   - //go:embed path/*.ext：使用通配符嵌入多个文件
   - //go:embed path/to/dir：嵌入整个目录
   - 多个 //go:embed 指令可以组合使用

4. 优势：
   - 单一二进制：所有资源打包在一个可执行文件中
   - 部署简化：无需单独分发资源文件
   - 性能提升：避免运行时文件系统I/O
   - 安全性：资源无法被外部修改

5. 使用场景：
   - 静态网站资源：HTML、CSS、JS文件
   - 配置文件：默认配置模板
   - 数据文件：初始化数据、种子数据
   - 模板文件：文本模板、邮件模板
   - 证书文件：SSL证书、密钥文件

6. 限制和注意事项：
   - 只能嵌入项目内的文件（不能跨模块边界）
   - 文件路径必须是相对路径
   - 嵌入的文件在编译时确定，运行时不可变
   - 会增加二进制文件大小

7. embed.FS 接口：
   - ReadFile(name string) ([]byte, error)
   - ReadDir(name string) ([]DirEntry, error)
   - Open(name string) (File, error)
   - 实现了 fs.FS 接口，可用于 http.FileServer

8. 最佳实践：
   - 对于小文件使用 string 或 []byte
   - 对于多文件或目录结构使用 embed.FS
   - 合理组织嵌入文件的目录结构
   - 考虑二进制文件大小的影响

9. 编译时行为：
   - go build 时自动处理 //go:embed 指令
   - 文件不存在会导致编译错误
   - 支持 go mod 和 go workspace

这个示例展示了 Go 中资源嵌入的现代方法，
是构建自包含应用程序的重要工具。
*/
