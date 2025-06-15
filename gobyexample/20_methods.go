package main

import "fmt"

/*
Go 语言中的方法（methods）及其接收者（receiver）的概念。
方法是一种特殊的函数，它与某个类型关联，可以访问和修改该类型的字段。
可以是值类型，可以说指针类型。
使用指针接收者类型可以避免方法调用时的复制，并允许方法修改接收的结构体。

Go 语言会自动处理方法调用中值和指针之间的转换。简写*pt.name为pt.name即可。
*/

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perim()) // 周长perimeter命名简写为perim

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perim())
}
