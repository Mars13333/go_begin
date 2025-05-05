package main // 创建文件默认的值是go_begin,改成main之后才能run
import "fmt"

func main() {
	hello("疯狂的世界")
}

func hello(str string) {
	fmt.Println("hello! ", str)
	//numbers := [16]int{1, 2, 3, 4}
	// n1 := [256]int{'a': 111, 'b': 8, 'c': 9}
	// fmt.Println(n1)
}
