/*
泛型

Go 语言的泛型特性从 1.18 版本开始引入，极大地增强了语言的表达能力，解决了之前版本中需要针对不同类型重复实现相同算法逻辑的问题。
泛型允许在编写代码时不指定具体的数据类型，而是在使用时再确定具体类型，
从而编写出更加通用和可重用的代码，避免了重复代码的出现，提高了代码的可维护性和灵活性。

可以像对常规类型一样在泛型类型上定义方法，但我们必须保留类型参数。类型是 List[T]，而不是 List。

在调用泛型函数时，我们通常可以依赖类型推断。
*/

package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo", SlicesIndex(s, "zoo"))

	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
