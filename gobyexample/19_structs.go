package main

import "fmt"

/*
结构体


Go 语言中的结构体是字段的类型化集合，用于将数据组合在一起形成记录。
在初始化结构体时可以命名字段，未指定的字段将被赋予零值。
使用 & 前缀获取结构体的指针，点（.）操作符访问结构体字段。
结构体是可变的，可以封装在构造函数中创建。
对于单个值，可以使用匿名结构体类型，这在表格驱动的测试中很常见。
*/

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
