package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

/*
Springf是啥？
Sprintf 是 Go 语言中的一个格式化输出函数，属于 fmt 包。
它用于将格式化的字符串输出到标准输出（通常是控制台）。
Sprintf 的格式化规则与 Printf 类似，但不会直接输出结果，而是返回一个字符串。
*/
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with 42"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)

	var ae *argError // 声明一个*argError类型的变量，初始值为nil
	// errors.As会判断err是否可以转换为*argError类型
	// &ae是ae变量的地址，类型为**argError，传递给errors.As后，匹配成功会自动把err赋值给ae
	if errors.As(err, &ae) { // 如果err是*argError类型，或者包含*argError类型的包装错误，如果是，就把 err 赋值给 ae。
		fmt.Println(ae.arg)     // 访问ae的arg字段
		fmt.Println(ae.message) // 访问ae的message字段
	} else {
		// 如果err不是*argError类型
		fmt.Println("err doesn't match argError")
	}
}
