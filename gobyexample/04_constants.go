package main

import (
	"fmt"
	"math"
)

const (
	s string = "constant"
)

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

/*
3e20 是一种科学计数法的表示方式，用于表示一个非常大的数字。
具体来说，3e20 表示 3×10^20，即 3 乘以 10 的 20 次方。

3e20 表示 3×10^20
2.5e-3 表示 2.5×10^−3


科学计数法
科学计数法是一种表示非常大或非常小的数字的简洁方式。它的格式是：
a×10 ^ b
*/
