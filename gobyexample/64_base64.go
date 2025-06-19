package main

import (
	b64 "encoding/base64" // 导入Base64编码包，使用别名b64
	"fmt"
)

func main() {

	// 定义要编码的原始数据字符串
	// 包含字母、数字、特殊字符，用于演示Base64编码
	data := "abc123!?$*&()'-=@~"

	// 使用标准Base64编码
	// StdEncoding：标准Base64编码，使用A-Z, a-z, 0-9, +, /字符
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码Base64字符串回原始数据
	// DecodeString返回字节切片，需要转换为字符串
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 使用URL安全的Base64编码
	// URLEncoding：URL安全编码，将+替换为-，/替换为_
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	// 解码URL安全的Base64字符串
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

/*
Go语言Base64编码示例 - encoding/base64包使用

主旨：
1. 演示encoding/base64包中Base64编码和解码的方法
2. 展示标准编码和URL安全编码的区别
3. 理解Base64编码的基本原理和用途
4. 学习字节切片和字符串的转换

关键特性：
- encoding/base64：Go标准库的Base64编码包
- Base64：将二进制数据编码为ASCII字符串
- 支持标准编码和URL安全编码
- 提供编码和解码功能

Base64编码原理：
- 将3字节（24位）数据编码为4个6位字符
- 使用64个ASCII字符：A-Z, a-z, 0-9, +, /
- 如果数据长度不是3的倍数，用=填充
- 编码后数据长度增加约33%

编码类型对比：
1. 标准编码（StdEncoding）：
   - 字符集：A-Z, a-z, 0-9, +, /
   - 填充字符：=
   - 用途：一般数据传输

2. URL安全编码（URLEncoding）：
   - 字符集：A-Z, a-z, 0-9, -, _
   - 填充字符：=
   - 用途：URL参数、文件名等

输出示例：
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~

编码函数详解：
- EncodeToString(data)：编码为字符串
- DecodeString(s)：解码字符串
- Encode(dst, src)：编码到字节切片
- Decode(dst, src)：解码到字节切片

常用编码模式：
```go
// 标准编码
encoded := b64.StdEncoding.EncodeToString([]byte("Hello, World!"))
decoded, _ := b64.StdEncoding.DecodeString(encoded)

// URL安全编码
urlEncoded := b64.URLEncoding.EncodeToString([]byte("Hello, World!"))
urlDecoded, _ := b64.URLEncoding.DecodeString(urlEncoded)

// 无填充编码
noPadding := b64.StdEncoding.WithPadding(b64.NoPadding)
encoded := noPadding.EncodeToString([]byte("data"))
```

实际应用场景：
- 图片数据编码（data URI）
- 二进制文件传输
- API认证令牌
- 配置文件编码
- 数据库存储
- 网络协议数据

Base64编码表：
```
索引  字符  索引  字符  索引  字符  索引  字符
0     A     16    Q     32    g     48    w
1     B     17    R     33    h     49    x
2     C     18    S     34    i     50    y
3     D     19    T     35    j     51    z
4     E     20    U     36    k     52    0
5     F     21    V     37    l     53    1
6     G     22    W     38    m     54    2
7     H     23    X     39    n     55    3
8     I     24    Y     40    o     56    4
9     J     25    Z     41    p     57    5
10    K     26    a     42    q     58    6
11    L     27    b     43    r     59    7
12    M     28    c     44    s     60    8
13    N     29    d     45    t     61    9
14    O     30    e     46    u     62    +
15    P     31    f     47    v     63    /
```

性能考虑：
- Base64编码增加约33%的数据大小
- 编码/解码是CPU密集型操作
- 大文件编码考虑流式处理
- 频繁编码考虑缓存结果

错误处理：
```go
// 检查解码错误
decoded, err := b64.StdEncoding.DecodeString(encoded)
if err != nil {
    fmt.Println("解码错误:", err)
    return
}
```

安全注意事项：
- Base64不是加密，只是编码
- 不要用于敏感数据保护
- URL安全编码避免特殊字符问题
- 注意编码后的数据大小增加

相关包：
- encoding/hex：十六进制编码
- encoding/json：JSON编码
- crypto/base64：加密相关编码
- bytes：字节切片操作

文件编码示例：
```go
// 编码文件内容
file, _ := os.ReadFile("file.txt")
encoded := b64.StdEncoding.EncodeToString(file)

// 解码并保存文件
decoded, _ := b64.StdEncoding.DecodeString(encoded)
os.WriteFile("decoded.txt", decoded, 0644)
```
*/
