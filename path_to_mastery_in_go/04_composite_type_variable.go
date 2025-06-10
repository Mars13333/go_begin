
package main

import "fmt"

/*
Go语言中的复合类型包括 结构体、数组、切片和map。
对于复合类型变量，最常见的值构造方式就是对其内部元素进行逐个赋值
*/

var s myStruct
s.name="tony"
s.age=23

var a[5] int
a[0]=13
a[1]=14
a[2]=15
a[3]=16
a[4]=17

sl:=make([]int,5,5)
sl[0]=23
sl[1]=24
sl[2]=25
sl[3]=26
sl[4]=27


m:=make(map[int]string)
m[1]="hello"
m[2]="gopher"
m[3]="!"


/*
复合字面值语法

但这样的值构造方式让代码显得有些烦琐，尤其是在构造组成较为复杂的复合类型变量的初值时。
Go提供的复合字面值（composite literal）语法可以作为复合类型变量的初值构造器。

复合字面值由俩部分组成，一部分是类型，一部分是由大括号包裹的字面值。
*/


s:=myStruct{"tony",23}
a:=[5]int{13,14,15,16,17}
sl:=[]int{23,24,25,26,27}
m:=map[int]string{1:"hello",2:"gopher",3:"!"}