package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("unint:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	/*
		make([][]int, 3) 创建了一个长度为 3 的二维切片。
		外层循环初始化每个子切片的长度。
		内层循环填充每个子切片的值。
		最终结果是一个不规则的二维切片，每个子切片的长度不同。

	*/
	towD := make([][]int, 3)
	for i := range 3 {
		innerlen := i + 1
		towD[i] = make([]int, innerlen)
		for j := range innerlen {
			towD[i][j] = i + j
		}
	}
	fmt.Println("2d:", towD)
}
