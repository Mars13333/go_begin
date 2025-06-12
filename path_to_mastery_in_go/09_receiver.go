package main

import (
	"fmt"
	"reflect"
)

type T struct {
	a int
}

func (t T) M1() {
	t.a = 10
}

func (t *T) M2() {
	t.a = 11
}

func main1() {
	var t T // t.a = 0
	fmt.Println(t.a)
	t.M1()
	fmt.Println(t.a)
	t.M2()
	fmt.Println(t.a)
}

/*
0
0
11
M1和M2方法体内都对字段a做了修改，
但M1（采用值类型receiver）修改的只是实例的副本，对原实例并没有影响，因此M1调用后，输出t.a的值仍为0。
而M2（采用指针类型receiver）修改的是实例本身，因此M2调用后，t.a的值变为了11。
*/

/*
仅关乎调用

是不是T类型实例只能调用receiver为T类型的方法，不能调用receiver为*T类型的方法呢？答案是否定的。
无论是T类型实例还是*T类型实例，都既可以调用receiver为T类型的方法，
也可以调用receiver为*T类型的方法。
*/

/*
方法集合与接口调用

Go语言的一个创新是，自定义类型与接口之间的实现关系是松耦合的：
如果某个自定义类型T的方法集合是某个接口类型的方法集合的超集，那么就说类型T实现了该接口，
并且类型T的变量可以被赋值给该接口类型的变量，即我们说的方法集合决定接口实现。

要判断一个自定义类型是否实现了某接口类型，我们首先要识别出自定义类型的方法集合和接口类型的方法集合。
但有些时候它们并不明显，尤其是当存在结构体嵌入、接口嵌入和类型别名时。
这里我们实现了一个工具函数，它可以方便地输出一个自定义类型或接口类型的方法集合。
*/

func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elemTyp := v.Elem()

	n := elemTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty! \n", elemTyp)
		return
	}

	fmt.Printf("%s's method set:\n", elemTyp)
	for i := 0; i < n; i++ {
		fmt.Println("-", elemTyp.Method(i).Name)
	}
	fmt.Println("\n")
}

/*
接下来，用该工具函数输出示例中的接口类型和自定义类型的方法集合：
*/
type MyInterface interface {
	M3()
	M4()
}

type G struct{}

func (g G) M3()  {}
func (g *G) M4() {}

func main() {
	var g G
	var pt *G
	DumpMethodSet(&g)
	DumpMethodSet(&pt)
	DumpMethodSet((*MyInterface)(nil))
}

/*
main.G's method set:
- M3

*main.G's method set:
- M3
- M4

main.MyInterface's method set:
- M3
- M4

在上述输出结果中，T、*T和Interface各自的方法集合一目了然。
我们看到T类型的方法集合中只包含M1，**无法成为Interface类型的方法集合的超集**，
因此这就是本条开头例子中编译器认为变量t不能赋值给Interface类型变量的原因。
在输出的结果中，我们还看到*T类型的方法集合为\[M1, M2]​。*T类型没有直接实现M1，
但M1仍出现在了*T类型的方法集合中。这符合**Go语言规范：
对于非接口类型的自定义类型T，其方法集合由所有receiver为T类型的方法组成；
而类型*T的方法集合则包含所有receiver为T和*T类型的方法**。
也正因为如此，pt才能成功赋值给Interface类型变量。
*/

/*
另外
definded类型不继承underlying类型的方法集合！！！
但是
类型别名与原类型拥有完全相同的方法集合，无论原类型是接口类型还是非接口类型。！！！
*/
