## 原生 SQL

Query Raw SQL with Scan

```go
type Result struct {
  ID   int
  Name string
  Age  int
}

var result Result
db.Raw("SELECT id, name, age FROM users WHERE id = ?", 3).Scan(&result)

db.Raw("SELECT id, name, age FROM users WHERE name = ?", "jinzhu").Scan(&result)

var age int
db.Raw("SELECT SUM(age) FROM users WHERE role = ?", "admin").Scan(&age)

var users []User
db.Raw("UPDATE users SET name = ? WHERE age = ? RETURNING id, name", "jinzhu", 20).Scan(&users)
```

Exec with Raw SQL

```go
db.Exec("DROP TABLE users")
db.Exec("UPDATE orders SET shipped_at = ? WHERE id IN ?", time.Now(), []int64{1, 2, 3})

// Exec with SQL Expression
db.Exec("UPDATE users SET money = ? WHERE name = ?", gorm.Expr("money * ? + ?", 10000, 1), "jinzhu")

```

## DryRun 模式

在不执行的情况下生成 SQL 及其参数，可以用于准备或测试生成的 SQL

```go
stmt := db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement
stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 ORDER BY `id`
stmt.Vars         //=> []interface{}{1}
```

## ToSQL

返回生成的 SQL 但不执行。

GORM使用 database/sql 的参数占位符来构建 SQL 语句，它会自动转义参数以避免 SQL 注入，但我们不保证生成 SQL 的安全，请只用于调试。

```go
sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
  return tx.Model(&User{}).Where("id = ?", 100).Limit(10).Order("age desc").Find(&[]User{})
})
sql //=> SELECT * FROM "users" WHERE id = 100 AND "users"."deleted_at" IS NULL ORDER BY age desc LIMIT 10

```

## Row & Rows

Get result as *sql.Row

```go
// 使用 GORM API 构建 SQL
row := db.Table("users").Where("name = ?", "jinzhu").Select("name", "age").Row()
row.Scan(&name, &age)

// 使用原生 SQL
row := db.Raw("select name, age, email from users where name = ?", "jinzhu").Row()
row.Scan(&name, &age, &email)
```

Get result as *sql.Rows

```go
// 使用 GORM API 构建 SQL
rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows()
defer rows.Close()
for rows.Next() {
  rows.Scan(&name, &age, &email)

  // 业务逻辑...
}

// 原生 SQL
rows, err := db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows()
defer rows.Close()
for rows.Next() {
  rows.Scan(&name, &age, &email)

  // 业务逻辑...
}
```


## 将 sql.Rows 扫描至 model


使用 ScanRows 将一行记录扫描至 struct，例如：

```go
rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows() // (*sql.Rows, error)
defer rows.Close()

var user User
for rows.Next() {
  // ScanRows 将一行扫描至 user
  db.ScanRows(rows, &user)

  // 业务逻辑...
}
```

### 子句（Clause）

GORM 内部使用 SQL builder 生成 SQL。对于每个操作，GORM 都会创建一个 *gorm.Statement 对象，所有的 GORM API 都是在为 statement 添加、修改 子句，最后，GORM 会根据这些子句生成 SQL

例如，当通过 First 进行查询时，它会在 Statement 中添加以下子句

```go
var limit = 1
clause.Select{Columns: []clause.Column{{Name: "*"}}}
clause.From{Tables: []clause.Table{{Name: clause.CurrentTable}}}
clause.Limit{Limit: &limit}
clause.OrderBy{Columns: []clause.OrderByColumn{
  {
    Column: clause.Column{
      Table: clause.CurrentTable,
      Name:  clause.PrimaryKey,
    },
  },
}}
```

然后 GORM 在 Query callback 中构建最终的查询 SQL，像这样：

```go
Statement.Build("SELECT", "FROM", "WHERE", "GROUP BY", "ORDER BY", "LIMIT", "FOR")
```

生成 SQL：
```sql
SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
```


### 子句构造器

不同的数据库, 子句可能会生成不同的 SQL，例如：

```go
db.Offset(10).Limit(5).Find(&users)
// SQL Server 会生成
// SELECT * FROM "users" OFFSET 10 ROW FETCH NEXT 5 ROWS ONLY
// MySQL 会生成
// SELECT * FROM `users` LIMIT 5 OFFSET 10
```

### 子句选项

GORM 定义了很多 子句，其中一些 子句提供了你可能会用到的选项

尽管很少会用到它们，但如果你发现 GORM API 与你的预期不符合。这可能可以很好地检查它们，例如

```go
db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&user)
// INSERT IGNORE INTO users (name,age...) VALUES ("jinzhu",18...);
```

### StatementModifier

---

### 注意

以下关于association（09~15）相关不再关注，因为更推荐软引用！