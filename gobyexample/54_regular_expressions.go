package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// 直接匹配字符串，返回是否匹配
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 编译正则表达式，返回*regexp.Regexp对象
	// 编译后的正则表达式可以重复使用，性能更好
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 使用编译后的正则表达式匹配字符串
	fmt.Println(r.MatchString("peach"))

	// 查找第一个匹配的字符串
	fmt.Println(r.FindString("peach punch"))

	// 查找第一个匹配的字符串的索引位置 [开始位置, 结束位置]
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// 查找第一个匹配的字符串及其子匹配（捕获组）
	// 返回完整匹配和所有子匹配的切片
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// 查找第一个匹配的字符串及其子匹配的索引位置
	// 返回完整匹配和所有子匹配的索引切片
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 查找所有匹配的字符串，-1表示查找所有匹配
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// 查找所有匹配的字符串及其子匹配的索引位置
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// 查找前2个匹配的字符串
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 匹配字节切片
	fmt.Println(r.Match([]byte("peach")))

	// 使用MustCompile编译正则表达式，如果编译失败会panic
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// 替换所有匹配的字符串
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// 使用函数替换匹配的字符串
	// ReplaceAllFunc接受一个函数，对每个匹配项调用该函数
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

/*
Go语言正则表达式示例 - regexp包使用

主旨：
1. 演示regexp包中常用正则表达式函数的使用方法
2. 展示正则表达式的编译、匹配、查找、替换等操作
3. 理解正则表达式的基本语法和捕获组概念
4. 学习如何高效处理文本匹配和替换

关键特性：
- regexp：Go标准库的正则表达式包
- 支持编译和直接匹配两种方式
- 提供丰富的查找、替换、分割功能
- 支持字节切片和字符串操作

正则表达式语法说明：
- p([a-z]+)ch：匹配以p开头，ch结尾，中间是1个或多个小写字母的字符串
- ([a-z]+)：捕获组，匹配1个或多个小写字母
- +：量词，表示1个或多个
- [a-z]：字符类，匹配任意小写字母

常用函数详解：

1. 匹配函数：
   - MatchString(pattern, s)：直接匹配字符串
   - Compile(pattern)：编译正则表达式
   - MustCompile(pattern)：编译正则表达式，失败时panic

2. 查找函数：
   - FindString(s)：查找第一个匹配
   - FindStringIndex(s)：查找第一个匹配的索引
   - FindStringSubmatch(s)：查找第一个匹配及其子匹配
   - FindAllString(s, n)：查找所有匹配，n=-1表示查找所有

3. 替换函数：
   - ReplaceAllString(s, repl)：替换所有匹配
   - ReplaceAllFunc(s, func)：使用函数替换匹配

输出示例：
true
true
peach
idx: [0 5]
[peach ea]
[0 5 1 3]
[peach punch pinch]
all: [[0 5 1 3] [6 11 7 10] [12 18 13 17]]
[peach punch]
true
regexp: p([a-z]+)ch
a <fruit>
a PEACH

性能优化建议：
- 对于重复使用的正则表达式，使用Compile编译后复用
- 避免在循环中重复编译正则表达式
- 使用MustCompile简化错误处理（确定正则表达式正确时）

实际应用场景：
- 文本验证（邮箱、手机号、身份证等）
- 文本提取和解析
- 文本替换和格式化
- 日志分析和处理
- 配置文件解析

注意事项：
- 正则表达式语法遵循RE2标准
- 不支持某些高级特性（如回溯引用）
- 编译失败时会返回错误，需要处理
- 捕获组索引从1开始（0是完整匹配）
*/
