package main

import (
	"fmt"
	"time" // 导入时间处理包
)

func main() {

	// 获取当前时间
	now := time.Now()
	fmt.Println(now)

	// 获取Unix时间戳（秒级，从1970年1月1日00:00:00 UTC开始计算）
	fmt.Println(now.Unix())
	// 获取Unix时间戳（毫秒级）
	fmt.Println(now.UnixMilli())
	// 获取Unix时间戳（纳秒级）
	fmt.Println(now.UnixNano())

	// 从Unix时间戳（秒级）创建时间对象
	fmt.Println(time.Unix(now.Unix(), 0))
	// 从Unix时间戳（纳秒级）创建时间对象
	fmt.Println(time.Unix(0, now.UnixNano()))
}

/*
Go语言Unix时间戳示例 - Epoch时间处理

主旨：
1. 演示Unix时间戳的获取和转换方法
2. 展示不同精度的时间戳格式（秒、毫秒、纳秒）
3. 理解Unix Epoch的概念和意义
4. 学习时间戳与时间对象的相互转换

关键概念：
- Unix Epoch：Unix时间的起点，1970年1月1日00:00:00 UTC
- Unix时间戳：从Epoch开始经过的秒数
- 时间戳精度：秒、毫秒、纳秒等不同精度

Unix时间戳获取方法：
- Unix()：获取秒级时间戳（int64）
- UnixMilli()：获取毫秒级时间戳（int64）
- UnixNano()：获取纳秒级时间戳（int64）

时间戳转换为时间对象：
- time.Unix(sec, nsec)：从秒和纳秒创建时间
- time.Unix(sec, 0)：从秒级时间戳创建时间
- time.Unix(0, nsec)：从纳秒级时间戳创建时间

输出示例：
2024-01-15 10:30:45.123456789 +0800 CST
1705294245
1705294245123
1705294245123456789
2024-01-15 10:30:45 +0800 CST
2024-01-15 10:30:45.123456789 +0800 CST

Unix Epoch详解：
- 起始时间：1970年1月1日00:00:00 UTC
- 也称为：Unix时间、POSIX时间、Epoch时间
- 标准：ISO 8601和RFC 3339
- 用途：计算机系统中的标准时间表示

时间戳精度对比：
- 秒级：适合大多数应用场景
- 毫秒级：适合需要毫秒精度的场景
- 纳秒级：适合高精度时间测量

实际应用场景：
- 数据库时间字段存储
- API接口时间参数
- 日志时间戳记录
- 缓存过期时间
- 文件修改时间
- 网络协议时间同步

时间戳优势：
- 标准化：跨平台、跨语言兼容
- 紧凑：数字格式，存储效率高
- 计算方便：可以直接进行数学运算
- 时区无关：统一使用UTC基准

注意事项：
- Unix时间戳使用UTC时区
- 2038年问题：32位系统的时间戳溢出
- 精度选择：根据应用需求选择合适的精度
- 时区转换：时间戳转换为本地时间需要考虑时区

相关函数：
```go
// 获取当前时间戳
now := time.Now()
sec := now.Unix()        // 秒
msec := now.UnixMilli()  // 毫秒
nsec := now.UnixNano()   // 纳秒

// 从时间戳创建时间
t1 := time.Unix(sec, 0)           // 从秒创建
t2 := time.Unix(0, nsec)          // 从纳秒创建
t3 := time.Unix(sec, nsec%1e9)    // 从秒和纳秒创建
```
*/
