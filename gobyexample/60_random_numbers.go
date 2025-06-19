package main

import (
	"fmt"
	"math/rand/v2" // 导入Go 1.22+的新随机数包
)

func main() {

	// 生成0到99之间的随机整数
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// 生成0.0到1.0之间的随机浮点数
	fmt.Println(rand.Float64())

	// 生成5.0到10.0之间的随机浮点数
	// 公式：(rand.Float64() * 5) + 5
	// rand.Float64()生成0-1，乘以5得到0-5，加5得到5-10
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 使用PCG算法创建可重现的随机数生成器
	// NewPCG(seed1, seed2)：使用两个种子初始化PCG生成器
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	// 使用相同的种子创建另一个生成器，结果应该相同
	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}

/*
Go语言随机数生成示例 - math/rand/v2包使用

主旨：
1. 演示math/rand/v2包中随机数生成的方法
2. 展示整数和浮点数随机数的生成
3. 理解随机数种子和可重现性
4. 学习PCG算法和自定义随机数生成器

关键特性：
- math/rand/v2：Go 1.22+的新随机数包，性能更好
- 支持整数和浮点数随机数生成
- 提供PCG（Permuted Congruential Generator）算法
- 支持可重现的随机数序列

随机数生成函数：
- IntN(n)：生成[0, n)范围内的随机整数
- Float64()：生成[0.0, 1.0)范围内的随机浮点数
- Float32()：生成[0.0, 1.0)范围内的随机浮点数

PCG算法说明：
- PCG：Permuted Congruential Generator
- 高质量、高性能的随机数生成算法
- 支持两个种子参数，提供更好的随机性
- 适合需要可重现随机数的场景

种子和可重现性：
- 相同种子产生相同的随机数序列
- 用于测试、调试、游戏等需要确定性的场景
- 不同种子产生不同的随机数序列

输出示例：
87,65
0.123456789
7.89,9.12
42,17
42,17

随机数范围控制：
```go
// 整数范围
rand.IntN(100)           // [0, 100)
rand.IntN(10) + 1        // [1, 11)

// 浮点数范围
rand.Float64() * 10      // [0.0, 10.0)
(rand.Float64() * 5) + 5 // [5.0, 10.0)

// 负范围
rand.IntN(100) - 50      // [-50, 50)
```

PCG生成器使用：
```go
// 创建PCG生成器
s := rand.NewPCG(seed1, seed2)
r := rand.New(s)

// 使用生成器
r.IntN(100)
r.Float64()
```

实际应用场景：
- 游戏开发（随机事件、地图生成）
- 测试数据生成
- 模拟和建模
- 密码学应用
- 算法随机化
- 负载测试

随机数质量：
- 全局随机数：适合一般用途
- PCG生成器：高质量、可重现
- 加密随机数：crypto/rand包，用于安全场景

注意事项：
- 默认使用全局随机数生成器
- 相同种子产生相同序列
- 不同种子产生不同序列
- PCG算法性能优于旧版本
- 不适合密码学用途（使用crypto/rand）

性能考虑：
- math/rand/v2比v1版本性能更好
- PCG算法比线性同余法质量更高
- 全局生成器适合大多数场景
- 自定义生成器适合需要控制的场景

安全注意事项：
- math/rand不适用于密码学
- 需要加密随机数时使用crypto/rand
- 种子不应该从可预测的来源获取
- 避免在安全敏感场景使用可重现的随机数
*/
