/*
closures 闭包

闭包是函数式编程中的一个概念，它允许一个函数访问其外部作用域的变量。
在 Go 语言中，匿名函数可以捕获周围作用域的变量，从而形成闭包。

这里描述了一个名为 intseq 的函数，它返回另一个在 intseq 函数体内匿名定义的函数。
返回的匿名函数捕获了变量 i，形成了一个闭包。


在 Go 中，通过匿名函数可以创建闭包。
闭包可以用于封装状态，例如计数器或可变配置。
每次调用闭包时，它都会访问和修改其捕获的变量，从而保持状态。
*/

package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq() // 返回的是函数，所以下面调用页使用小括号，如果不使用将打印地址值

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
	fmt.Println(newInts)   // 0xdb8420
	fmt.Println(newInts()) // 1
}
