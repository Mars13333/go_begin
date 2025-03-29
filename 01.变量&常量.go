package main

// 在go语言中 采用的是后置类型的声明方式： <命名> <类型>
// 通常使用关键字var来声明变量

var a int
var b int = 1
var c = 1

// 多个变量同时声明
var (
	x, y int
	z    float64
)

// 如果有初始值可以简短声明 使用:=
func main() {
	a := 1
	b := int32(1)
	println(a)
	println(b)
}

// A 常量
// 定义时候必须指定值
// 指定的值的类型主要有三类：布尔，数字，字符串，其中数字类型包含（rune,integer,float-point,complex）都属于基本数据类型
// 不能用:=
const A = 64
const (
	A1 = 4
	B1 = 0.2
)
