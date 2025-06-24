## 智能选择字段

在 GORM 中，您可以使用 Select 方法有效地选择特定字段。 这在Model字段较多但只需要其中部分的时候尤其有用，比如编写API响应。

```go
type User struct {
  ID     uint
  Name   string
  Age    int
  Gender string
  // 很多很多字段
}

type APIUser struct {
  ID   uint
  Name string
}

// 在查询时，GORM 会自动选择 `id `, `name` 字段
db.Model(&User{}).Limit(10).Find(&APIUser{})
// SQL: SELECT `id`, `name` FROM `users` LIMIT 10
```

注意 在 QueryFields 模式中, 所有的模型字段（model fields）都会被根据他们的名字选择。

```go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  QueryFields: true,
})

// 当 QueryFields 被设置为 true 时，此行为默认进行
db.Find(&user)
// SQL: SELECT `users`.`name`, `users`.`age`, ... FROM `users`

// 开启 QueryFields 并使用会话模式（Session mode）
db.Session(&gorm.Session{QueryFields: true}).Find(&user)
// SQL: SELECT `users`.`name`, `users`.`age`, ... FROM `users`
```

### QueryFields 模式推荐策略分析

`QueryFields` 模式**不推荐默认开启**，但在特定场景下很有用。

#### QueryFields 模式的作用

**默认行为 vs QueryFields 模式**

```go
// 默认行为（不开启 QueryFields）
db.Find(&user)
// SQL: SELECT * FROM users

// 开启 QueryFields 模式
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
    QueryFields: true,
})
db.Find(&user)
// SQL: SELECT `users`.`name`, `users`.`age`, `users`.`email`, ... FROM users
```

#### 优缺点分析

**优点 ✅**

**1. 性能优化**
```go
type User struct {
    ID       uint
    Name     string
    Email    string
    Avatar   []byte  // 大字段
    Profile  string  // 大字段
    // ... 50个字段
}

// 开启 QueryFields，只查询需要的字段
db.Session(&gorm.Session{QueryFields: true}).Find(&users)
// 避免传输不必要的大字段数据
```

**2. 网络传输优化**
```go
// 在微服务架构中，减少网络传输
// 只查询实际需要的字段，而不是 SELECT *
```

**3. 明确的字段依赖**
```go
// 让SQL更明确，便于数据库优化器分析
// SELECT `users`.`name`, `users`.`email` FROM users
// 而不是 SELECT * FROM users
```

**缺点 ❌**

**1. SQL 语句变长**
```go
// 字段很多时，SQL会变得很长
// SELECT `users`.`id`, `users`.`name`, `users`.`email`, `users`.`phone`, ... FROM users
```

**2. 缓存效率降低**
```go
// 每次查询的SQL都明确列出字段，可能影响SQL缓存
```

**3. 调试复杂化**
```go
// 日志中的SQL语句变得冗长，不如 SELECT * 简洁
```

#### 推荐的使用策略

**❌ 不推荐全局开启**
```go
// 不推荐：全局开启
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
    QueryFields: true,  // 不推荐全局开启
})
```

**✅ 推荐按需开启**
```go
// 推荐：在特定场景下使用
func getUsersForAPI(db *gorm.DB) []User {
    var users []User
    
    // 只在需要性能优化的场景下开启
    db.Session(&gorm.Session{QueryFields: true}).
       Select("id", "name", "email").  // 结合 Select 使用
       Find(&users)
    
    return users
}
```

#### 实际应用场景

**场景1：大表查询优化**
```go
type User struct {
    ID          uint
    Name        string
    Email       string
    Avatar      []byte    // 1MB 大字段
    Resume      string    // 10KB 大字段
    Settings    string    // 5KB JSON字段
    // ... 更多字段
}

// 列表页面只需要基本信息
func getUserList(db *gorm.DB) []User {
    var users []User
    
    // 使用 QueryFields + Select 优化
    db.Session(&gorm.Session{QueryFields: true}).
       Select("id", "name", "email").
       Find(&users)
    
    return users
}
```

**场景2：API 响应优化**
```go
// API 返回时，明确只查询需要的字段
func getUserAPI(c *gin.Context) {
    var users []User
    
    // 根据API需求，只查询必要字段
    db.Session(&gorm.Session{QueryFields: true}).
       Select("id", "name", "email", "created_at").
       Find(&users)
    
    c.JSON(200, users)
}
```

**场景3：报表查询**
```go
// 报表查询，只需要统计字段
func getUserStats(db *gorm.DB) {
    type UserStat struct {
        Date  time.Time
        Count int
    }
    
    var stats []UserStat
    
    // 报表查询，明确字段更好
    db.Session(&gorm.Session{QueryFields: true}).
       Table("users").
       Select("DATE(created_at) as date, COUNT(*) as count").
       Group("DATE(created_at)").
       Scan(&stats)
}
```

#### 最佳实践建议

**1. 按场景选择**
```go
// ✅ 推荐：根据具体需求选择
func getUsers(db *gorm.DB, scenario string) []User {
    var users []User
    
    switch scenario {
    case "list":
        // 列表页面，只要基本字段
        db.Select("id", "name", "email").Find(&users)
    case "detail":
        // 详情页面，需要所有字段
        db.Find(&users)
    case "export":
        // 导出场景，使用 QueryFields 优化
        db.Session(&gorm.Session{QueryFields: true}).Find(&users)
    }
    
    return users
}
```

**2. 结合其他优化手段**
```go
// ✅ 最佳实践：结合多种优化方法
func optimizedQuery(db *gorm.DB) {
    var users []User
    
    db.Session(&gorm.Session{QueryFields: true}).  // 明确字段
       Select("id", "name", "email").              // 只选择需要的
       Where("status = ?", "active").              // 减少数据量
       Limit(100).                                 // 分页
       Find(&users)
}
```

**总结：**
- ❌ **不推荐全局开启** - 会让所有SQL变复杂
- ✅ **推荐按需使用** - 在性能敏感场景下开启
- ✅ **结合 Select 使用** - 获得最佳性能
- ✅ **用于大表优化** - 避免查询不必要的大字段

**什么时候使用：** 表字段很多（>20个字段）、包含大字段（BLOB、TEXT）、API响应需要优化、网络带宽有限的环境

**什么时候不用：** 小表查询、开发调试阶段、需要所有字段的场景、简单的CRUD操作

## 锁

GORM 支持多种类型的锁，例如：

```go
// 基本的 FOR UPDATE 锁
db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&users)
// SQL: SELECT * FROM `users` FOR UPDATE
```

上述语句将会在事务（transaction）中锁定选中行（selected rows）。 可以被用于以下场景：当你准备在事务（transaction）中更新（update）一些行（rows）时，并且想要在本事务完成前，阻止（prevent）其他的事务（other transactions）修改你准备更新的选中行。

Strength 也可以被设置为 SHARE ，这种锁只允许其他事务读取（read）被锁定的内容，而无法修改（update）或者删除（delete）。

```go
db.Clauses(clause.Locking{
  Strength: "SHARE",
  Table: clause.Table{Name: clause.CurrentTable},
}).Find(&users)
// SQL: SELECT * FROM `users` FOR SHARE OF `users`
```

Table选项用于指定将要被锁定的表。 这在你想要 join 多个表，并且锁定其一时非常有用。

你也可以提供如 NOWAIT 的Options，这将尝试获取一个锁，如果锁不可用，导致了获取失败，函数将会立即返回一个error。 当一个事务等待其他事务释放它们的锁时，此Options（Nowait）可以阻止这种行为

```go
db.Clauses(clause.Locking{
  Strength: "UPDATE",
  Options: "NOWAIT",
}).Find(&users)
// SQL: SELECT * FROM `users` FOR UPDATE NOWAIT
```

Options也可以是SKIP LOCKED，设置后将跳过所有已经被其他事务锁定的行（any rows that are already locked by other transactions.）。 这次高并发情况下非常有用：那时你可能会想要对未经其他事务锁定的行进行操作（process ）。

## 子查询

### From 子查询

## Group 条件

## 带多个列的 In

## 命名参数

## Find 至 map

## FirstOrInit

### 使用 Attrs 进行初始化

### 为属性使用 Assign

## FirstOrCreate

### 配合 Attrs 使用 FirstOrCreate

### 配合 Assign 使用 FirstOrCreate

## 优化器、索引提示

### 索引提示

## 迭代

## FindInBatches

## 查询钩子

## Pluck

## Scope

### 定义 Scopes

### 在查询中使用 Scopes

## Count

### 得到匹配记录的 Count

### 配合 Distinct 和 Group 使用 Count