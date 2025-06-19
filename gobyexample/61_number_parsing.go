package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 解析浮点数字符串，64表示64位精度
	// ParseFloat(s, bitSize)：将字符串解析为float64
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 解析整数字符串，0表示自动检测进制，64表示64位
	// ParseInt(s, base, bitSize)：将字符串解析为int64
	// base=0：自动检测进制（0x开头为16进制，0开头为8进制，其他为10进制）
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// 解析16进制字符串
	// "0x1c8"是16进制表示，自动检测并解析为十进制
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// 解析无符号整数字符串
	// ParseUint(s, base, bitSize)：将字符串解析为uint64
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// 快速解析整数（默认10进制）
	// Atoi(s)：将字符串解析为int，等同于ParseInt(s, 10, 0)
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 演示解析错误：无法解析的字符串
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

/*
Go语言数字解析示例 - strconv包使用

主旨：
1. 演示strconv包中数字字符串解析的方法
2. 展示不同进制和精度的数字解析
3. 理解错误处理和返回值
4. 学习字符串到数字的转换技巧

关键特性：
- strconv：Go标准库的字符串转换包
- 支持整数、浮点数、无符号整数的解析
- 支持多种进制（2、8、10、16进制）
- 提供不同精度的数字类型支持

数字解析函数详解：

1. 浮点数解析：
   - ParseFloat(s, bitSize)：解析为float64
   - bitSize：32或64，表示精度
   - 支持科学计数法（如"1.23e4"）

2. 整数解析：
   - ParseInt(s, base, bitSize)：解析为int64
   - base：进制（0=自动检测，2=二进制，8=八进制，10=十进制，16=十六进制）
   - bitSize：8、16、32、64，表示位数

3. 无符号整数解析：
   - ParseUint(s, base, bitSize)：解析为uint64
   - 不支持负数

4. 快速解析：
   - Atoi(s)：快速解析为int（10进制）
   - 等同于ParseInt(s, 10, 0)

进制说明：
- 0：自动检测进制
- 2：二进制（如"1010"）
- 8：八进制（如"0777"）
- 10：十进制（如"123"）
- 16：十六进制（如"0xFF"）

自动检测规则：
- 0x或0X开头：16进制
- 0开头：8进制
- 其他：10进制

输出示例：
1.234
123
456
789
135
strconv.Atoi: parsing "wat": invalid syntax

错误处理：
```go
// 检查解析错误
if i, err := strconv.Atoi("123"); err != nil {
    fmt.Println("解析错误:", err)
} else {
    fmt.Println("解析成功:", i)
}
```

常用解析模式：
```go
// 浮点数
f, _ := strconv.ParseFloat("3.14", 64)

// 整数（自动检测进制）
i, _ := strconv.ParseInt("0xFF", 0, 64)  // 16进制
i, _ := strconv.ParseInt("0777", 0, 64)  // 8进制
i, _ := strconv.ParseInt("123", 0, 64)   // 10进制

// 无符号整数
u, _ := strconv.ParseUint("123", 10, 64)

// 快速解析
i, _ := strconv.Atoi("123")
```

实际应用场景：
- 配置文件数值解析
- 命令行参数处理
- 网络协议数据解析
- 数据库查询结果处理
- 用户输入验证
- 日志数据分析

性能考虑：
- Atoi比ParseInt更快（针对10进制）
- 错误检查会增加少量开销
- 频繁解析考虑缓存结果

注意事项：
- 总是检查错误返回值
- 注意数值范围限制
- 不同进制的表示方法
- 浮点数精度问题
- 无符号整数不支持负数

相关函数：
- strconv.FormatInt()：整数转字符串
- strconv.FormatFloat()：浮点数转字符串
- strconv.Itoa()：整数转字符串
- strconv.Quote()：字符串转引号格式
*/
