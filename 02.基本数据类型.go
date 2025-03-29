package main

/*
数字类型
数字类型主要分为有符号数和无符号数，有符号数可以用来表示负数，
除此之外它们还有位数的区别，不同的位数代表它们实际存储占用空间，以及取值的范围。

bool      the set of boolean (true, false)

uint8      the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8      the set of all signed  8-bit integers (-128 to 127)
int16      the set of all signed 16-bit integers (-32768 to 32767)
int32      the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64      the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32      the set of all IEEE-754 32-bit floating-point numbers
float64      the set of all IEEE-754 64-bit floating-point numbers

表示 复数 的基本数据类型，其实部和虚部均为 float32 类型，占用 8 字节内存（实部和虚部各 4 字节）
complex64      the set of all complex numbers with float32 real and imaginary parts
表示 高精度复数 的基本数据类型，其实部和虚部均为 float64 类型，占用 16 字节内存（实部和虚部各 8 字节）
complex128      the set of all complex numbers with float64 real and imaginary parts

byte      alias for uint8

表示 Unicode 码点（Unicode Code Point） 的基本数据类型，本质上是 int32 的别名（占 4 字节），语义上更强调“字符”而非数字。
rune      alias for int32
uint      either 32 or 64 bits
int      same size as uint

一种特殊的无符号整数类型，主要用于底层编程和与指针相关的操作。
uintptr      an unsigned integer large enough to store the uninterpreted bits of a pointer value

string      the set of string value (eg: "hi")
*/

/*
布尔类型： true false
var a bool
var a = true
a := true
const a = true

*/

/*
字符串
var a = "hello" //单行字符串
var c = "\"" // 转义符
var d = `line3  //多行输出
line1
line2
`
*/

/*
特殊类型
byte，uint8 别名，用于表示二进制数据的 bytes
rune，int32 别名, 用于表示一个符号
*/

func main() {

}
