package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// os.CreateTemp 创建临时文件
	// 第一个参数是目录（""表示使用系统默认临时目录）
	// 第二个参数是文件名前缀，系统会自动添加随机后缀
	f, err := os.CreateTemp("", "sample")
	check(err)

	// 显示生成的临时文件路径
	// 通常类似：/tmp/sample123456789 (Linux/macOS) 或 C:\Users\...\Temp\sample123456789 (Windows)
	fmt.Println("Temp file name:", f.Name())

	// defer 确保程序结束时删除临时文件
	defer os.Remove(f.Name())

	// 向临时文件写入数据
	// f 是已打开的文件句柄，可以直接写入
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// os.MkdirTemp 创建临时目录
	// 参数与 CreateTemp 类似：目录位置和名称前缀
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	// defer 确保程序结束时删除临时目录及其内容
	defer os.RemoveAll(dname)

	// 在临时目录中创建文件
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}

/*
Temporary Files and Directories 临时文件和目录示例主旨：

1. 核心功能：
   演示如何在 Go 中安全地创建和使用临时文件和目录，
   这些资源会在程序结束时自动清理。

2. 主要方法：
   - os.CreateTemp()：创建临时文件并返回文件句柄
   - os.MkdirTemp()：创建临时目录并返回路径
   - defer + os.Remove()：确保临时文件被删除
   - defer + os.RemoveAll()：确保临时目录及其内容被删除

3. 系统临时目录：
   - Linux/macOS：通常是 /tmp
   - Windows：通常是 %TEMP% 或 %TMP% 环境变量指定的位置
   - 传入空字符串 "" 会使用系统默认临时目录

4. 文件命名规则：
   - 前缀 + 随机后缀的组合
   - 确保文件名唯一性，避免冲突
   - 示例：sample123456789、sampledir987654321

5. 资源管理：
   - 使用 defer 确保资源清理
   - os.Remove()：删除单个文件
   - os.RemoveAll()：递归删除目录及其所有内容
   - 即使程序异常退出，defer 也会执行清理

6. 权限设置：
   - 临时文件默认权限：通常是 0600（仅所有者可读写）
   - 临时目录默认权限：通常是 0700（仅所有者可访问）
   - WriteFile 使用 0666 权限（所有用户可读写）

7. 实际应用场景：
   - 缓存文件：临时存储处理中的数据
   - 下载文件：先下载到临时位置，完成后移动
   - 测试环境：创建测试数据而不污染系统
   - 数据处理：大文件分块处理的中间结果
   - 配置备份：修改配置前的临时备份

8. 安全考虑：
   - 临时文件具有唯一名称，防止冲突
   - 适当的权限设置保护敏感数据
   - 自动清理防止磁盘空间泄漏
   - 在系统临时目录中创建，遵循系统安全策略

9. 最佳实践：
   - 总是使用 defer 进行清理
   - 检查所有操作的错误
   - 使用有意义的前缀名称
   - 考虑在程序开始时清理可能残留的临时文件

这个示例展示了 Go 中处理临时文件的标准模式，
是文件处理、缓存系统和测试代码中的重要工具。
*/
