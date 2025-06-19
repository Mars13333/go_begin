package main

import (
	"bufio" // 导入缓冲读取包
	"fmt"
	"io" // 导入IO接口包
	"os" // 导入操作系统包
)

// 错误检查辅助函数，简化错误处理
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 方法1：一次性读取整个文件到内存
	// os.ReadFile：读取整个文件内容，返回字节切片
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// 方法2：打开文件进行流式读取
	// os.Open：打开文件，返回文件句柄
	f, err := os.Open("/tmp/dat")
	check(err)

	// 读取前5个字节
	// f.Read：从文件当前位置读取指定字节数
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// 文件指针定位：从文件开头偏移6个字节
	// f.Seek：移动文件指针，io.SeekStart表示从文件开头计算
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	// 读取2个字节
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// 从当前位置向前偏移2个字节
	// io.SeekCurrent：从当前位置计算偏移
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// 从文件末尾向前偏移4个字节
	// io.SeekEnd：从文件末尾计算偏移
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	// 重新定位到第6个字节位置
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	// 使用io.ReadAtLeast确保至少读取2个字节
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 重置文件指针到文件开头
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// 使用缓冲读取器
	// bufio.NewReader：创建带缓冲的读取器，提高读取效率
	r4 := bufio.NewReader(f)
	// Peek：预览接下来的5个字节，不移动文件指针
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 关闭文件，释放系统资源
	f.Close()
}

/*
Go语言文件读取示例 - 多种文件读取方法

主旨：
1. 演示Go语言中多种文件读取方法
2. 展示文件指针定位和随机访问
3. 理解缓冲读取和流式读取的区别
4. 学习文件操作的最佳实践

关键特性：
- os.ReadFile：一次性读取整个文件
- os.Open + f.Read：流式读取文件
- f.Seek：文件指针定位
- bufio.NewReader：缓冲读取提高效率

文件读取方法对比：

1. 一次性读取（os.ReadFile）：
   - 优点：简单、快速
   - 缺点：占用内存多，不适合大文件
   - 适用：小文件、配置文件

2. 流式读取（os.Open + f.Read）：
   - 优点：内存占用少，支持大文件
   - 缺点：需要手动管理文件指针
   - 适用：大文件、网络流

3. 缓冲读取（bufio.NewReader）：
   - 优点：读取效率高，支持Peek操作
   - 缺点：需要额外内存
   - 适用：频繁读取、需要预览

文件指针定位（Seek）：
- io.SeekStart：从文件开头计算偏移
- io.SeekCurrent：从当前位置计算偏移
- io.SeekEnd：从文件末尾计算偏移
- 正数：向前偏移，负数：向后偏移

输出示例：
hello
go
5 bytes: hello
2 bytes @ 6: go
2 bytes @ 6: go
5 bytes: hello

读取函数详解：
- f.Read(b)：读取到字节切片，返回读取字节数
- f.Seek(offset, whence)：移动文件指针
- io.ReadAtLeast(f, b, min)：至少读取min字节
- r.Peek(n)：预览n个字节，不移动指针

错误处理模式：
```go
// 检查错误并处理
if err != nil {
    // 处理错误
    return err
}

// 使用辅助函数简化
func check(e error) {
    if e != nil {
        panic(e)
    }
}
```

实际应用场景：
- 配置文件读取
- 日志文件分析
- 大文件处理
- 网络数据流处理
- 数据库备份文件
- 媒体文件处理

性能优化建议：
- 小文件使用os.ReadFile
- 大文件使用流式读取
- 频繁读取使用缓冲读取
- 避免频繁的Seek操作
- 及时关闭文件句柄

文件操作最佳实践：
```go
// 使用defer确保文件关闭
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close()

// 使用缓冲读取器
reader := bufio.NewReader(file)
line, err := reader.ReadString('\n')

// 错误处理
if err != nil && err != io.EOF {
    return err
}
```

安全注意事项：
- 检查文件路径，避免路径遍历攻击
- 验证文件大小，防止内存溢出
- 处理文件权限错误
- 注意文件锁定问题

相关包和接口：
- os：文件系统操作
- io：IO接口定义
- bufio：缓冲IO操作
- ioutil：便捷IO函数（已废弃）

高级读取技巧：
```go
// 按行读取
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    // 处理每一行
}

// 按块读取
buffer := make([]byte, 1024)
for {
    n, err := file.Read(buffer)
    if n == 0 || err == io.EOF {
        break
    }
    // 处理读取的数据
}
```
*/
