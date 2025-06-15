/*
Go 语言中结构体和接口的嵌入（embedding）特性，这是一种表达类型无缝组合的方式。
*/

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Printf("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	// 放在main的外面也是可以的。 只要接口方法一致，都可以算成一个类型
	type describer interface {
		describe() string
	}

	// cannot use co (variable of type container) as describer value in variable declaration:
	// container does not implement describer (missing method decribe)
	// describe interface 名字打错
	var d describer = co
	fmt.Println("describer:", d.describe())
}

// package main

// import "fmt"

// type base struct {
// 	num int
// }

// func (b base) describe() string {
// 	return fmt.Sprintf("base with num=%v", b.num)
// }

// type container struct {
// 	base
// 	str string
// }

// func main() {

// 	co := container{
// 		base: base{
// 			num: 1,
// 		},
// 		str: "some name",
// 	}

// 	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

// 	fmt.Println("also num:", co.base.num)

// 	fmt.Println("describe:", co.describe())

// 	type describer interface {
// 		describe() string
// 	}

// 	var d describer = co
// 	fmt.Println("describer:", d.describe())
// }
