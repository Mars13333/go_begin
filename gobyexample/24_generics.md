# Go 泛型学习笔记

## 1. 泛型简介

Go 语言的泛型特性从 1.18 版本开始引入，极大地增强了语言的表达能力，解决了之前版本中需要针对不同类型重复实现相同算法逻辑的问题。

泛型允许在编写代码时不指定具体的数据类型，而是在使用时再确定具体类型，从而编写出更加通用和可重用的代码，避免了重复代码的出现，提高了代码的可维护性和灵活性。

## 2. 泛型参数列表语法

### 2.1 基本语法
- 使用 `[]` 包裹类型参数列表
- 放在函数名或类型名之后，参数列表之前
- 类型参数列表只包含类型参数的声明和约束
- 多个类型参数用逗号分隔

### 2.2 声明顺序
- 不需要按使用顺序声明
- Go 编译器会先扫描整个类型参数列表
- 收集所有类型参数的声明和约束
- 然后才进行类型检查

例如，以下三种写法是等价的：
```go
func SlicesIndex[S ~[]E, E comparable](s S, v E) int
func SlicesIndex[E comparable, S ~[]E](s S, v E) int
func SlicesIndex[E comparable, S ~[]E, F any](s S, v E) int
```

## 3. 类型约束

### 3.1 内置约束
- `any`：等同于 `interface{}`，表示任何类型
- `comparable`：表示可比较的类型（支持 `==` 和 `!=` 操作）
- `~T`：表示底层类型是 T 的类型

### 3.2 接口约束
- 可以使用任何接口作为约束
- 包括标准库中的接口
- 可以自定义接口作为约束

### 3.3 类型集接口
在 Go 1.18 之后，接口可以包含两种内容：
1. 方法声明（传统接口）
2. 类型约束（类型集接口）

例如：
```go
// 传统接口：只包含方法
type Stringer interface {
    String() string
}

// 类型集接口：只包含类型约束
type Ordered interface {
    ~int | ~float64 | ~string
}

// 混合接口：同时包含方法和类型约束
type StringableNumber interface {
    ~int | ~float64
    String() string
}
```

### 3.4 联合类型约束
- 使用 `|` 运算符组合多个类型
- 例如：`type Number interface { ~int | ~float64 }`

## 4. 常见约束示例

### 4.1 基本类型约束
```go
type Number interface {
    ~int | ~int32 | ~int64 | ~float32 | ~float64
}

type StringType interface {
    ~string
}
```

### 4.2 容器类型约束
```go
type SliceType interface {
    ~[]int | ~[]string
}

type MapType interface {
    ~map[string]int | ~map[int]string
}

type ChanType interface {
    ~chan int | ~chan string
}
```

### 4.3 标准库接口约束
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Stringer interface {
    String() string
}
```

## 5. 泛型使用示例

### 5.1 泛型函数
```go
func Min[T Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

func Sum[T Number](nums []T) T {
    var sum T
    for _, n := range nums {
        sum += n
    }
    return sum
}
```

### 5.2 泛型类型
```go
type Container[T any] struct {
    value T
}

func (c *Container[T]) Set(v T) {
    c.value = v
}

func (c *Container[T]) Get() T {
    return c.value
}
```

### 5.3 泛型方法
```go
type List[T any] struct {
    head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
    // 实现代码
}

func (lst *List[T]) AllElements() []T {
    // 实现代码
}
```

## 6. 注意事项

1. **约束的组合**：
   - 可以使用 `|` 组合多个类型
   - 可以使用 `&` 组合多个接口
   - 可以嵌套使用约束

2. **约束的限制**：
   - 不能使用 `interface{}` 作为约束（应该使用 `any`）
   - 不能使用未导出的类型作为约束
   - 不能使用类型别名作为约束

3. **类型推断**：
   - 在调用泛型函数时，通常可以依赖类型推断
   - 也可以显式指定类型参数

4. **性能考虑**：
   - 泛型代码在编译时会生成具体的类型实现
   - 不会带来运行时性能开销
   - 但可能会增加编译时间和二进制文件大小

## 7. 最佳实践

1. 优先使用类型推断，只在必要时显式指定类型参数
2. 合理使用约束，避免过度约束或约束不足
3. 保持代码简洁，避免过度使用泛型
4. 注意文档和注释，说明泛型函数的使用方式和约束条件
5. 考虑向后兼容性，避免破坏现有代码 