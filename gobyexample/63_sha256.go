package main

import (
	"crypto/sha256" // 导入SHA256哈希算法包
	"fmt"
)

func main() {
	// 定义要计算哈希的字符串
	s := "sha256 this string"

	// 创建SHA256哈希对象
	h := sha256.New()

	// 将字符串转换为字节切片并写入哈希对象
	h.Write([]byte(s))

	// 计算最终的哈希值，返回字节切片
	// Sum(nil)表示不追加额外数据到哈希结果
	bs := h.Sum(nil)

	// 输出原始字符串
	fmt.Println(s)
	// 以十六进制格式输出哈希值
	// %x：小写十六进制，%X：大写十六进制
	fmt.Printf("%x\n", bs)
}

/*
Go语言SHA256哈希示例 - crypto/sha256包使用

主旨：
1. 演示crypto/sha256包中SHA256哈希计算的方法
2. 展示字节切片和字符串的哈希处理
3. 理解哈希算法的基本概念和用途
4. 学习十六进制输出格式

关键特性：
- crypto/sha256：Go标准库的SHA256哈希算法包
- SHA256：安全哈希算法，产生256位（32字节）哈希值
- 支持字符串和字节切片的哈希计算
- 提供流式哈希计算接口

SHA256算法说明：
- SHA256：Secure Hash Algorithm 256-bit
- 输入：任意长度的数据
- 输出：256位（32字节）的固定长度哈希值
- 特性：确定性、雪崩效应、抗碰撞性
- 用途：数据完整性验证、数字签名、密码存储

哈希计算步骤：
1. 创建哈希对象：sha256.New()
2. 写入数据：h.Write([]byte(data))
3. 计算哈希：h.Sum(nil)
4. 格式化输出：十六进制表示

输出示例：
sha256 this string
1af1dfa857bf1d8814fe1af8983c18076819922e8f0a88bdf6d0191d2d6836d9

哈希算法对比：
- MD5：128位，已不推荐用于安全用途
- SHA1：160位，已不推荐用于安全用途
- SHA256：256位，当前推荐标准
- SHA512：512位，更高安全性

常用哈希函数：
```go
// SHA256
h := sha256.New()
h.Write([]byte("data"))
hash := h.Sum(nil)

// SHA1（不推荐用于安全用途）
h1 := sha1.New()
h1.Write([]byte("data"))
hash1 := h1.Sum(nil)

// MD5（不推荐用于安全用途）
h2 := md5.New()
h2.Write([]byte("data"))
hash2 := h2.Sum(nil)
```

十六进制输出格式：
```go
fmt.Printf("%x\n", hash)  // 小写十六进制
fmt.Printf("%X\n", hash)  // 大写十六进制
fmt.Printf("%x", hash)    // 无换行符
```

实际应用场景：
- 文件完整性验证
- 密码哈希存储
- 数字签名
- 区块链技术
- 数据去重
- 缓存键生成

安全注意事项：
- SHA256是密码学安全的哈希算法
- 适合用于数据完整性验证
- 不适合直接用于密码存储（应使用bcrypt、scrypt等）
- 哈希值不可逆，无法从哈希值恢复原始数据

性能考虑：
- SHA256计算相对快速
- 适合大量数据处理
- 内存使用固定（32字节输出）
- 支持流式处理大文件

错误处理：
- SHA256计算不会失败
- 主要注意输入数据的编码
- 处理大文件时注意内存使用

相关包：
- crypto/sha1：SHA1哈希算法
- crypto/md5：MD5哈希算法
- crypto/sha512：SHA512哈希算法
- crypto/hmac：HMAC消息认证码
- crypto/rand：加密随机数生成

哈希值验证：
```go
// 计算两个数据的哈希值并比较
hash1 := sha256.Sum256([]byte("data1"))
hash2 := sha256.Sum256([]byte("data2"))
if hash1 == hash2 {
    fmt.Println("哈希值相同")
} else {
    fmt.Println("哈希值不同")
}
```

文件哈希计算：
```go
// 计算文件哈希
file, _ := os.Open("file.txt")
defer file.Close()

h := sha256.New()
io.Copy(h, file)
hash := h.Sum(nil)
fmt.Printf("%x\n", hash)
```
*/
