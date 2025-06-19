package main

import (
	"fmt"
	s "strings" // 导入strings包并别名为s，方便调用
)

// 创建fmt.Println的别名，简化代码
var p = fmt.Println

func main() {

	// 检查字符串是否包含子串
	p("Contains:  ", s.Contains("test", "es"))
	// 统计子串在字符串中出现的次数
	p("Count:     ", s.Count("test", "t"))
	// 检查字符串是否以指定前缀开头
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	// 检查字符串是否以指定后缀结尾
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	// 查找子串在字符串中第一次出现的位置（索引从0开始）
	p("Index:     ", s.Index("test", "e"))
	// 使用分隔符连接字符串切片
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	// 重复字符串指定次数
	p("Repeat:    ", s.Repeat("a", 5))
	// Replace函数：替换字符串中的子串
	// 语法：Replace(s, old, new, n)
	// s: 原字符串, old: 要替换的子串, new: 新子串, n: 替换次数
	// n = -1 表示替换所有匹配项
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	// n = 1 表示只替换第一个匹配项
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	// 按分隔符分割字符串为切片
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	// 转换为小写
	p("ToLower:   ", s.ToLower("TEST"))
	// 转换为大写
	p("ToUpper:   ", s.ToUpper("test"))
}

/*
Go语言字符串函数示例 - strings包常用函数

主旨：
1. 演示strings包中常用字符串操作函数
2. 展示字符串的查找、替换、分割、连接等操作
3. 理解字符串处理的基本方法
4. 特别说明Replace函数的不同用法

关键特性：
- strings包提供了丰富的字符串操作函数
- 大部分函数都是纯函数，不修改原字符串
- 支持字符串查找、替换、分割、连接等操作
- 函数命名清晰，易于理解和使用

Replace函数详解：

Replace(s, old, new, n) 函数参数说明：
- s: 原字符串
- old: 要替换的子串
- new: 新子串
- n: 替换次数

两个示例的区别：

1. s.Replace("foo", "o", "0", -1)
   - 原字符串: "foo"
   - 查找: "o"
   - 替换为: "0"
   - 次数: -1 (替换所有匹配项)
   - 结果: "f00" (两个o都被替换为0)

2. s.Replace("foo", "o", "0", 1)
   - 原字符串: "foo"
   - 查找: "o"
   - 替换为: "0"
   - 次数: 1 (只替换第一个匹配项)
   - 结果: "f0o" (只有第一个o被替换为0)

Replace函数替换次数参数说明：
- n > 0: 替换前n个匹配项
- n = 0: 不进行任何替换
- n < 0: 替换所有匹配项

其他重要函数说明：

Contains: 检查字符串是否包含子串
Count: 统计子串出现次数
HasPrefix/HasSuffix: 检查前缀/后缀
Index: 查找子串位置（返回-1表示未找到）
Join: 连接字符串切片
Repeat: 重复字符串
Split: 分割字符串为切片
ToLower/ToUpper: 大小写转换

输出结果示例：
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
ToLower:    test
ToUpper:    TEST

注意事项：
- 字符串在Go中是不可变的，所有操作都返回新字符串
- 大部分函数区分大小写
- Index函数返回-1表示未找到子串
- Join函数可以连接任意字符串切片
*/
