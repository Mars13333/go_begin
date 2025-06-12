package main

import (
	"errors"
	"fmt"
)

/*
nil error值!=nil
对于接口类型变量的内部表示，最能引起大家好奇心的例子如下
*/

type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad error"),
}

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

func main() {
	e := returnsError()
	if e != nil {
		fmt.Println("error: ", e)
		return
	}
	fmt.Println("ok")
}

/*
初学者的思路大致是这样的：p为nil，returnsError返回p，那么main函数中的e就等于nil，
于是程序输出ok后退出。但真实的运行结果是什么样的呢？我们来看一下：
error:  <nil>

我们看到：示例程序并未如初学者预期的那样输出ok，程序显然是进入了错误处理分支，输出了e的值。
于是疑惑出现了：明明returnsError函数返回的p值为nil，为何却满足了if e != nil的条件进入错误处理分支呢？
要想弄清楚这个问题，非了解**接口类型变量的内部表示**不可。
*/
