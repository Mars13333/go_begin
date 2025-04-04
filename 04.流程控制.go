package main

func main() {

}

/*
if、switch、for 来进行条件判断和流程控制

if 主要用于条件判断，语法为：

if 条件 {
  # 业务代码
}

package main

import "fmt"

func main() {
    age := 7

    if age > 6 {
        fmt.Println("It's primary school")
    }
}

---

&& ||

func main() {
    age := 7

    if age > 6 && age <= 12 {
        fmt.Println("It's primary school")
    }
}

---

if else if else

func main() {
    age := 13

    if age > 6 && age <= 12 {
        fmt.Println("It's primary school")
    } else if age > 12 && age <= 15 {
        fmt.Println("It's middle school")
    } else {
        fmt.Println("It's high school")
    }
}
*/

/*
switch

如果我们的条件分支太多，可以考虑使用 switch 替换 if, 例如：
注意：在 Go 中 switch 只要匹配中了就会中止剩余的匹配项，这和 Java 很大不一样，它需要使用 break 来主动跳出

package main

import "fmt"

func main() {
    age := 10

    switch age {
    case 5:
        fmt.Println("The age is 5")
    case 7:
        fmt.Println("The age is 7")
    case 10:
        fmt.Println("The age is 10")
    default:
        fmt.Println("The age is unkown")
    }
}

---


switch 的 case 条件可以是多个值，例如：
func main() {
    age := 7

    switch age {
    case 7, 8, 9, 10, 11, 12:
        fmt.Println("It's primary school")
    case 13, 14, 15:
        fmt.Println("It's middle school")
    case 16, 17, 18:
        fmt.Println("It's high school")
    default:
        fmt.Println("The age is unkown")
    }
}


---

使用条件判断


func main() {
    age := 7

    switch {
    case age >= 6 && age <= 12:
        fmt.Println("It's primary school")
    case age >= 13 && age <= 15:
        fmt.Println("It's middle school")
    case age >= 16 && age <= 18:
        fmt.Println("It's high school")
    default:
        fmt.Println("The age is unkown")
    }
}
*/

/*
for


func main() {
    for i := 0; i < 2; i++ {
        fmt.Println("loop with index", i)
    }
}


---


for...range

使用 for..range 对数组、切片、map、 字符串等进行循环操作，例如：
func main() {
    numbers := []int{1, 2, 3}

    for i, v := range numbers {
        fmt.Printf("numbers[%d] is %d\n", i, v)
    }
}
注意: 这里的 i、v 是切片元素的位置索引和值。


func main() {
    cityCodes := map[string]int{
        "北京": 1,
        "上海": 2,
    }

    for i, v := range cityCodes {
        fmt.Printf("%s is %d\n", i, v)
    }
}
注意: 这里的 i、v 是 map 的 一组键值对的键和值。


---

使用 continue 和 break 对循环进行控制，例如：

func main() {
    numbers := []int{1, 2, 3, 4, 5}

    for i, v := range numbers {
        if v == 4 {
            break
        }

        if v%2 == 0 {
            continue
        }

        fmt.Printf("numbers[%d] is %d\n", i, v)
    }
}
注意：

break 会结束所有循环。
continue 会跳过当前循环直接进入下一次循环。

*/
