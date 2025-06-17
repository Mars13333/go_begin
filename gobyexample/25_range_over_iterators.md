# Go 1.23 迭代器特性

Go 语言版本 1.23 引入的迭代器特性，它扩展了 for-range 的适用范围，使得开发者可以对几乎任何类型的数据结构进行遍历，从而提高了代码的灵活性和可读性。

## 核心概念

### iter.Seq[T] 类型

`iter.Seq[T]` 是一个函数类型，它的定义大致是：

```go
type Seq[T any] func(yield func(T) bool)
```

它接收一个 `yield` 函数作为参数，用于产出元素并控制迭代流程。

### yield 函数

`yield` 函数是迭代器实现的核心，可以理解为"产出"或"提交"元素：

1. **功能**：接收当前元素的值并返回一个布尔值，表示是否继续迭代
2. **工作方式**：
   - 把当前元素"产出"给外部的 for-range 循环
   - 通过返回值判断是否应该继续迭代
3. **返回值**：
   - `true`：继续迭代
   - `false`：停止迭代（例如当使用 `break` 时）

### 实现迭代器

实现一个返回 `iter.Seq[T]` 的函数时，必须返回一个符合这个签名的函数：

```go
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// 迭代逻辑
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}
```

这种模式在 Go 中被称为"函数返回函数"或"高阶函数"，是函数式编程的常见模式。

### 泛型参数 T

在 `func (lst *List[T]) All() iter.Seq[T]` 中的三个 `T` 都是同一个类型参数：

1. `(lst *List[T])` 中的 `T`：接收者类型中使用的泛型参数
2. `iter.Seq[T]` 中的 `T`：返回类型中的泛型参数
3. `func(yield func(T) bool)` 中的 `T`：返回的函数中使用的泛型参数

这里没有声明新的类型参数，而是使用了 `List[T]` 结构体已经声明的类型参数 `T`。

## 使用迭代器

### 基本用法

```go
// 在for-range循环中使用迭代器
for e := range lst.All() {
	fmt.Println(e)
}
```

### 收集迭代器结果

```go
// 使用slices.Collect将迭代器收集为切片
all := slices.Collect(lst.All())
```

### 无限序列

可以创建无限序列的迭代器，在使用时通过条件控制何时停止：

```go
// 斐波那契数列迭代器
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

// 使用时通过break控制何时停止
for n := range genFib() {
	if n >= 10 {
		break
	}
	fmt.Println(n)
}
```

#### 迭代器停止机制

迭代器的停止机制展示了生成逻辑和终止条件的分离：

1. 迭代器函数（如 `genFib()`）本身不知道何时停止，它只负责生成值序列
2. 停止条件（如 `if n >= 10`）是在**使用**迭代器的循环中定义的，而不是在迭代器内部
3. 当循环中执行 `break` 时，Go 运行时会让 `yield` 函数返回 `false`
4. 当 `yield(a)` 返回 `false` 时，迭代器内部的 `if !yield(a)` 条件成立，执行 `return` 语句，结束迭代器函数

这种设计的优点是：
- 迭代器专注于值的生成逻辑
- 使用者决定何时停止迭代
- 同一个迭代器可以用于不同的终止条件

这就是迭代器模式的优雅之处：生成逻辑和终止条件可以分离，迭代器专注于生成值，使用者决定何时停止。

## 迭代器的优势

1. 可以对几乎任何类型的数据结构进行遍历
2. 提高了代码的灵活性和可读性
3. 使自定义数据结构的遍历变得更加简洁直观
4. 支持惰性计算和无限序列 