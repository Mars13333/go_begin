package main

import "fmt"

/*
map原理

map是Go语言提供的一种抽象数据类型，它表示一组无序的键值对（key-value)。

map不支持 “零值可用”。未显式赋初值的map类型变量的零值为nil。对处于零值状态的map变量进行操作将会导致运行时panic。

创建map类型有俩种方式：
1.复合字面值
2.make
*/

// 复合字面值
func main1() {
	var statusText = map[int]string{
		1: "OK",
		2: "Created",
		3: "Accepted",
	}

	fmt.Println(statusText)
}

// make
func main2() {
	m1 := make(map[int]string)
	m1[11] = "lala"
	m1[22] = "haha"
	fmt.Println(m1)
}

/*
和切片一样，map也是引用类型，将map类型变量作为函数参数传入不会有很大的性能损耗，
并且在函数内部对map变量的修改在函数外部也是可见的。
*/

/*
map常用操作
*/
func main() {
	m := make(map[string]int)
	m["key1"] = 111 // 1.插入数据
	m["key2"] = 222
	m["key3"] = 333
	fmt.Println(m)
	fmt.Println(len(m)) // len(m) 获取数据个数

	// 2.查找 ，使用"comma ok"惯用法
	// 并不关心某个key对应的value，而仅仅关心某个key是否在map中，
	// 因此我们使用**空标识符**（blank identifier）忽略了可能返回的数据值，而仅关心ok的值是否为true
	_, ok := m["key13"]
	if !ok {
		fmt.Println("key13 not exist") // key13 not exist
	}

	// 3.读取
	v := m["key3"]
	fmt.Println(v)
	v2 := m["key13"]
	fmt.Println(v2) // 其实不存在key13的value, 但是打印了0，这个其实是一个合法的值0,这是value类型int的零值
	// 所以，仍然可以借助comma ok方法来获取
	v3, ok := m["key13"]
	if !ok {
		// key13 不存在
		fmt.Println("key13 not exist")
	} else {
		fmt.Println(v3)
	}

	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println()

	// 4.删除数据  
	m2:=map[string]int{
		"key1":1,
		"key2":2, // 最后一行也必须要有逗号
	}
	fmt.Println(m2)
	delete(m2,"key2")
	delete(m2,"key22") //注意，即使要删除的不存在，delete也不会导致panic
	fmt.Println(m2)


	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println()

	// 5.遍历数据 可以像对待切片那样通过for range进行遍历
	// map遍历次序是随机的。如果需要稳定的次序，可以把key存储到切片中
	m3:=map[int]int{
		1:11,
		2:12,
		3:13,
		4:14,
		5:15,
		6:16,
	}
	for k,v:=range m3{
		fmt.Printf("[%d,%d]\n",k,v)
	}
}


/*
map实现原理

和切片相比，map类型的内部实现要复杂得多。Go运行时使用一张哈希表来实现抽象的map类型。
运行时实现了map操作的所有功能，包括查找、插入、删除、遍历等。
在编译阶段，Go编译器会将语法层面的map操作重写成运行时对应的函数调用。下面是大致的对应关系

与语法层面map类型变量一一对应的是runtime.hmap类型的实例。
hmap是map类型的header，可以理解为map类型的描述符，它存储了后续map类型操作所需的所有信息

详见： 06_data_map.md !!!
*/

