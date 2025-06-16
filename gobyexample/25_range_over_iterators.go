package main

// 模拟一个链表，泛型类型参数T可以是任何类型
type List[T any] struct {
	head, tail *node[T]
}

// 链表的节点，泛型类型参数T可以是任何类型
type node[T any] struct {
	next *node[T]
	val  T
}
