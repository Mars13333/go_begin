/*
泛型

详见：24_generics.md

Go 语言的泛型特性从 1.18 版本开始引入，极大地增强了语言的表达能力，解决了之前版本中需要针对不同类型重复实现相同算法逻辑的问题。
泛型允许在编写代码时不指定具体的数据类型，而是在使用时再确定具体类型，
从而编写出更加通用和可重用的代码，避免了重复代码的出现，提高了代码的可维护性和灵活性。

可以像对常规类型一样在泛型类型上定义方法，但我们必须保留类型参数。类型是 List[T]，而不是 List。

在调用泛型函数时，我们通常可以依赖类型推断。
*/

/*

go的泛型参数列表的语法：
- 使用[]包裹类型参数列表
- 放在函数名或类型名之后，参数列表之前
- 类型参数列表只包含类型参数的声明和约束
- 约束如：any、comparable、~[]E
- 多个类型参数用逗号隔开
- 不需要按使用顺序声明，因为Go编译器会先扫描整个类型参数列表，收集所有的类型参数
的声明和约束，然后才进行类型检查，所以SlicesIndex等同于：
这三种写法都是等价的
func SlicesIndex[S ~[]E, E comparable](s S, v E) int
func SlicesIndex[E comparable, S ~[]E](s S, v E) int
func SlicesIndex[E comparable, S ~[]E, F any](s S, v E) int  // 即使有其他类型参数也可以


另外，
~本身不表示切片，它表示“底层类型是”
[]才表示切片类型
~[]组合在一起表示“底层类型是切片”

// 使用 ~ 约束底层类型是 int
type Number interface {
    ~int | ~int32 | ~int64
}

// 使用 ~ 约束底层类型是 string
type StringType interface {
    ~string
}

// 使用 ~ 约束底层类型是切片
type SliceType interface {
    ~[]int | ~[]string
}

*/

package main

import "fmt"

// func SlicesIndex[S ~[]E, E comparable](s S, v E) int
// 这种方式看着更顺眼，不亏顾虑E是否还未声明就使用的疑虑
func SlicesIndex[E comparable, S ~[]E](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// 模拟一个链表，泛型类型参数T可以是任何类型
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

	// 一般不需要像下面那样显示描述类型，go会自动推断
	fmt.Println("index of zoo", SlicesIndex(s, "zoo"))

	// _ = SlicesIndex[[]string, string](s, "zoo")
	// 因为上面定义换成了[E comparable, S ~[]E]形式，所以这里显示调用的话需要遵循类型参数的顺序
	_ = SlicesIndex[string, []string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
