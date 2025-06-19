package main

import (
	"fmt"
	"time" // 导入时间处理包
)

func main() {
	p := fmt.Println

	// 获取当前时间
	now := time.Now()
	p(now)

	// 创建指定时间
	// time.Date(年, 月, 日, 时, 分, 秒, 纳秒, 时区)
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 获取时间的各个组成部分
	p(then.Year())       // 年份
	p(then.Month())      // 月份
	p(then.Day())        // 日期
	p(then.Hour())       // 小时
	p(then.Minute())     // 分钟
	p(then.Second())     // 秒
	p(then.Nanosecond()) // 纳秒
	p(then.Location())   // 时区

	// 获取星期几
	p(then.Weekday())

	// 时间比较
	p(then.Before(now)) // then是否在now之前
	p(then.After(now))  // then是否在now之后
	p(then.Equal(now))  // then是否等于now

	// 计算时间差
	diff := now.Sub(then)
	p(diff)

	// 获取时间差的不同单位表示
	p(diff.Hours())       // 小时数
	p(diff.Minutes())     // 分钟数
	p(diff.Seconds())     // 秒数
	p(diff.Nanoseconds()) // 纳秒数

	// 时间加减操作
	p(then.Add(diff))  // then + diff
	p(then.Add(-diff)) // then - diff
}

/*
Go语言时间处理示例 - time包使用

主旨：
1. 演示time包中时间创建、获取、比较、计算的方法
2. 展示时间对象的各个组成部分
3. 理解时间差的计算和表示
4. 学习时间加减操作

关键特性：
- time：Go标准库的时间处理包
- 支持时间的创建、格式化、比较、计算
- 提供丰富的时区支持
- 支持高精度时间操作（纳秒级）

时间创建方法：
- time.Now()：获取当前时间
- time.Date()：创建指定时间
- time.Parse()：从字符串解析时间
- time.Unix()：从Unix时间戳创建时间

时间组成部分：
- Year()：年份
- Month()：月份
- Day()：日期
- Hour()：小时
- Minute()：分钟
- Second()：秒
- Nanosecond()：纳秒
- Location()：时区
- Weekday()：星期几

时间比较方法：
- Before(t)：是否在指定时间之前
- After(t)：是否在指定时间之后
- Equal(t)：是否等于指定时间

时间差计算：
- Sub(t)：计算两个时间的差值
- Hours()：时间差的小时数
- Minutes()：时间差的分钟数
- Seconds()：时间差的秒数
- Nanoseconds()：时间差的纳秒数

时间加减操作：
- Add(d)：时间加时间差
- AddDate(y, m, d)：时间加年月日

输出示例：
2024-01-15 10:30:45.123456789 +0800 CST
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
123456h 55m 46.472069552s
123456.929297
7407415.757821
444444945.46926
444444945469260000
2024-01-15 10:30:45.123456789 +0800 CST
1995-05-22 06:39:12.179317685 +0800 CST

时间格式化：
```go
// 自定义格式
t.Format("2006-01-02 15:04:05")
t.Format("2006年01月02日 15时04分05秒")

// 常用格式常量
time.RFC3339     // "2006-01-02T15:04:05Z07:00"
time.RFC822      // "02 Jan 06 15:04 MST"
time.ANSIC       // "Mon Jan _2 15:04:05 2006"
```

时区处理：
```go
// 获取指定时区
loc, _ := time.LoadLocation("Asia/Shanghai")
t := time.Now().In(loc)

// 常用时区
time.UTC        // 协调世界时
time.Local      // 本地时区
```

实际应用场景：
- 日志记录和时间戳
- 定时任务和调度
- 性能测量和基准测试
- 数据分析和统计
- 缓存过期时间
- 会话超时处理

注意事项：
- Go使用2006-01-02 15:04:05作为时间格式模板
- 时间对象是不可变的，操作返回新对象
- 时区信息很重要，避免时区混淆
- 高精度时间操作注意性能影响
- 时间比较要考虑时区因素
*/
