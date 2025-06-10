package main

import "fmt"
import "testing"

/*
字符串类型是现代编程语言中最常使用的数据类型之一。
在Go语言的先祖之一C语言当中，字符串类型并没有被显式定义，
而是以字符串字面值常量或以'\0'结尾的字符类型（char）数组来呈现的。

这给C程序员在使用字符串时带来一些问题，诸如：

类型安全性差；
字符串操作要时时刻刻考虑结尾的'\0'；
字符串数据可变（主要指以字符数组形式定义的字符串类型）​；
获取字符串长度代价大（O(n)的时间复杂度）​；
未内置对非ASCII字符（如中文字符）的处理。

Go语言修复了C语言的这些“缺陷”​，内置了string类型，统一了对字符串的抽象。
在Go语言中，无论是字符串常量、字符串变量还是代码中出现的字符串字面量，它们的类型都被统一设置为string。
*/

func main1() {
	/*
	   1.string类型的数据是不可变的

	   一旦声明了一个string类型的标识符，无论是常量还是变量，该标识符所指代的数据在整个程序的生命周期内便无法更改。

	   试图将string转换为一个切片并通过该切片对其内容进行修改，但结果事与愿违。
	   对string进行切片化后，Go编译器会为切片变量重新分配底层存储而**不是共用string的底层存储**，
	   因此对切片的修改并未对原string的数据产生任何影响。
	*/

	//原始字符串
	var s string = "hello"
	fmt.Println("original string: ", s)

	// 切片化后试图改变原字符串
	sl := []byte(s)
	sl[0] = 't'
	fmt.Println("slice: ", string(sl))
	fmt.Println("after reslice, the original string still is: ", string(s))
}

func main2() {
	/*
		2.零值可用
		Go string类型支持“零值可用”的理念。Go字符串无须像C语言中那样考虑结尾'\0'字符，因此其零值为""，长度为0。
	*/

	/*
		3.获取长度的时间复杂度是O(1)级别
		Go string类型数据是不可变的，因此一旦有了初值，那块数据就不会改变，其长度也不会改变。
		Go将这个长度作为一个字段存储在运行时的string类型的内部表示结构中。
		这样获取string长度的操作，即len(s)实际上就是读取存储在运行时中的那个长度值，这是一个代价极低的O(1)操作。
	*/

	var s string
	fmt.Println(s)      // s=""
	fmt.Println(len(s)) // 0
}

func main3() {
	/*
		4.支持通过+/+=操作符进行字符串连接

		Go还提供了其他一些构造字符串的方法,
		比如：fmt.Sprintf，strings.Join，strings.Builder，bytes.Buffer
		但直接使用+或者+=即可，最自然、开发体验最好的一种。
	*/
	s := "hello"
	s = s + ", "
	s += "gopher!"
	fmt.Println(s) //hello, gopher!
}

func main4() {
	/*
		5.支持各种比较关系操作符

		由于Go string是不可变的，因此如果两个字符串的长度不相同，那么无须比较具体字符串数据即可断定两个字符串是不同的。
		如果长度相同，则要进一步判断数据指针是否指向同一块底层存储数据。
		如果相同，则两个字符串是等价的；如果不同，则还需进一步比对实际的数据内容。
	*/

	// ==
	s1 := "世界和平"
	s2 := "世界" + "和平"
	// s2[0]='A' // cannot assign to s2[0] (value of type byte)
	/*
		！！！！
		字符串是不可变的, 所以创建出来s2后，内容就不能被修改。
		而下列的s1,s2重新赋值成"Go","23456"等，是允许的，属于被重新赋值。并不违反字符串的不可变性，而是正常地将变量指向了新的字符串对象。
	*/
	fmt.Println(s1 == s2) // true

	//!=
	s1 = "GO"
	s2 = "C"
	fmt.Println(s1 != s2) //true

	// <和<=
	s1 = "12345"
	s2 = "23456"
	fmt.Println(s1 < s2)  //true
	fmt.Println(s1 <= s2) //true

	//>和>=
	s1 = "12345"
	s2 = "123"
	fmt.Println(s1 > s2)  //true
	fmt.Println(s1 >= s2) //true
}

func main5() {

	/*
			6.对非ASCII字符提供原生支持

			something different?
		Go语言源文件默认采用的Unicode字符集。Unicode字符集是目前市面上最流行的字符集，几乎囊括了所有主流非ASCII字符（包括中文字符）​。
		Go字符串的每个字符都是一个Unicode字符，并且这些Unicode字符是以UTF-8编码格式存储在内存当中的。
	*/
	// 中文字符		Unicode码点		UTF8编码
	// 中					U+4E2D				E4B8AD
	// 国					U+56FD				E59BBD
	// 欢					U+6B22				E6ACA2
	// 迎					U+8FCE				E8BF8E
	// 你					U+60A8				E682A8
	s := "中国欢迎你"
	rs := []rune(s)
	sl := []byte(s)
	for i, v := range rs {
		var utf8Bytes []byte
		for j := i * 3; j < (i+1)*3; j++ {
			utf8Bytes = append(utf8Bytes, sl[j])
		}
		fmt.Println("%s => %X => %X\n", string(v), v, utf8Bytes)
	}
}

func main6() {
	/*
		7.原生支持多行字符串
	*/
	s := `好于知识界，当春乃发生。
随风潜入夜，润物细无声。
	`
	fmt.Println(s)
}

func main() {
	/*
		8.字符串的高效转换
		string和[​]rune、​[​]byte可以双向转换。下面就是从[​]rune或[​]byte反向转换为string的例子。

		无论是string转slice还是slice转string，转换都是要**付出代价**的，
		这些代价的根源在于string是不可变的，运行时要为转换后的类型**分配新内存**。

		想要更高效地进行转换，唯一的方法就是减少甚至避免额外的内存分配操作。

		slice类型是不可比较的，而string类型是可比较的。
		因此在日常Go编码中，我们会经常遇到将slice临时转换为string的情况。Go编译器为这样的场景提供了优化。
		在运行时中有一个名为slicebytetostringtmp的函数就是协助实现这一优化的。

		该函数的“秘诀”就在于不为string新开辟一块内存，而是直接使用slice的底层存储。
		当然使用这个函数的前提是：在原slice被修改后，这个string不能再被使用了。
	*/
	rs := []rune{
		0x4E2D,
		0x56FD,
		0x6B22,
		0x8FCE,
		0x60A8,
	}

	s := string(rs)
	fmt.Println(s)

	sl := []byte{
		0xE4, 0xB8, 0xAD,
		0xE5, 0x9B, 0xBD,
		0xE6, 0xAC, 0xA2,
		0xE8, 0xBF, 0x8E,
		0xE6, 0x82, 0xA8,
	}

	s = string(sl) //s重新赋值
	fmt.Println(s)

	fmt.Println()
	fmt.Println("-------")
	fmt.Println()

	/*
		Go编译器对用在**for-range循环**中的string到[​]byte的转换也有优化处理，
		它不会为[​]byte进行额外的内存分配，而是直接使用string的底层数据。

		从结果看到，convertWithOptimize函数将string到\[​]byte的转换放在for-range循环中，Go编译器对其进行了优化，节省了一次内存分配操作。

		此外，Go语言还在标准库中提供了strings和strconv包，可以辅助Gopher对string类型数据进行更多高级操作。
	*/
	fmt.Println(testing.AllocsPerRun(1, convert))             // 1
	fmt.Println(testing.AllocsPerRun(1, convertWithOptimize)) // 0

}

func convert() {
	s := "中国欢迎你，北京欢迎你"
	sl := []byte(s)
	for _, v := range sl {
		_ = v
	}
}

func convertWithOptimize() {
	s := "中国欢迎你，北京欢迎你"
	for _, v := range []byte(s) {
		_ = v
	}
}
