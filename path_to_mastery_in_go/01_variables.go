// go是静态编译型语言，遵循使用变量前先声明，以下是go常见的变量声明形式

package main

import "fmt"

var a int32            // var 名称 类型
var s string = "hello" // var 名称 类型 赋值
var i = 13             // var 名称 赋值（省略类型）
var (                  //批量声明
	crlf       = []byte("\r\n")
	colonSpace = []byte(":")
)

func block() {
	//对于接受默认类型的变量
	n1 := 17 //简要声明（省略var,省略类型）
	f1 := 3.14
	s1 := "hello, gopher!"

	//对于不接受默认类型的变量，依然可以使用短变量声明形式，只是在“:=”右侧要进行显式转型
	// 因为即便是int也分为int32,int64; float也分不同的，看需要。
	n2 := int32(17)
	f2 := float32(3.14)
	s2 := []byte("hello, gopher!")
	fmt.Println(n1) // 注意：仅打印一个，go编译器会阻止程序运行
	fmt.Println(f1)
	fmt.Println(s1)
	fmt.Println(n2)
	fmt.Println(f2)
	fmt.Println(s2)
	// why here.
	// 因为简要声明无法作为包级别变量！上面var都是包级变量，只能用var, 并且推荐形式为：var variableName = InitExpression
	// 仅可作为函数或方法体内部的局部变量
}

func main(){
	block()
}


// 尽量使用短声明。 简明扼要， 见闻之意， 联系上下文， 不推荐一长串的方式。