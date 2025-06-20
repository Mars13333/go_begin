package main

import (
	"fmt"
	"testing"
)

// IntMin 返回两个整数中的最小值
// 这是我们要测试的函数
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TestIntMinBasic 是一个基础的单元测试
// 测试函数必须以 Test 开头，接受 *testing.T 参数
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Errorf 报告测试失败，但继续执行其他测试
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// TestIntMinTableDriven 展示表驱动测试模式
// 这是 Go 中推荐的测试多个用例的方法
func TestIntMinTableDriven(t *testing.T) {
	// 定义测试用例表：输入和期望输出
	var tests = []struct {
		a, b int // 输入参数
		want int // 期望结果
	}{
		{0, 1, 0},   // 测试用例1
		{1, 0, 0},   // 测试用例2
		{2, -2, -2}, // 测试用例3：负数
		{0, -1, -1}, // 测试用例4：零和负数
		{-1, 0, -1}, // 测试用例5：负数和零
	}

	// 遍历每个测试用例
	for _, tt := range tests {
		// 为每个子测试创建描述性名称
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		// t.Run 创建子测试，每个子测试独立运行
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// BenchmarkIntMin 是基准测试函数
// 基准测试函数必须以 Benchmark 开头，接受 *testing.B 参数
func BenchmarkIntMin(b *testing.B) {
	// b.N 是基准测试框架自动调整的迭代次数
	// 框架会运行足够多次以获得稳定的性能测量
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

/*
Testing and Benchmarking 测试和基准测试示例主旨：

1. 核心功能：
   演示 Go 语言内置的测试框架，包括单元测试和性能基准测试。

2. 测试类型：
   - 单元测试：验证函数功能的正确性
   - 表驱动测试：用结构化方式测试多个用例
   - 基准测试：测量函数的性能表现

3. 测试函数规范：
   - 测试函数：func TestXxx(*testing.T)
   - 基准函数：func BenchmarkXxx(*testing.B)
   - 文件名：必须以 _test.go 结尾

4. 运行命令：
   - go test：运行所有测试
   - go test -v：详细输出
   - go test -bench=.：运行基准测试
   - go test -run=TestName：运行特定测试

5. 测试方法：
   - t.Errorf()：报告错误但继续执行
   - t.Fatalf()：报告错误并立即停止
   - t.Run()：创建子测试
   - t.Skip()：跳过测试

6. 表驱动测试优势：
   - 结构化：测试用例清晰组织
   - 可扩展：容易添加新的测试用例
   - 可读性：输入输出一目了然
   - 子测试：每个用例独立运行和报告

7. 基准测试特点：
   - 自动调整：b.N 自动调整迭代次数
   - 性能测量：测量执行时间和内存分配
   - 比较工具：可以比较不同实现的性能
   - 输出格式：ns/op（纳秒每操作）

8. 最佳实践：
   - 为每个公开函数编写测试
   - 使用表驱动测试处理多个用例
   - 测试边界条件和错误情况
   - 保持测试简单和专注
   - 使用有意义的测试名称

9. 测试覆盖率：
   - go test -cover：显示代码覆盖率
   - go test -coverprofile=coverage.out：生成覆盖率报告
   - go tool cover -html=coverage.out：生成HTML报告

10. CI/CD 集成：
    - 测试是持续集成的重要组成部分
    - 确保代码质量和功能正确性
    - 性能回归检测

这个示例展示了 Go 测试的核心模式，
是保证代码质量和性能的重要工具。
*/
