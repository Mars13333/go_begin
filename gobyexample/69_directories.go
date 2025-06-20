package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// os.Mkdir 创建单级目录，权限为 0755 (rwxr-xr-x)
	err := os.Mkdir("subdir", 0755)
	check(err)

	// defer 确保程序结束时清理创建的目录和文件
	defer os.RemoveAll("subdir")

	// 创建空文件的辅助函数
	createEmptyFile := func(name string) {
		d := []byte("")                    // 空字节切片
		check(os.WriteFile(name, d, 0644)) // 权限 0644 (rw-r--r--)
	}

	// 在子目录中创建文件
	createEmptyFile("subdir/file1")

	// os.MkdirAll 递归创建多级目录（类似 mkdir -p）
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	// 创建更多测试文件
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// os.ReadDir 读取目录内容，返回 DirEntry 切片
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	// 遍历目录条目，显示名称和是否为目录
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// os.Chdir 改变当前工作目录（类似 cd 命令）
	err = os.Chdir("subdir/parent/child")
	check(err)

	// 读取当前目录（"."）的内容
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 返回到原始目录（向上三级）
	err = os.Chdir("../../..")
	check(err)

	// filepath.WalkDir 递归遍历目录树
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
	check(err)
}

// visit 是 filepath.WalkDir 的回调函数
// 对遍历到的每个文件和目录都会调用此函数
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err // 如果有错误，返回错误停止遍历
	}
	// 打印路径和是否为目录
	fmt.Println(" ", path, d.IsDir())
	return nil // 返回 nil 继续遍历
}

/*
Directories 目录操作示例主旨：

1. 核心功能：
   演示 Go 语言中常用的目录操作，包括创建、读取、遍历和导航目录。

2. 主要操作：
   - os.Mkdir()：创建单级目录
   - os.MkdirAll()：递归创建多级目录（mkdir -p）
   - os.ReadDir()：读取目录内容
   - os.Chdir()：改变当前工作目录
   - os.RemoveAll()：递归删除目录及其内容
   - filepath.WalkDir()：递归遍历目录树

3. 权限设置：
   - 0755：目录权限（rwxr-xr-x）- 所有者可读写执行，组和其他用户可读执行
   - 0644：文件权限（rw-r--r--）- 所有者可读写，组和其他用户只读

4. 目录结构示例：
   subdir/
   ├── file1
   └── parent/
       ├── file2
       ├── file3
       └── child/
           └── file4

5. 遍历模式：
   - 直接读取：使用 os.ReadDir() 读取单个目录
   - 递归遍历：使用 filepath.WalkDir() 遍历整个目录树
   - 工作目录切换：使用 os.Chdir() 改变当前位置

6. 错误处理：
   - 每个操作都检查错误
   - 使用 defer 确保资源清理
   - visit 函数中的错误会停止遍历

7. 实际应用场景：
   - 文件系统管理工具
   - 日志文件整理
   - 项目构建系统
   - 文件备份和同步
   - 配置文件查找

8. 安全考虑：
   - 适当的文件权限设置
   - 错误处理防止程序崩溃
   - 使用 defer 确保资源清理

这个示例展示了 Go 中进行目录操作的标准模式，
是文件系统编程和系统管理工具开发的基础。
*/
