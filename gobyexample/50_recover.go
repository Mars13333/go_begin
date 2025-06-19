package main

import "fmt"

// mayPanic函数：主动触发panic用于演示
func mayPanic() {
	panic("a problem")
}

func main() {

	// 使用defer注册一个匿名函数，用于捕获panic
	// 这个defer函数会在main函数返回前执行
	defer func() {
		// 使用Go语言特有的"if with initialization"语法
		// 语法：if initialization; condition { ... }
		// 这里：r := recover() 是初始化表达式，r != nil 是条件判断
		// 这种写法的优势：
		// 1. 变量r的作用域限制在if块内，避免污染外部作用域
		// 2. 代码更简洁，不需要在外部声明变量
		// 3. 只能有一个初始化表达式，多个会报错
		if r := recover(); r != nil {
			// 捕获到panic，打印恢复信息和panic值
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	// 调用会触发panic的函数
	mayPanic()

	// 这行代码只有在没有panic或panic被recover捕获后才会执行
	fmt.Println("After mayPanic()")
}

/*
Go语言recover机制示例 - panic恢复和异常处理

主旨：
1. 演示recover函数的使用和效果
2. 展示defer + recover的组合使用
3. 说明panic和recover的配合机制
4. 理解异常恢复的执行流程
5. 展示Go语言if语句的特殊语法

关键特性：
- recover(): 捕获panic，恢复正常执行
- recover只能在defer函数中调用才有效
- 如果没有panic，recover()返回nil
- recover会停止panic传播，程序继续执行

Go语言if语句特殊语法详解：

1. 基本语法：if initialization; condition { ... }
2. 初始化表达式：在条件判断前执行的语句
3. 条件判断：决定是否执行if块的条件
4. 变量作用域：初始化语句中的变量只在if块内有效

if语句语法示例：
```go
// 传统写法
r := recover()
if r != nil {
    // 使用r
}

// Go特有写法（带初始化）
if r := recover(); r != nil {
    // 使用r，r的作用域限制在if块内
}

// 更多例子：
if file, err := os.Open("file.txt"); err == nil {
    defer file.Close()
    // 使用file
}

if value, ok := someInterface.(string); ok {
    // value是string类型
}

if err := someFunction(); err != nil {
    return err
}
```

执行流程分析：

1. 程序开始执行main函数
2. defer注册匿名函数（用于recover）
3. 调用mayPanic()函数
4. mayPanic()触发panic("a problem")
5. panic开始向上传播
6. defer函数执行，调用recover()
7. recover()捕获panic，返回"a problem"
8. 程序恢复正常执行，打印"Recovered. Error: a problem"
9. main函数继续执行，打印"After mayPanic()"

如果没有recover会发生什么：
- panic会继续向上传播
- 程序会异常终止
- 输出：panic: a problem

recover的重要规则：
1. 只能在defer函数中调用recover()
2. recover()只能捕获同一个goroutine中的panic
3. recover()会停止panic传播，但不会恢复panic发生点之后的代码
4. 一个panic只能被recover()一次

实际应用场景：
- Web服务器中捕获panic，避免整个服务崩溃
- 数据库连接池中的异常恢复
- 第三方库调用的异常处理
- 优雅的错误处理和日志记录

注意事项：
- 不要过度使用recover，应该优先使用正常的错误处理
- recover主要用于处理不可预期的异常情况
- 在生产环境中，应该记录panic信息用于调试
- if的初始化语法是Go语言的特色，让代码更简洁安全
*/
