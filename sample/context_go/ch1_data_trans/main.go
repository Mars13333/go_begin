package main

import (
	"context"
	"fmt"
)

// 定义一个私有类型，避免外部包用同样的 key
type userKey struct{}

// 作为上下文，可以传递数据

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, userKey{}, User{Name: "Alice"})
	GetUser(ctx)
}

type User struct {
	Name string
}

func GetUser(ctx context.Context) {
	fmt.Println("User Name:", ctx.Value(userKey{}).(User).Name) // 断言
}
