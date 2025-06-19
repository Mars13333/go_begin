package main

import (
	"fmt"
	"os"
)

func main() {

	// 创建文件，如果文件已存在会被覆盖
	f := createFile("./tmp/defer2.txt")
	// 使用defer确保文件在函数结束时被关闭
	// defer语句会在函数返回前执行，无论函数是正常返回还是panic
	defer closeFile(f)
	// 向文件写入数据
	writeFile(f)
}

// createFile函数：创建或覆盖文件
// 注意：os.Create的行为是"创建或覆盖"，不是"仅创建"
func createFile(p string) *os.File {
	fmt.Println("creating")
	// os.Create会：
	// 1. 如果文件不存在：创建新文件
	// 2. 如果文件已存在：截断（清空）现有文件
	// 3. 返回文件句柄和可能的错误
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

// writeFile函数：向文件写入数据
func writeFile(f *os.File) {
	fmt.Println("writing")
	// 使用fmt.Fprintln向文件写入字符串"data"并换行
	fmt.Fprintln(f, "data")
}

// closeFile函数：关闭文件并处理可能的错误
func closeFile(f *os.File) {
	fmt.Println("closing")
	// 关闭文件，释放系统资源
	err := f.Close()

	if err != nil {
		panic(err)
	}
}

/*
Go语言defer机制示例 - 资源管理和延迟执行

主旨：
1. 演示defer关键字的使用和效果
2. 展示defer在资源管理中的重要作用
3. 说明defer的执行顺序（后进先出LIFO）
4. 理解文件操作的完整流程

关键特性：
- defer：延迟执行，在函数返回前执行
- defer的执行顺序：后进先出（LIFO）
- defer常用于资源清理：关闭文件、解锁、恢复panic等
- defer语句在函数开始时求值，但延迟到函数结束时执行

关于createFile函数的疑问解答：

os.Create的行为确实容易让人困惑：
1. 函数名是"Create"，但实际行为是"创建或覆盖"
2. 如果文件不存在：创建新文件
3. 如果文件已存在：截断（清空）现有文件内容
4. 这更像是"确保文件存在且为空"的操作

如果想要"仅创建新文件"的行为，应该使用：
- os.OpenFile(path, os.O_CREATE|os.O_EXCL, 0644) // 仅创建，文件存在则失败
- 或者先检查文件是否存在：os.Stat(path)

执行流程：
1. createFile: 创建/覆盖文件
2. defer closeFile: 注册关闭操作
3. writeFile: 写入数据
4. main函数结束: 执行defer closeFile

输出顺序：
creating
writing
closing

defer的优势：
- 确保资源被正确释放
- 即使发生panic也能执行清理操作
- 代码更清晰，资源管理更安全
*/
