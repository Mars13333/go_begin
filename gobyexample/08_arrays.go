package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	/*
		省略号 ... 表示让编译器根据初始化列表中的元素数量自动推断数组的长度。
		在这个例子中，初始化列表中有 5 个元素 {1, 2, 3, 4, 5}，因此编译器会推断数组的长度为 5。
		这种方式更加灵活，因为不需要显式指定数组的长度。
	*/
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)

}
