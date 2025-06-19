package main

import (
	"fmt"
	"slices" // 导入Go 1.21+的slices包，提供切片操作功能
)

func main() {
	// 创建一个字符串切片用于排序演示
	str := []string{"c", "a", "b"}
	// 使用slices.Sort对字符串切片进行原地排序（按字典序）
	slices.Sort(str)
	fmt.Println("Strings:", str)

	// 创建一个整数切片用于排序演示
	ints := []int{7, 2, 4}
	// 使用slices.Sort对整数切片进行原地排序（按数值大小）
	slices.Sort(ints)
	fmt.Println("Ints:", ints)

	// 使用slices.IsSorted检查切片是否已经排序
	s := slices.IsSorted(ints)
	fmt.Println("Sorted:", s)
}

/*
Go语言排序示例 - 使用slices包进行切片排序

主旨：
1. 演示Go 1.21+引入的slices包的基本排序功能
2. 展示对字符串和整数切片的排序操作
3. 说明slices.Sort()是原地排序，会修改原切片!!!
4. 展示如何检查切片是否已排序

关键特性：
- slices.Sort()：对切片进行原地排序，支持多种数据类型
- slices.IsSorted()：检查切片是否已经按升序排列!!!
- 相比sort包，slices包提供了更简洁的API和更好的类型安全
*/
