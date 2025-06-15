/*
range 可以遍历多种内置数据结构
可以使用 range 来求切片中数字的和，并指出数组也可以以相同的方式。

在数组和切片上使用 range 会为每个元素提供索引和值。
有时候不需要索引，所以使用空白标识符 _ 忽略它。但有时我们确实需要索引。

在映射上使用 range 会遍历键值对。
range 也可以只遍历映射的键。

在字符串上使用 range 会遍历 Unicode 码点。第一个值是字符（rune）的起始字节索引，第二个值是字符本身。
*/

package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { // 忽略index
		sum += num
	}
	fmt.Println("sum:", sum)

	for idx, num := range nums { // 前一个是index，无论怎么命名
		if num == 3 {
			fmt.Println("index:", idx)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for kaaa, vbbb := range kvs { // 前一个是k,后一个v，无论怎么命名
		fmt.Printf("%s -> %s\n", kaaa, vbbb)
	}

	for k := range kvs { // 前一个是k,后一个v，如果只写一个，默认是key
		fmt.Println("just key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c) // 可以使用string()显式转换成文本
	}

}
