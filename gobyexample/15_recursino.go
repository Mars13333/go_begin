/*
递归

栈溢出风险：对于较大的 n 值，递归调用可能会因为栈溢出而失败。
性能问题：递归函数可能会重复计算相同的值，导致性能下降。

为了解决这些问题，可以考虑使用迭代方法(for)或尾递归优化（尽管 Go 语言的编译器不直接支持尾递归优化）。
*/

package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// 匿名函数同样支持递归，但是需要变量接收
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

}
