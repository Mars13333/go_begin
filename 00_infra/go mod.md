# Go Mod 命令使用指南

## 概述

`go mod` 是Go语言的模块管理命令，用于管理Go模块的依赖关系。它是Go 1.11+引入的官方依赖管理工具，替代了之前的GOPATH模式。

## 基本语法

```bash
go mod [子命令] [参数]
```

## 子命令详解

### 1. `go mod init` - 初始化模块

**功能**: 在当前目录初始化一个新的Go模块

```bash
# 基本用法
go mod init module-name

# 示例
go mod init github.com/username/project-name
go mod init my-project
```

**生成文件**: `go.mod`

### 2. `go mod tidy` - 整理依赖

**功能**: 添加缺失的依赖，移除未使用的依赖

```bash
# 整理当前模块的依赖
go mod tidy

# 整理指定模块的依赖
go mod tidy ./path/to/module
```

**作用**:
- 添加代码中引用但未在go.mod中声明的依赖
- 移除go.mod中存在但代码中未使用的依赖
- 更新go.sum文件

### 3. `go mod download` - 下载依赖

**功能**: 下载模块依赖到本地缓存

```bash
# 下载所有依赖
go mod download

# 下载特定模块
go mod download github.com/gin-gonic/gin

# 下载并显示详细信息
go mod download -x
```

### 4. `go mod verify` - 验证依赖

**功能**: 验证依赖的完整性和正确性

```bash
# 验证所有依赖
go mod verify

# 验证特定模块
go mod verify github.com/gin-gonic/gin
```

### 5. `go mod vendor` - 创建vendor目录

**功能**: 将依赖复制到vendor目录

```bash
# 创建vendor目录
go mod vendor

# 使用vendor模式编译
go build -mod=vendor
```

### 6. `go mod graph` - 显示依赖图

**功能**: 显示模块依赖关系图

```bash
# 显示所有依赖关系
go mod graph

# 过滤特定模块的依赖
go mod graph | grep github.com/gin-gonic/gin
```

### 7. `go mod why` - 解释依赖关系

**功能**: 解释为什么需要某个依赖

```bash
# 解释特定模块的依赖原因
go mod why github.com/gin-gonic/gin

# 解释特定包的依赖原因
go mod why github.com/gin-gonic/gin/gin.go
```

### 8. `go mod edit` - 编辑go.mod文件

**功能**: 以编程方式编辑go.mod文件

```bash
# 添加依赖
go mod edit -require=github.com/gin-gonic/gin@v1.9.0

# 移除依赖
go mod edit -droprequire=github.com/gin-gonic/gin

# 添加替换规则
go mod edit -replace=old/path=new/path@version

# 添加排除规则
go mod edit -exclude=github.com/gin-gonic/gin@v1.8.0
```

## go.mod 文件结构

```go
module github.com/username/project-name

go 1.21

require (
    github.com/gin-gonic/gin v1.9.0
    github.com/go-sql-driver/mysql v1.7.0
)

require (
    github.com/bytedance/sonic v1.9.0 // indirect
    github.com/gabriel-vasile/mimetype v1.4.2 // indirect
    // ... 其他间接依赖
)

replace github.com/gin-gonic/gin => ./local/gin

exclude github.com/gin-gonic/gin v1.8.0
```

## go.sum 文件

**功能**: 存储模块的加密哈希值，用于验证模块完整性

```go
github.com/bytedance/sonic v1.9.0 h1:6iJ6NqdoxCDr6mbY8h18oSO+cShGSMRGCEo7F2h0x8s=
github.com/bytedance/sonic v1.9.0/go.mod h1:i736AoUSYt75HyZLoJW9ERYxcy6eaN6h4BZXU064P/U=
// ... 更多哈希值
```

## 实际使用示例

### 1. 创建新项目

```bash
# 创建项目目录
mkdir my-go-project
cd my-go-project

# 初始化模块
go mod init github.com/username/my-go-project

# 添加依赖
go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql

# 整理依赖
go mod tidy
```

### 2. 管理现有项目

```bash
# 进入项目目录
cd existing-project

# 如果项目没有go.mod，初始化
go mod init

# 整理依赖
go mod tidy

# 验证依赖
go mod verify
```

### 3. 处理依赖冲突

```bash
# 查看依赖图
go mod graph

# 查看特定依赖的原因
go mod why github.com/conflicting/package

# 手动编辑依赖版本
go mod edit -require=github.com/conflicting/package@v1.2.3

# 使用替换规则
go mod edit -replace=github.com/conflicting/package=github.com/forked/package@v1.2.3
```

### 4. 离线开发

```bash
# 下载所有依赖到本地
go mod download

# 创建vendor目录
go mod vendor

# 使用vendor模式编译
go build -mod=vendor
```

## 常用工作流程

### 1. 新项目初始化流程

```bash
# 1. 创建项目目录
mkdir new-project && cd new-project

# 2. 初始化模块
go mod init github.com/username/new-project

# 3. 编写代码并添加依赖
go get github.com/gin-gonic/gin

# 4. 整理依赖
go mod tidy

# 5. 验证项目
go mod verify
```

### 2. 依赖更新流程

```bash
# 1. 查看当前依赖
go list -m all

# 2. 查看可更新的依赖
go list -m -u all

# 3. 更新特定依赖
go get -u github.com/gin-gonic/gin

# 4. 更新所有依赖
go get -u ./...

# 5. 整理依赖
go mod tidy

# 6. 验证更新
go mod verify
```

### 3. 问题排查流程

```bash
# 1. 查看依赖图
go mod graph

# 2. 解释特定依赖
go mod why github.com/problematic/package

# 3. 验证依赖完整性
go mod verify

# 4. 重新下载依赖
go mod download

# 5. 清理并重新整理
go mod tidy
```

## 环境变量配置

### 1. 代理设置

```bash
# 设置代理
go env -w GOPROXY=https://goproxy.cn,direct

# 设置私有仓库
go env -w GOPRIVATE=*.gitlab.com,*.gitee.com

# 设置校验数据库
go env -w GOSUMDB=sum.golang.org
```

### 2. 模块模式设置

```bash
# 启用模块模式
go env -w GO111MODULE=on

# 禁用模块模式
go env -w GO111MODULE=off

# 自动模式
go env -w GO111MODULE=auto
```

## 最佳实践

### 1. 版本管理

- **使用语义化版本**: 遵循 `v1.2.3` 格式
- **锁定版本**: 在生产环境中使用固定版本
- **定期更新**: 定期更新依赖以获取安全补丁

### 2. 依赖管理

- **最小化依赖**: 只添加必要的依赖
- **定期清理**: 使用 `go mod tidy` 清理未使用的依赖
- **审查依赖**: 定期审查依赖的安全性和维护状态

### 3. 项目结构

- **模块命名**: 使用有意义的模块名
- **版本兼容**: 确保依赖版本兼容性
- **文档化**: 在README中说明依赖要求

### 4. 开发流程

- **版本控制**: 将go.mod和go.sum加入版本控制
- **CI/CD**: 在CI/CD中验证依赖
- **测试**: 在更新依赖后运行测试

## 常见问题解决

### 1. 依赖下载失败

```bash
# 设置代理
go env -w GOPROXY=https://goproxy.cn,direct

# 清理缓存
go clean -modcache

# 重新下载
go mod download
```

### 2. 版本冲突

```bash
# 查看冲突
go mod graph

# 使用替换规则
go mod edit -replace=old/path=new/path@version

# 整理依赖
go mod tidy
```

### 3. 模块路径问题

```bash
# 检查模块路径
go mod edit -module=new/path

# 更新导入路径
find . -name "*.go" -exec sed -i 's|old/path|new/path|g' {} \;
```

## 命令对比表

| 命令 | 功能 | 使用场景 |
|------|------|----------|
| `go mod init` | 初始化模块 | 新项目创建 |
| `go mod tidy` | 整理依赖 | 依赖管理 |
| `go mod download` | 下载依赖 | 离线准备 |
| `go mod verify` | 验证依赖 | 安全检查 |
| `go mod vendor` | 创建vendor | 离线开发 |
| `go mod graph` | 显示依赖图 | 问题排查 |
| `go mod why` | 解释依赖 | 依赖分析 |
| `go mod edit` | 编辑go.mod | 手动管理 |

## 总结

`go mod` 是Go语言现代化的依赖管理工具，提供了完整的模块生命周期管理功能。掌握这些命令对于Go项目开发至关重要，能够有效管理项目依赖，提高开发效率。
