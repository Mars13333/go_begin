package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// filepath.Join 用于跨平台地连接路径组件
	// 会自动使用正确的路径分隔符（Windows用\，Unix用/）
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Join 会清理路径，处理多余的分隔符和相对路径
	fmt.Println(filepath.Join("dir1//", "filename"))       // 清理多余的 //
	fmt.Println(filepath.Join("dir1/../dir1", "filename")) // 处理 .. 相对路径

	// filepath.Dir 返回路径的目录部分（除了最后一个元素）
	fmt.Println("Dir(p):", filepath.Dir(p))
	// filepath.Base 返回路径的最后一个元素（文件名）
	fmt.Println("Base(p):", filepath.Base(p))

	// filepath.IsAbs 检查路径是否为绝对路径
	fmt.Println(filepath.IsAbs("dir/file"))  // false - 相对路径
	fmt.Println(filepath.IsAbs("/dir/file")) // true - 绝对路径

	filename := "config.json"

	// filepath.Ext 获取文件扩展名（包含点号）
	ext := filepath.Ext(filename)
	fmt.Println(ext) // 输出: .json

	// 使用 strings.TrimSuffix 去除扩展名，获取文件名主体
	fmt.Println(strings.TrimSuffix(filename, ext)) // 输出: config

	// filepath.Rel 计算从一个路径到另一个路径的相对路径
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel) // 输出: t/file

	// 计算跨目录的相对路径
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel) // 输出: ../c/t/file
}

/*
File Path 文件路径处理示例主旨：

1. 核心功能：
   Go 的 filepath 包提供了跨平台的文件路径操作功能，
   自动处理不同操作系统的路径分隔符差异。

2. 主要方法：
   - filepath.Join()：连接路径组件，自动使用正确的分隔符
   - filepath.Dir()：获取路径的目录部分
   - filepath.Base()：获取路径的文件名部分
   - filepath.Ext()：获取文件扩展名
   - filepath.IsAbs()：检查是否为绝对路径
   - filepath.Rel()：计算相对路径

3. 跨平台特性：
   - Windows: 使用反斜杠 (\) 作为路径分隔符
   - Unix/Linux/macOS: 使用正斜杠 (/) 作为路径分隔符
   - filepath 包自动处理这些差异

4. 路径清理：
   - 自动处理多余的分隔符 (dir1// -> dir1/)
   - 处理相对路径标记 (dir1/../dir1 -> dir1)
   - 规范化路径格式

5. 实用场景：
   - 构建配置文件路径
   - 处理用户上传文件的路径
   - 计算相对路径用于链接生成
   - 文件名和扩展名的分离处理

6. 安全性：
   - 避免手动字符串拼接可能导致的路径错误
   - 防止路径遍历攻击（通过路径清理）
   - 确保跨平台兼容性

这个示例展示了 Go 中进行文件路径操作的标准方法，
是文件系统操作和Web开发中的重要工具。
*/
