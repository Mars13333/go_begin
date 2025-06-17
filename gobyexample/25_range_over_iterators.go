/*
Go 语言版本 1.23 引入的迭代器特性，它扩展了 for-range 的适用范围，
使得开发者可以对几乎任何类型的数据结构进行遍历，从而提高了代码的灵活性和可读性。

这些新特性让开发者能够对几乎任何类型的数据结构进行遍历，提高了代码的灵活性和可读性，
使自定义数据结构的遍历变得更加简洁直观。

使用 iter.Seq[T] 类型创建自定义迭代器，扩展了 for-range 循环的适用范围

通过泛型实现的链表结构及其迭代方法

为自定义数据结构实现迭代器接口，使其可以直接用于 for-range 循环

以及使用 slices.Collect() 将迭代器收集为切片
*/

package main

import(
	"fmt"
	"iter"
	"slices"
)

// 模拟一个链表，泛型类型参数T可以是任何类型
type List[T any] struct {
	head, tail *element[T]
}

// 链表的节点，泛型类型参数T可以是任何类型
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

/*
yield 在 Go 的迭代器实现中是一个关键概念。
yield 可以翻译为"产出"或"让渡"。为了便于理解，你可以把它读作"产出元素"或"提交元素"。
yield 是一个函数参数，它接收当前元素的值并返回一个布尔值，表示是否继续迭代。

具体工作方式：
iter.Seq[T] 是一个函数类型（这里作为函数的返回类型），它接收一个 yield 函数作为参数
当你实现 iter.Seq[T] 时，你需要调用这个 yield 函数来"产出"每个元素
yield 函数返回 false 时表示迭代应该停止（例如当使用 break 时）


在迭代器上下文中，yield 函数的作用是：
把当前元素"产出"给外部的 for-range 循环
判断是否应该继续迭代（通过返回值）
每次调用 yield(e.val) 时，就相当于把 e.val 这个值提交给了 for-range 循环使用。
如果 for-range 循环决定继续，yield 返回 true；
如果 for-range 循环决定停止（比如遇到 break），yield 返回 false。

iter.Seq[T] 本身是一个函数类型，它的定义大致是：
type Seq[T any] func(yield func(T) bool)
所以当你实现一个返回 iter.Seq[T] 的函数时，
你必须返回一个符合这个签名的函数。
这就是为什么函数体必须是 return func(yield func(T) bool) { ... }。


这种模式在 Go 中被称为"函数返回函数"或"高阶函数"，
是函数式编程的一种常见模式。
在迭代器的情况下，你返回的是一个封装了迭代逻辑的函数，
这个函数会在 for-range 循环中被调用。

*/

/*
所有这三处的 T 都是同一个类型参数，它们共享相同的类型。如果 List[int] 调用 All()，那么所有的 T 都会是 int 类型。
这里没有声明新的类型参数，而是使用了 List[T] 结构体已经声明的类型参数 T。真正的类型参数声明在 type List[T any] 那一行。
*/

/*
这里返回的是一个函数，而不是布尔值。这个返回的函数本身没有返回值（是 void 函数）。
关于 yield 函数的返回值：
yield 是作为参数传入的函数，它返回布尔值
当调用 yield(e.val) 时，会得到一个布尔值
如果这个布尔值是 false，则执行 return，结束迭代
如果是 true，则继续循环
在 for-range 循环中使用时，Go 运行时会提供 yield 函数的实现，并在适当的时候返回 true 或 false。
例如，当循环中使用 break 时，yield 会返回 false。
所以，yield 函数的返回值是由调用者（for-range 循环）决定的，
而不是由 All() 方法决定的。
All() 方法只是根据 yield 的返回值决定是否继续迭代。
*/
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}


func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
