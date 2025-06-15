package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// range 可以用于整数类型，它会生成从 0 到该整数减 1 的序列。
	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
		// break 结束循环
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
			// continue 跳过本次循环
		}
		fmt.Println(n)
	}

}
