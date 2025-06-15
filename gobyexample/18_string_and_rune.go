/*
介绍Go 语言中字符串和 Unicode 码点（rune）的概念，以及如何在 Go 中处理和操作字符串。

Go 语言和标准库将字符串视为 UTF-8 编码文本的容器。

Go 语言中的字符串是 UTF-8 编码的只读字节切片。

在其他语言中，字符串由“字符”组成；而在 Go 中，字符的概念被称为“rune”，它是一个表示 Unicode 码点的整数。

由于字符串等价于 []byte，这将产生存储在内的原始字节的长度。

要计算字符串中有多少个 rune，可以使用 utf8 包。

一些泰语字符由可以跨越多个字节的 UTF-8 码点表示，因此这个计数的结果可能会令人惊讶。

*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("Len:", len(s)) // 18

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) //e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s)) //6

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' { // 不能使用双引号：invalid operation: r == "t" (mismatched types rune and untyped string)
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found sua")
	}
}
