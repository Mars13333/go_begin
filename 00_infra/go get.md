# Go Get 命令使用指南

## 概述

`go get` 是Go语言中用于下载和安装包的命令，主要用于管理项目依赖。

## 基本语法

```bash
go get [参数] 包路径
```

## 常用参数

| 参数 | 说明 |
|------|------|
| `-d` | 只下载，不安装 |
| `-u` | 更新包及其依赖 |
| `-t` | 同时下载测试依赖 |
| `-v` | 显示详细信息 |
| `-insecure` | 允许不安全的HTTP连接 |

## 基本使用场景

### 1. 安装单个包

```bash
# 安装标准库包
go get golang.org/x/text

# 安装第三方包
go get github.com/gin-gonic/gin
```

### 2. 安装特定版本

```bash
# 安装特定版本
go get github.com/gin-gonic/gin@v1.9.0

# 安装最新版本
go get github.com/gin-gonic/gin@latest
```

### 3. 更新包

```bash
# 更新包到最新版本
go get -u github.com/gin-gonic/gin

# 更新包及其所有依赖
go get -u -t github.com/gin-gonic/gin
```

### 4. 只下载不安装

```bash
# 只下载源码，不编译安装
go get -d github.com/gin-gonic/gin
```

## 实际示例

### 在Go模块项目中使用

```bash
# 进入项目目录
cd your-project

# 初始化Go模块（如果还没有go.mod）
go mod init your-project-name

# 安装依赖
go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql
```

### 常用包安装示例

```bash
# Web框架
go get github.com/gin-gonic/gin
go get github.com/gorilla/mux

# 数据库驱动
go get github.com/go-sql-driver/mysql
go get github.com/lib/pq

# 配置管理
go get github.com/spf13/viper

# 日志库
go get go.uber.org/zap
go get github.com/sirupsen/logrus

# 测试框架
go get github.com/stretchr/testify
```

### 版本管理

```bash
# 查看可用版本
go list -m -versions github.com/gin-gonic/gin

# 安装特定版本
go get github.com/gin-gonic/gin@v1.9.0

# 升级到最新版本
go get -u github.com/gin-gonic/gin
```

### 批量更新

```bash
# 更新所有依赖
go get -u ./...

# 更新所有依赖并显示详细信息
go get -u -v ./...
```

## 网络配置

### 设置代理（解决网络问题）

```bash
# 设置国内代理
go env -w GOPROXY=https://goproxy.cn,direct

# 设置官方代理
go env -w GOPROXY=https://proxy.golang.org,direct

# 查看当前代理设置
go env GOPROXY
```

### 私有仓库配置

```bash
# 设置私有仓库
go env -w GOPRIVATE=*.gitlab.com,*.gitee.com

# 配置Git认证（如果需要）
git config --global url."https://username:token@github.com/".insteadOf "https://github.com/"
```

## 相关命令

### 依赖管理

```bash
# 整理依赖
go mod tidy

# 下载依赖
go mod download

# 创建vendor目录
go mod vendor

# 查看依赖图
go mod graph
```

### 包信息查询

```bash
# 查看包信息
go list -m github.com/gin-gonic/gin

# 查看所有依赖
go list -m all

# 查看依赖更新
go list -m -u all
```

## 注意事项

### 1. Go版本要求
- 需要Go 1.11+支持模块模式
- 建议使用Go 1.16+以获得更好的模块支持

### 2. 项目结构
- 确保项目有 `go.mod` 文件
- 在项目根目录执行命令

### 3. 常见问题解决

#### 依赖冲突
```bash
# 清理并重新整理依赖
go mod tidy
```

#### 版本不兼容
```bash
# 查看依赖关系
go mod graph

# 手动指定版本
go get package@version
```

#### 网络超时
```bash
# 设置超时时间
go env -w GOSUMDB=sum.golang.org
go env -w GOSUMDB_TIMEOUT=30s
```

## 最佳实践

1. **使用模块模式**: 确保项目使用Go模块
2. **定期更新**: 定期使用 `go get -u ./...` 更新依赖
3. **版本锁定**: 在生产环境中锁定依赖版本
4. **依赖审查**: 定期审查和清理未使用的依赖
5. **使用代理**: 配置合适的代理以提高下载速度

## 命令对比

| 命令 | 用途 | 适用场景 |
|------|------|----------|
| `go get` | 下载和安装包 | 添加新依赖 |
| `go mod tidy` | 整理依赖 | 清理未使用依赖 |
| `go mod download` | 下载依赖 | 离线环境准备 |
| `go mod vendor` | 创建vendor | 确保依赖版本一致 |

## 总结

`go get` 是Go语言依赖管理的核心命令，掌握其使用方法对于Go项目开发至关重要。结合 `go mod` 系列命令，可以有效地管理项目依赖。
