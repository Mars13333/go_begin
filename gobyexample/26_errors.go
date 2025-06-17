package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cant't work with 42")
	}
	return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("cant't boil water")

/*
fmt.Errorf 的 %w ?
%w 是 Go 1.13 引入的错误包装（error wrapping）语法。
它用于将错误包装在另一个错误中，同时保留原始错误的信息。
创建一个新的错误，消息是 "make tea: " + ErrPower 的错误消息
同时将 ErrPower 作为被包装的错误保存起来
这样既提供了上下文信息，又保留了原始错误
*/

/*
fmt.Errorf 类型就是err ?
是的，fmt.Errorf 返回的类型实现了 error 接口。在 Go 1.13+ 中，它还实现了 Unwrap() 方法，
用于错误链的展开。
*/
func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("make tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	fmt.Println("--------------------------------")

	/*
		为什么i为4的时候，会打印出Now it is dark.
		makeTea中 arg==4时，返回的是fmt.Errorf("make tea: %w", ErrPower)
		为什么errors.Is(err, ErrPower)为true？


		这是错误包装机制的核心特性：
		// 当 arg == 4 时
		return fmt.Errorf("make tea: %w", ErrPower)
		这行代码创建了一个包装错误，其中：
		错误消息是 "make tea: cant't boil water"
		被包装的原始错误是 ErrPower。


		errors.Is(err, ErrPower) 的工作原理：
		首先检查 err 是否等于 ErrPower（不相等）
		然后检查 err 是否实现了 Unwrap() 方法（是的）
		调用 err.Unwrap() 得到被包装的错误 ErrPower
		检查被包装的错误是否等于 ErrPower（相等）
		因此返回 true。
		详见下方，演示错误包装机制的示例。
	*/
	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Println("unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}

	// 演示错误包装机制的示例
	demonstrateErrorWrapping()
}

// 演示错误包装机制的示例
func demonstrateErrorWrapping() {
	fmt.Println("\n=== 错误包装机制演示 ===")

	// 创建原始错误
	originalErr := errors.New("原始错误")

	// 包装错误
	wrappedErr := fmt.Errorf("包装后的错误: %w", originalErr)

	fmt.Printf("原始错误: %v\n", originalErr)
	fmt.Printf("包装错误: %v\n", wrappedErr)

	// 直接比较
	fmt.Printf("wrappedErr == originalErr: %t\n", wrappedErr == originalErr) // false

	// 使用 errors.Is 比较
	fmt.Printf("errors.Is(wrappedErr, originalErr): %t\n", errors.Is(wrappedErr, originalErr)) // true

	// 多层包装
	doubleWrappedErr := fmt.Errorf("再次包装: %w", wrappedErr)
	fmt.Printf("双层包装错误: %v\n", doubleWrappedErr)
	fmt.Printf("errors.Is(doubleWrappedErr, originalErr): %t\n", errors.Is(doubleWrappedErr, originalErr)) // true
}
