package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	// 以默认格式输出结构体
	fmt.Printf("struct1: %v\n", p)

	// 以字段名+值的格式输出结构体
	fmt.Printf("struct2: %+v\n", p)

	// 以Go语法格式输出结构体（含类型信息）
	fmt.Printf("struct3: %#v\n", p)

	// 输出变量的类型
	fmt.Printf("type: %T\n", p)

	// 输出布尔值
	fmt.Printf("bool: %t\n", true)

	// 十进制整数
	fmt.Printf("int: %d\n", 123)

	// 二进制整数
	fmt.Printf("bin: %b\n", 14)

	// 输出对应的Unicode字符
	fmt.Printf("char: %c\n", 33)

	// 十六进制整数
	fmt.Printf("hex: %x\n", 456)

	// 浮点数（小数点形式）
	fmt.Printf("float1: %f\n", 78.9)

	// 浮点数（科学计数法，小写e）
	fmt.Printf("float2: %e\n", 123400000.0)
	// 浮点数（科学计数法，大写E）
	fmt.Printf("float3: %E\n", 123400000.0)

	// 普通字符串
	fmt.Printf("str1: %s\n", "\"string\"")

	// 带双引号的字符串（Go语法格式）
	fmt.Printf("str2: %q\n", "\"string\"")

	// 字符串的十六进制表示
	fmt.Printf("str3: %x\n", "hex this")

	// 指针的值（内存地址）
	fmt.Printf("pointer: %p\n", &p)

	// 指定宽度的整数，右对齐
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// 指定宽度和精度的浮点数，右对齐
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// 指定宽度和精度的浮点数，左对齐
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// 指定宽度的字符串，右对齐
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// 指定宽度的字符串，左对齐
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// Sprintf格式化为字符串
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// Fprintf格式化输出到os.Stderr
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}

/*
Go语言字符串格式化示例 - fmt包格式化输出

主旨：
1. 演示fmt包中Printf/Sprintf/Fprintf等格式化输出函数的用法
2. 展示各种常用格式化占位符的效果
3. 理解宽度、精度、对齐等格式化参数

常用格式化占位符说明：
- %v    ：默认格式输出
- %+v   ：结构体字段名+值
- %#v   ：Go语法格式输出
- %T    ：类型
- %t    ：布尔值
- %d    ：十进制整数
- %b    ：二进制整数
- %c    ：Unicode字符
- %x    ：十六进制整数或字符串
- %f    ：浮点数（小数点）
- %e/%E ：浮点数（科学计数法）
- %s    ：字符串
- %q    ：带引号字符串
- %p    ：指针
- %6d   ：宽度为6的整数，右对齐
- %-6d  ：宽度为6的整数，左对齐
- %.2f  ：保留2位小数的浮点数
- %6.2f ：宽度为6，保留2位小数，右对齐
- %-6.2f：宽度为6，保留2位小数，左对齐
- %6s   ：宽度为6的字符串，右对齐
- %-6s  ：宽度为6的字符串，左对齐

其他说明：
- Printf: 格式化输出到标准输出
- Sprintf: 格式化为字符串
- Fprintf: 格式化输出到指定io.Writer（如os.Stderr）
- 格式化参数可组合使用，灵活控制输出样式
- 结构体格式化可用于调试和日志输出
*/
