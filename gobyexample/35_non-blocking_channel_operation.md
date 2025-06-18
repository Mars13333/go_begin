# Go 非阻塞 Channel 操作

## 核心概念

### 什么是非阻塞 Channel 操作

准确来说，应该叫 **"non-blocking channel operations"**（非阻塞的 channel 操作），而不是 "non-blocking channel"。

- **Channel 本身**：仍然是阻塞的
- **操作方式**：通过 `select` + `default` 实现非阻塞操作

### 阻塞 vs 非阻塞操作

**阻塞操作**（正常的 channel 操作）：
```go
messages <- "hi"        // 会一直等待，直到有接收者
msg := <-messages       // 会一直等待，直到有数据
```

**非阻塞操作**（通过 select + default）：
```go
select {
case messages <- "hi":
    // 能发送就发送
default:
    // 不能发送就立即执行这里，不等待
}

select {
case msg := <-messages:
    // 有数据就接收
default:
    // 没数据就立即执行这里，不等待
}
```

## select 语句中的 case

### case 后面不是条件判断

在 `select` 中，`case` 后面是 **channel 操作**，不是条件判断：

1. **channel 接收操作**：
```go
case msg := <-ch:        // 从 ch 接收数据
case <-ch:               // 从 ch 接收数据但不保存
```

2. **channel 发送操作**：
```go
case ch <- value:        // 向 ch 发送 value
```

3. **default**：
```go
default:                 // 当所有 channel 操作都无法进行时执行
```

### 执行逻辑

`case messages <- msg` 的含义是：
- **尝试**向 `messages` channel 发送 `msg`
- 如果能够发送（有接收者在等待），就执行这个 case
- 如果不能发送（没有接收者），就检查下一个 case 或执行 default

这是 **channel 操作的就绪状态检查**，而不是普通的条件判断。

## 示例分析

### 1. 非阻塞接收

```go
messages := make(chan string)  // 无缓冲channel，没有数据

select {
case msg := <-messages:
    fmt.Println("received message", msg)  // 不会执行
default:
    fmt.Println("no message received")    // 立即执行
}
```

**结果**：`no message received`

### 2. 非阻塞发送

```go
msg := "hi"
select {
case messages <- msg:
    fmt.Println("sent message", msg)      // 不会执行
default:
    fmt.Println("no message sent")        // 立即执行
}
```

**结果**：`no message sent`
**原因**：`messages` 是无缓冲 channel，且没有接收者

### 3. 多路非阻塞选择

```go
select {
case msg := <-messages:
    fmt.Println("received message", msg)  // 不会执行
case sig := <-signals:
    fmt.Println("received signal", sig)   // 不会执行
default:
    fmt.Println("no activity")            // 立即执行
}
```

**结果**：`no activity`
**原因**：两个 channel 都没有数据可接收

## 对比：普通 switch vs select

### 普通 switch（条件判断）

```go
switch x {
case 1:           // x == 1 时执行
case 2:           // x == 2 时执行
}
```

### select（channel 操作）

```go
select {
case ch1 <- data:     // 能向 ch1 发送时执行
case data := <-ch2:   // 能从 ch2 接收时执行
case <-ch3:           // 能从 ch3 接收时执行
default:              // 所有 channel 操作都不能进行时执行
}
```

## 使用场景

非阻塞 channel 操作适用于：

1. **轮询检查**：定期检查 channel 是否有数据，但不想阻塞
2. **尝试性操作**：尝试发送或接收，失败了就做其他事情
3. **避免死锁**：在可能发生阻塞的地方提供备选方案
4. **性能优化**：避免不必要的等待时间

## 关键要点

1. **Channel 本身特性不变**：channel 仍然是阻塞的
2. **操作方式改变**：通过 `select` + `default` 让操作变成非阻塞
3. **case 是操作而非条件**：`case ch <- data` 是尝试发送操作
4. **立即返回**：有 `default` 的 `select` 永远不会阻塞
5. **就绪状态检查**：Go 运行时检查哪个 channel 操作可以立即执行 