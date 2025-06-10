package main

import "fmt"

/*
Go语言中的复合类型包括 结构体、数组、切片和map。
对于复合类型变量，最常见的值构造方式就是对其内部元素进行逐个赋值
*/

type myStruct struct {
	name string
	age  int
}

func test1() {
	var s myStruct
	s.name = "tony"
	s.age = 23

}

func test2() {
	var a [5]int
	a[0] = 13
	a[1] = 14
	a[2] = 15
	a[3] = 16
	a[4] = 17
}

func test3() {
	sl := make([]int, 5, 5)
	sl[0] = 23
	sl[1] = 24
	sl[2] = 25
	sl[3] = 26
	sl[4] = 27
}

func test4() {
	m := make(map[int]string)
	m[1] = "hello"
	m[2] = "gopher"
	m[3] = "!"

}

/*
复合字面值语法

但这样的值构造方式让代码显得有些烦琐，尤其是在构造组成较为复杂的复合类型变量的初值时。
Go提供的复合字面值（composite literal）语法可以作为复合类型变量的初值构造器。

复合字面值由俩部分组成，一部分是类型，一部分是由大括号包裹的字面值。
*/

func test5() {
	s := myStruct{"tony", 23}
	a := [5]int{13, 14, 15, 16, 17}
	sl := []int{23, 24, 25, 26, 27}
	m := map[int]string{1: "hello", 2: "gopher", 3: "!"}
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(sl)
	fmt.Println(m)
}

/*

虽然数组在 Go 中是基本类型，但在实际开发中，!!!!!!!数组的使用频率相对较低!!!!!!，主要原因如下
数组的长度是类型的一部分，一旦声明，长度不能改变。这在某些动态场景下不够灵活。相比之下，切片（[]Type）是动态的，可以动态扩展和收缩，因此更常用。

切片是基于数组的动态视图，提供了更多的灵活性和便利性。例如：
切片可以动态扩展，使用 append 函数可以方便地添加元素。
切片支持切片操作（如 slice[start:end]），可以方便地获取子切片。

尽管数组的使用频率较低，但在某些场景下，数组仍然是非常有用的：
固定长度的集合：当你知道集合的长度是固定的，并且不会改变时，数组是一个很好的选择。例如，表示一周的天数（7 个元素）或一个固定大小的缓存。
性能优化：数组的长度是固定的，因此在某些情况下，数组的性能可能比切片更好，尤其是在数组较小且频繁使用时。

数组

声明一个数组：var arrayName [length]Type
length：数组的长度，必须是一个常量。
Type：数组中每个元素的类型。
arrayName：数组的变量名。

var arr [5]int // 声明一个长度为 5 的整数数组

初始化
var arr [5]int = [5]int{1, 2, 3, 4, 5} // 显式初始化
var arr [5]int = [5]int{1, 2}         // 部分初始化，未指定的元素为零值
var arr [5]int = [5]int{1, 2, 3, 4, 5} // 省略类型声明
var arr = [5]int{1, 2, 3, 4, 5}        // 省略类型声明和长度
var arr [5]int // 零值是 [0 0 0 0 0]

*/
