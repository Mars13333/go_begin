package main

import "sync"

var mu sync.Mutex

// 注：mu虽然未显式初始化，但是Mutex满足零值可用，所以可以直接使用

func f() {
	mu.Lock()
	defer mu.Unlock()
	bizOperation()
}

func g() {
	mu.Lock()
	bizOperation()
	mu.Unlock()
}

func bizOperation() {
	// 可能会产生错误
}

/*
使用defer让函数更简洁更健壮
当函数bizOperation抛出panic时，
函数g无法释放mutex，
而函数f则可以通过deferred函数释放mutex，
让后续函数依旧可以申请mutex资源。
*/
