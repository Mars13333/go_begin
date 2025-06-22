# Go语言命名规范指南

## 概述

Go语言有明确的命名规范，遵循这些规范可以让代码更易读、更符合社区习惯，并且能通过代码检查工具。

## 模块命名规范

### 1. **推荐使用连字符（-）**
```bash
# 推荐
go mod init mysql-demo
go mod init web-server
go mod init api-gateway
go mod init user-management

# 不推荐
go mod init mysql_demo
go mod init web_server
go mod init api_gateway
```

### 2. **命名原则**
- **描述性**：模块名应该描述项目功能
- **简洁**：避免过长或过于复杂的名称
- **一致性**：与Go标准库和社区习惯保持一致

### 3. **实际示例**
```bash
# 数据库相关
go mod init mysql-demo
go mod init redis-cache
go mod init postgres-example

# Web相关
go mod init web-server
go mod init api-gateway
go mod init http-client

# 学习项目
go mod init go-learning
go mod init algorithm-demo
go mod init concurrency-example
```

## 包命名规范

### 1. **使用小写字母**
```go
// 推荐
package mysql
package redis
package web
package database

// 不推荐
package MySQL
package Redis
package Web
```

### 2. **避免下划线和连字符**
```go
// 推荐
package userdb
package httputil
package jsonutil

// 不推荐
package user_db
package http_util
package json-util
```

### 3. **简洁明了**
```go
// 推荐
package db      // 而不是 database
package http    // 而不是 httputil
package json    // 而不是 jsonutil

// 避免缩写
package userdatabase  // 而不是 userdb
```

## 变量命名规范

### 1. **使用小驼峰（camelCase）**
```go
// 推荐
var userName string
var databaseConnection *sql.DB
var maxRetryCount int
var isConnected bool

// 不推荐
var user_name string
var database_connection *sql.DB
var max_retry_count int
var is_connected bool
```

### 2. **局部变量**
```go
func main() {
    // 局部变量使用小驼峰
    userName := "John"
    userAge := 25
    isActive := true
    
    fmt.Printf("User: %s, Age: %d, Active: %t\n", userName, userAge, isActive)
}
```

### 3. **全局变量**
```go
// 包级变量
var (
    databaseConnection *sql.DB
    maxRetryCount      int = 3
    defaultTimeout     time.Duration = 30 * time.Second
)
```

## 函数命名规范

### 1. **使用小驼峰**
```go
// 推荐
func getUserInfo() {}
func connectDatabase() {}
func isValidEmail(email string) bool {}
func calculateTotalPrice(items []Item) float64 {}

// 不推荐
func get_user_info() {}
func connect_database() {}
func is_valid_email(email string) bool {}
```

### 2. **函数命名原则**
- **动词开头**：描述动作，如 `get`, `set`, `create`, `delete`
- **描述性**：清楚表达函数功能
- **简洁**：避免过长名称

### 3. **实际示例**
```go
// 获取数据
func getUserByID(id int) (*User, error) {}
func getProductList() ([]Product, error) {}

// 创建数据
func createUser(user *User) error {}
func createOrder(order *Order) error {}

// 验证数据
func isValidEmail(email string) bool {}
func isValidPhone(phone string) bool {}

// 处理数据
func processPayment(payment *Payment) error {}
func validateInput(input string) error {}
```

## 类型命名规范

### 1. **使用大驼峰（PascalCase）**
```go
// 推荐
type User struct {
    Name  string
    Age   int
    Email string
}

type DatabaseConnection interface {
    Connect() error
    Close() error
}

// 不推荐
type user struct {}
type database_connection interface {}
```

### 2. **结构体字段**
```go
type User struct {
    // 字段名使用大驼峰（可导出）
    Name     string
    Age      int
    Email    string
    
    // 私有字段使用小驼峰
    password string
    token    string
}
```

### 3. **接口命名**
```go
// 推荐：以 -er 结尾或描述行为
type Reader interface {
    Read(p []byte) (n int, err error)
}

type DatabaseConnection interface {
    Connect() error
    Close() error
    Query(query string) (*sql.Rows, error)
}

// 不推荐
type database interface {}
type connection interface {}
```

## 常量命名规范

### 1. **使用大驼峰或全大写**
```go
// 推荐：大驼峰
const (
    MaxRetryCount = 3
    DefaultTimeout = 30
    MaxConnections = 100
)

// 推荐：全大写（用于配置常量）
const (
    DEFAULT_TIMEOUT = 30
    MAX_RETRY_COUNT = 3
    DATABASE_URL = "localhost:3306"
)

// 不推荐
const (
    max_retry_count = 3
    default_timeout = 30
)
```

### 2. **枚举常量**
```go
// 使用 iota 定义枚举
const (
    StatusPending = iota
    StatusActive
    StatusInactive
    StatusDeleted
)

// 或者使用字符串常量
const (
    StatusPending  = "pending"
    StatusActive   = "active"
    StatusInactive = "inactive"
)
```

## 导出标识符规范

### 1. **大写开头：可导出**
```go
// 可导出的变量
var UserName string
var MaxRetryCount int

// 可导出的函数
func GetUserInfo() {}
func CreateUser() {}
func ValidateEmail() {}

// 可导出的类型
type User struct {}
type DatabaseConnection interface {}
```

### 2. **小写开头：私有**
```go
// 私有变量
var userName string
var maxRetryCount int

// 私有函数
func getUserInfo() {}
func createUser() {}
func validateEmail() {}

// 私有类型
type user struct {}
type databaseConnection interface {}
```

## 特殊命名规范

### 1. **错误变量**
```go
// 错误变量通常以 Err 开头
var (
    ErrNotFound = errors.New("not found")
    ErrInvalidInput = errors.New("invalid input")
    ErrDatabaseConnection = errors.New("database connection failed")
)
```

### 2. **测试函数**
```go
// 测试函数以 Test 开头
func TestGetUser(t *testing.T) {}
func TestCreateUser(t *testing.T) {}
func TestValidateEmail(t *testing.T) {}

// 基准测试以 Benchmark 开头
func BenchmarkGetUser(b *testing.B) {}
func BenchmarkCreateUser(b *testing.B) {}
```

### 3. **示例函数**
```go
// 示例函数以 Example 开头
func ExampleGetUser() {}
func ExampleCreateUser() {}
func ExampleValidateEmail() {}
```

## 完整示例

```go
package main

import (
    "database/sql"
    "errors"
    "fmt"
    "time"
)

// 常量：大驼峰或全大写
const (
    MaxRetryCount = 3
    DefaultTimeout = 30 * time.Second
    DEFAULT_DATABASE_URL = "localhost:3306"
)

// 错误变量：以 Err 开头
var (
    ErrUserNotFound = errors.New("user not found")
    ErrInvalidEmail = errors.New("invalid email format")
)

// 类型：大驼峰
type User struct {
    ID       int
    Name     string
    Email    string
    Age      int
    password string // 私有字段
}

type DatabaseConnection interface {
    Connect() error
    Close() error
    Query(query string) (*sql.Rows, error)
}

// 全局变量：小驼峰
var (
    databaseConnection *sql.DB
    maxRetryCount      int = 3
    isConnected        bool
)

// 私有函数：小驼峰
func connectDatabase() error {
    // 实现
    return nil
}

func getUserInfo(userID int) (*User, error) {
    // 实现
    return &User{}, nil
}

func isValidEmail(email string) bool {
    // 实现
    return true
}

// 导出函数：大写开头
func GetUserByID(userID int) (*User, error) {
    return getUserInfo(userID)
}

func CreateUser(user *User) error {
    // 实现
    return nil
}

func ValidateEmail(email string) error {
    if !isValidEmail(email) {
        return ErrInvalidEmail
    }
    return nil
}

func main() {
    // 局部变量：小驼峰
    userName := "John"
    userAge := 25
    isActive := true
    
    fmt.Printf("User: %s, Age: %d, Active: %t\n", userName, userAge, isActive)
}
```

## 代码检查工具

### 1. **golint**
```bash
# 安装
go get -u golang.org/x/lint/golint

# 使用
golint ./...
```

### 2. **gofmt**
```bash
# 格式化代码
gofmt -w .

# 检查格式
gofmt -d .
```

### 3. **go vet**
```bash
# 检查代码问题
go vet ./...
```

### 4. **静态检查工具**
```bash
# 安装 staticcheck
go install honnef.co/go/tools/cmd/staticcheck@latest

# 使用
staticcheck ./...
```

## 最佳实践总结

### 1. **命名原则**
- **清晰性**：名称应该清楚表达含义
- **一致性**：在整个项目中保持一致的命名风格
- **简洁性**：避免过长或过于复杂的名称
- **描述性**：名称应该描述功能或用途

### 2. **常见错误**
- 使用下划线命名变量和函数
- 模块名使用下划线而不是连字符
- 包名使用大写字母
- 导出标识符使用小写开头

### 3. **检查清单**
- [ ] 模块名使用连字符（-）
- [ ] 包名使用小写字母
- [ ] 变量和函数使用小驼峰
- [ ] 类型和接口使用大驼峰
- [ ] 常量使用大驼峰或全大写
- [ ] 导出标识符使用大写开头
- [ ] 错误变量以 Err 开头
- [ ] 测试函数以 Test 开头

## 总结

遵循Go语言的命名规范可以：
1. **提高代码可读性**
2. **符合社区习惯**
3. **通过代码检查工具**
4. **便于团队协作**
5. **减少维护成本**

记住：**好的命名是代码自文档化的第一步**！