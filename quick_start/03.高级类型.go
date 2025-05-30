package main

import (
	"fmt"
	"sync"
)

func main() {
	//数组()
	//切片()
	//切片2()
	//切片3()
	//maps()
	//maps2()
	//maps3()
	//maps4()
}

func 数组() {
	/*
		定义：由若干相同类型的元素组成的序列
		数组的长度是固定的，声明后无法改变
		数组的长度是数组类型的一部分，eg：元素类型相同但是长度不同的两个数组是不同类型的
		需要严格控制程序所使用内存时，数组十分有用，因为其长度固定，避免了内存二次分配操作
	*/
	// 定义长度为5的数组
	var arr1 [5]int
	for i := 0; i < 5; i++ {
		arr1[i] = i
	}
	printHelper("arr1", arr1)

	// 以下赋值会报类型不匹配错误，因为数组长度是数组类型的一部分
	//arr1 = [3]int{1, 2, 3}
	arr1 = [5]int{2, 3, 4, 5, 6} // 长度和元素类型都相同，才可以正确赋值

	// 简写模式，在定义的同时给出赋值
	arr2 := [5]int{0, 1, 2, 3, 4}
	printHelper("arr2", arr2)

	// 数组元素类型相同并且数组长度相同的情况下，数组可以进行比较
	fmt.Println(arr1 == arr2)

	// 也可以不显式的定义数组长度，由编译器完成长度计算
	var arr3 = [...]int{0, 1, 2, 3, 4}
	printHelper("arr3", arr3)

	// 定义前四个元素为默认值0,最后一个元素为-1
	var arr4 = [...]int{4: -1}
	printHelper("arr4", arr4)

	// 多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD: ", twoD)
}

func printHelper(name string, arr [5]int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v[%v]: %v\n", name, i, arr[i])
	}

	// len 获取长度
	fmt.Printf("len of %v: %v\n", name, len(arr))

	// cap 也可以用来获取数组长度
	fmt.Printf("cap of %v: %v\n", name, cap(arr))

	fmt.Println()
}

func 切片() {
	/*
		切片组成要素：
		指针：指向底层数组
		长度：切片中元素的长度，不能大于容量
		容量：指针所指向的底层数组的总容量
	*/

	// 使用make初始化
	slice1 := make([]int, 5) //初始化长度和容量为5的切片
	fmt.Println("slice1", slice1)
	slice2 := make([]int, 6, 10) // 初始化长度为5,容量为10的切片
	fmt.Println("slice2", slice2)

	// 使用简短定义
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice3", slice3)

	// 使用数组来初始化切片
	arr := [5]int{1, 2, 3, 4, 5}
	slice4 := arr[0:3] // 左闭右开区间，最终切片为[1,2,3]
	fmt.Println("slice4", slice4)

	// 使用切片来初始切片
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := sliceA[0:3] // 左闭右开区间，最终切片为[1,2,3]
	fmt.Println("sliceB", sliceB)

}

func 切片2() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println("len: ", len(slice))
	fmt.Println("cap: ", cap(slice))

	// 改变切片长度
	slice = append(slice, 6)
	fmt.Println("after append operation: ")
	fmt.Println(slice)
	fmt.Println("len: ", len(slice))
	fmt.Println("cap: ", cap(slice)) // 注意，底层数组容量不够时，会重新分配数组空间，通常为俩倍
}

func 切片3() {
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[0:3]
	// 使用 copy 方法可以避免共享同一个底层数组
	anotherSlice := make([]int, len(slice))
	copy(anotherSlice, slice)
	fmt.Println("before modifying underlying array:")
	fmt.Println("slice: ", slice)
	fmt.Println("newSlice: ", newSlice)
	fmt.Println("anotherSlice: ", anotherSlice)
	fmt.Println()

	newSlice[0] = 6 // 多个切片共享一个底层数组的情况
	fmt.Println("after modifying underlying array:")
	fmt.Println("slice: ", slice)
	fmt.Println("newSlice: ", newSlice)
	fmt.Println("anotherSlice: ", anotherSlice)
}

func maps() {
	/*
		在 Go 语言里面，map 一种无序的键值对, 它是数据结构 hash 表的一种实现方式
		使用关键字 map[keyType]valueType
		注意：
			必须指定 key, value 的类型，插入的纪录类型必须匹配。
			key 具有唯一性，插入纪录的 key 不能重复。
			KeyType 可以为基础数据类型（例如 bool, 数字类型，字符串）, 不能为数组，切片，map，它的取值必须是能够使用 == 进行比较。
			ValueType 可以为任意类型。
			无序性。
			线程不安全, 一个 goroutine 在对 map 进行写的时候，另外的 goroutine 不能进行读和写操作，Go 1.6 版本以后会抛出 runtime 错误信息
	*/

	// 使用var声明
	var cMap map[string]int //只定义，此时cMap为nil
	fmt.Println(cMap == nil)
	//cMap["北京"] = 1 //会报错，因为cMap为nil

	// 使用make
	cMap2 := make(map[string]int)
	cMap2["北京"] = 1
	fmt.Println("cMap2: ", cMap2)

	// 指定初始容量
	//说明：在使用 make 初始化 map 的时候，可以指定初始容量，这在能预估 map key 数量的情况下，减少动态分配的次数，从而提升性能。
	cMap2 = make(map[string]int, 10)
	cMap2["北京"] = 1
	fmt.Println("cMap2: ", cMap2)

	// 简短声明方式
	cMap3 := map[string]int{"北京": 111}
	fmt.Println("cMap3: ", cMap3)

	// 一些基本操作
	cMap3["北京"] = 1 //写

	//code := cMap3["北京"] // 读
	//fmt.Println(code)

	//code = cMap3["广州"] // 读不存在 key
	//fmt.Println(code)

	//多值返回的机制
	//Go 的 map 查询会返回两个值：
	//第一个值：键对应的值（不存在时返回零值）。
	//第二个值（ok）：布尔类型，表示键是否存在。
	code, ok := cMap3["北京"] // 检查 key 是否存在
	if ok {
		fmt.Println(code)
	} else {
		fmt.Println("key not exist")
	}

	//delete(cMap3, "北京") // 删除 key
	//fmt.Println("北京")
}

func maps2() {
	// map的循环和无序性
	cMap := map[string]int{"北京": 1, "上海": 2, "广州": 3, "深圳": 4}
	for city, code := range cMap {
		fmt.Printf("%s:%d", city, code)
		fmt.Println()
	}
}

func maps3() {
	// 线程不安全
	cMap := make(map[string]int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		cMap["北京"] = 1
		wg.Done()
	}()

	go func() {
		cMap["上海"] = 2
		wg.Done()
	}()

	wg.Wait()

	// 在 Go 1.6 之后的版本，多次运行此段代码，你将遇到这样的错误信息：
	// 解决之道: 1.对读写枷锁；2.使用security map,例如 sync.map
}

func maps4() {
	// map嵌套
	provinces := make(map[string]map[string]int)

	provinces["北京"] = map[string]int{
		"东城区": 1,
		"西城区": 2,
		"朝阳区": 3,
		"海淀区": 4,
	}

	fmt.Println(provinces["北京"])
}

/*
自定义类型
通过自定义类型来表示一些特殊的数据结构和业务逻辑。
使用关键字 type 来声明：type NAME TYPE

单次声明
type City string

批量声明
type (
    B0 = int8
    B1 = int16
    B2 = int32
    B3 = int64
)

type (
    A0 int8
    A1 int16
    A2 int32
    A3 int64
)


package main

import "fmt"

type Age int
type Height int

func main() {
    age := Age(12)
    height := Height(175)

    fmt.Println(height / age)
}

修复方法使用显式转换:
fmt.Println(int(height) / int(age))
*/

/*
结构体
数组、切片和 Map 可以用来表示同一种数据类型的集合，但是当我们要表示不同数据类型的集合时就需要用到结构体。
结构体是由零个或多个任意类型的值聚合成的实体

关键字 type 和 struct 用来定义结构体：
package main

import "fmt"

type Student struct {
    Age     int
    Name    string
}

func main() {
    stu := Student{
        Age:     18,
        Name:    "name",
    }
    fmt.Println(stu)

    // 在赋值的时候，字段名可以忽略
    fmt.Println(Student{20, "new name"})

    return
}

---

通常结构体中一个字段占一行，但是类型相同的字段，也可以放在同一行，例如：
type Student struct{
    Age           int
    Name, Address string
}

---

结构体中的字段如果是小写字母开头，那么其他 package 就无法直接使用该字段，例如：
// 在包 pk1 中定义 Student 结构体
package pk1
type Student struct{
    Age  int
    name string
}

// 在另外一个包 pk2 中调用 Student 结构体
package pk2

func main(){
    stu := Student{}
    stu.Age = 18        //正确
    stu.name = "name"  // 错误，因为`name` 字段为小写字母开头，不对外暴露
}

---

结构体中可以内嵌结构体 但是需要注意的是：如果嵌入的结构体是本身，那么只能用指针！！！
package main

import "fmt"

type Tree struct {
    value       int
    left, right *Tree
}

func main() {
    tree := Tree{
        value: 1,
        left: &Tree{
            value: 1,
            left:  nil,
            right: nil,
        },
        right: &Tree{
            value: 2,
            left:  nil,
            right: nil,
        },
    }

    fmt.Printf(">>> %#v\n", tree)
}

---

结构体是可以比较的 前提是结构体中的字段类型是可以比较的
package main

import "fmt"

type Tree struct {
    value       int
    left, right *Tree
}

func main() {
    tree1 := Tree{
        value: 2,
    }

    tree2 := Tree{
        value: 1,
    }

    fmt.Printf(">>> %#v\n", tree1 == tree2)
}


---


结构体内嵌匿名成员
声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫 匿名成员
package main

import "fmt"

type Person struct {
    Age  int
    Name string
}

type Student struct {
    Person
}

func main() {
    per := Person{
        Age:  18,
        Name: "name",
    }

    stu := Student{Person: per}

    fmt.Println("stu.Age: ", stu.Age)
    fmt.Println("stu.Name: ", stu.Name)
}
*/

/*
函数

函数是语句序列的集合，能够将一个大的工作分解为小的任务，对外隐藏了实现细节
函数组成：
	函数名
	参数列表(parameter-list)
	返回值(result-list)
	函数体(body)

func name(parameter-list) (result-list){
    body
}

---

单返回值函数
func plus(a, b int) (res int){
    return a + b
}

---

多返回值函数
func multi()(string, int){
    return "name", 18
}


---

命名返回值
// 被命名的返回参数的值为该类型的默认零值
// 该例子中 name 默认初始化为空字符串，height 默认初始化为 0
func namedReturnValue()(name string, height int){
    name = "xiaoming"
    height = 180
    return
}

---

参数可变函数
func sum(nums ...int)int{
    fmt.Println("len of nums is : ", len(nums))
    res := 0
    for _, v := range nums{
        res += v
    }
    return res
}

func main(){
    fmt.Println(sum(1))
    fmt.Println(sum(1,2))
    fmt.Println(sum(1,2,3))
}

---

匿名函数
func main(){
    func(name string){
       fmt.Println(name)
    }("禾木课堂")
}


---

闭包
func main() {
    addOne := addInt(1)
    fmt.Println(addOne())
    fmt.Println(addOne())
    fmt.Println(addOne())

    addTwo := addInt(2)
    fmt.Println(addTwo())
    fmt.Println(addTwo())
    fmt.Println(addTwo())
}

func addInt(n int) func() int {
    i := 0
    return func() int {
        i += n
        return i
    }
}

---

函数作为参数
func sayHello(name string) {
    fmt.Println("Hello ", name)
}

func logger(f func(string), name string) {
    fmt.Println("start calling method sayHello")
    f(name)
    fmt.Println("end calling method sayHellog")
}

func main() {
    logger(sayHello, "禾木课堂")
}

---

传值和传引用
func sendValue(name string) {
    name = "hemuketang"
}

func sendAddress(name *string) {
    *name = "hemuketang"
}

func main() {
    // 传值和传引用
    str := "禾木课堂"
    fmt.Println("before calling sendValue, str : ", str)
    sendValue(str)
    fmt.Println("after calling sendValue, str : ", str)

    fmt.Println("before calling sendAddress, str : ", str)
    sendAddress(&str)
    fmt.Println("after calling sendAddress, str: ", str)
}
*/

/*
方法

方法主要源于 OOP 语言，在传统面向对象语言中 (例如 C++), 我们会用一个“类”来封装属于自己的数据和函数，这些类的函数就叫做方法。

虽然 Go 不是经典意义上的面向对象语言，但是我们可以在一些接收者（自定义类型，结构体）上定义函数，同理这些接收者的函数在 Go 里面也叫做方法。

方法（method）的声明和函数很相似, 只不过它必须指定接收者：
func (t T) F() {}

注意：
	接收者的类型只能为用关键字 type 定义的类型，例如自定义类型，结构体。
	同一个接收者的方法名不能重复 (没有重载)，如果是结构体，方法名还不能和字段名重复。
	值作为接收者无法修改其值，如果有更改需求，需要使用指针类型。

简单的例子：
package main

type T struct{}

func (t T) F()  {}

func main() {
    t := T{}
    t.F()
}

---

接收者可以同时为值和指针
在 Go 语言中，方法的接收者可以同时为值或者指针，例如：
package main

type T struct{}

func (T) F()  {}
func (*T) N() {}

func main() {
    t := T{}
    t.F()
    t.N()

    t1 := &T{} // 指针类型
    t1.F()
    t1.N()
}


*/

/*
接口

接口类型是一种抽象类型，是方法的集合，其他类型实现了这些方法就是实现了这个接口。

定义接口
type interface_name interface {
	method_name1 [return_type]
	method_name2 [return_type]
	method_name3 [return_type]
	...
	method_namen [return_type]
}

---


简单示例：打印不同几何图形的面积和周长
package main

import (
    "fmt"
    "math"
)

type geometry interface {
    area() float32
    perim() float32
}

type rect struct {
    len, wid float32
}

func (r rect) area() float32 {
    return r.len * r.wid
}

func (r rect) perim() float32 {
    return 2 * (r.len + r.wid)
}

type circle struct {
    radius float32
}

func (c circle) area() float32 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float32 {
    return 2 * math.Pi * c.radius
}

func show(name string, param interface{}) {
    switch param.(type) {
    case geometry:
        // 类型断言
        fmt.Printf("area of %v is %v \n", name, param.(geometry).area())
        fmt.Printf("perim of %v is %v \n", name, param.(geometry).perim())
    default:
        fmt.Println("wrong type!")
    }
}

func main() {
    rec := rect{
        len: 1,
        wid: 2,
    }
    show("rect", rec)

    cir := circle{
        radius: 1,
    }
    show("circle", cir)

    show("test", "test param")
}


---

接口中可以内嵌接口

对上述例子做以下修改：

首先添加 tmp 接口，该接口定义了 area() 方法
将 tmp 作为 geometry 接口中的匿名成员，并且将 geometry 接口中原本定义的 area() 方法删除
完成以上两步后，geometry 接口将会拥有 tmp 接口所定义的所有方法。运行结果和上述例子相同。

type tmp interface{
    area() float32
}

type geometry interface {
    // area() float32
    tmp
    perim() float32
}
*/
