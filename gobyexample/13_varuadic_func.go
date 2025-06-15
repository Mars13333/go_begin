package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	/*
		sum(nums...) 是一种语法，用于将一个切片（slice）的元素作为可变参数（variadic arguments）传递给函数。
		这种语法允许你将一个切片中的所有元素展开为独立的参数，传递给支持可变参数的函数。
	*/
}
