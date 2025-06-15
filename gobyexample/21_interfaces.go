/*
接口

接口在 Go 中是方法签名的命名集合，用于定义一组行为。

在 Go 中实现接口只需要实现接口中的所有方法。

如果变量具有接口类型，则可以调用接口中的方法。

有时了解接口值的运行时类型很有用。一种方法是使用类型断言，另一种是类型切换。

只要实现了接口，就可以将结构体实例用作接口方法的参数。
*/

package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return math.Pi * c.radius * 2
}

// measure 测量
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// detect 发现
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)

}
