package main

import (
	"cmp" // 导入cmp包，提供比较函数，cmp是"compare"的缩写
	"fmt"
	"slices" // 导入slices包，提供切片操作功能
)

func main() {

	// 创建一个水果字符串切片用于自定义排序演示
	fruits := []string{"peach", "banana", "kiwi"}

	// 定义一个按字符串长度排序的比较函数
	// cmp.Compare(a, b) 返回：
	// -1 如果 a < b
	//  0 如果 a == b
	// +1 如果 a > b
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b)) // [kiwi peach banana]
		// return -cmp.Compare(len(a), len(b)) // [banana peach kiwi]

	}
	// 使用自定义比较函数对水果切片进行排序
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// 定义一个Person结构体用于复杂对象排序演示
	type Person struct {
		name string
		age  int
	}
	// 创建Person切片
	peoole := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	// 使用匿名函数按年龄对Person切片进行排序
	slices.SortFunc(peoole, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(peoole)

}

/*
Go语言自定义排序示例 - 使用slices.SortFunc进行自定义排序

主旨：
1. 演示如何使用slices.SortFunc进行自定义排序
2. 展示如何为不同类型的数据定义比较函数
3. 说明cmp包的作用和cmp.Compare函数的使用
4. 展示结构体切片的自定义排序方法

关键特性：
- slices.SortFunc()：使用自定义比较函数进行排序
- cmp.Compare()：提供标准化的比较结果（-1, 0, +1）
- 支持任意类型的自定义排序逻辑
- 原地排序，会修改原切片

slices包排序函数对比分析：

1. Sort() vs SortFunc() vs SortStableFunc():
   - Sort(): 使用默认排序规则（字典序/数值序）
   - SortFunc(): 使用自定义比较函数，不保证稳定性
   - SortStableFunc(): 使用自定义比较函数，保证稳定性（相等元素相对位置不变）

2. Sorted() vs SortedFunc() vs SortedStableFunc():
   - Sorted(): 检查是否按默认规则排序
   - SortedFunc(): 使用自定义函数检查是否排序
   - SortedStableFunc(): 使用自定义函数检查是否稳定排序

3. IsSorted():
   - 检查切片是否已按升序排列（使用默认比较规则）

cmp包说明：
- cmp是"compare"的缩写，提供标准化的比较操作
- cmp.Compare(a, b)返回int值：-1(a<b), 0(a==b), +1(a>b)
- 提供类型安全的比较操作，避免手动编写比较逻辑
*/
