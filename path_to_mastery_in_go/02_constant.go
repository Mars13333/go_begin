package main

import "fmt"

/*
Go常量在声明时并不显式指定类型，也就是说可以使用 无类型常量（untyped constant）。
*/

const (
	SeekStart   = 0
	SeekCurrent = 1
	SeekEnd     = 2
)


/*
有类型常量的烦恼
Go是对类型安全要求十分严格的编程语言。Go要求，两个类型即便拥有相同的底层类型，
也仍然是不同的数据类型，不可以被相互比较或混在一个表达式中进行运算。
*/

type myInt int

// func test1(){
// 	var a int =5
// 	var b myInt = 6
// 	fmt.Println(a+b) // invalid operation: a + b (mismatched types int and myInt)
// }

/*
Go在处理不同类型的变量间的运算时不支持隐式的类型转换。
Go的设计者认为，隐式转换 带来的便利性不足以抵消其带来的诸多问题。要解决上面的编译错误，必须进行显式类型转换
*/


func test2(){
	var a int =5
	var b myInt = 6
	fmt.Println(a+int(b)) // 显示类型转换
}

/*
而将有类型常量与变量混合在一起进行运算求值时也要遵循这一要求，
即如果有类型常量与变量的类型不同，那么混合运算的求值操作会报错
*/


/*
无类型常量是Go语言推荐的实践，它拥有和字面值一样的灵活特性，
可以直接用于更多的表达式而不需要进行显式类型转换，从而简化了代码编写。
就是 无类型常量 提开发人员做了显示类型转换。
*/


/*
iota实现枚举常量
_ 可以跳过
*/

const(
	_=iota
	F_A	//1
	F_B //2
	_
	F_C //4
)

func main(){
	// test1()
	test2()
	// fmt.Println(F_A)
	fmt.Println(F_B)
	fmt.Println(F_C)
}

