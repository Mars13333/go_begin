package main

import "os"

func main() {
	// 主动触发panic，程序会立即停止执行并打印错误信息
	// 注意：这行代码之后的代码永远不会执行
	panic("a problem")

	// 以下代码永远不会执行，因为上面的panic会终止程序
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

/*
Go语言panic机制示例 - 程序异常终止

主旨：
1. 演示panic函数的使用和效果
2. 说明panic会立即终止程序执行
3. 展示panic在错误处理中的应用
4. 理解panic与recover的配合使用

关键特性：
- panic(): 主动触发程序异常终止
- panic会立即停止当前函数执行，开始执行defer函数
- 如果没有recover捕获，程序会完全退出
- panic后的代码永远不会执行

如果注释掉 panic("a problem") 会发生什么：

1. 程序会继续执行到os.Create("/tmp/file")
2. 尝试在/tmp目录下创建file文件
3. 如果创建成功（err == nil），程序正常结束
4. 如果创建失败（err != nil），会执行panic(err)，程序异常终止

执行结果对比：
- 有panic("a problem"): 程序立即终止，输出 "panic: a problem"
- 注释掉后: 程序尝试创建文件，成功则正常结束，失败则panic(err)

注意事项：
- panic应该用于不可恢复的错误情况
- 生产环境中应该使用recover来捕获panic
- 文件操作等可能失败的操作应该使用正常的错误处理机制
*/
