package main

import "fmt"

/*
零值

Go语言中的每个原生类型都有其默认值，这个默认值就是这个类型的零值.

所有整型类型：0
浮点类型：0.0
布尔类型：false
字符串类型：""
指针、interface、切片（slice）​、channel、map、function：nil
另外，Go的零值初始是递归的，即数组、结构体等类型的零值初始化就是对其组成元素逐一进行零值初始化。

*/

/*
零值可用

按传统的思维，对于值为nil的变量，我们要先为其赋上合理的值后才能使用。
但Go从诞生以来就一直秉承着尽量保持“零值可用”的理念。
*/

func main1() {
	var zeroSlice []int // 没有初始化赋值， 但是可以正常使用。 即零值可用。
	zeroSlice = append(zeroSlice, 1)
	zeroSlice = append(zeroSlice, 2)
	zeroSlice = append(zeroSlice, 3)
	fmt.Println(zeroSlice)
}

/*
不过Go并非所有类型都是零值可用的！并且零值可用也有一定的限制！
*/

func main2() {
	var s []int
	//s[0]=12 // panic: runtime error: index out of range [0] with length 0
	s = append(s, 12) // success
	fmt.Println(s)
}

/*
另外，map也不支持零值可用
另外，零值可用的类型尽量避免复制
*/

func main() {
	var m map[string]int
	m["go"] = 1 // panic: assignment to entry in nil map

	// m2:make(map[string]int)
	// m2["go"]=1
}
