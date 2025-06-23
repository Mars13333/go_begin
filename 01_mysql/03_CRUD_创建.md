## 创建记录

```go
user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

result := db.Create(&user) // 通过数据的指针来创建

user.ID             // 返回插入数据的主键
result.Error        // 返回 error
result.RowsAffected // 返回插入记录的条数
```

还可以使用 Create() 创建多项记录：

```go
users := []*User{
    {Name: "Jinzhu", Age: 18, Birthday: time.Now()},
    {Name: "Jackson", Age: 19, Birthday: time.Now()},
}

result := db.Create(users) // pass a slice to insert multiple row

result.Error        // returns error
result.RowsAffected // returns inserted records count
```

注意：无法向 ‘create’ 传递结构体，应该传入数据的指针.

## 用指定的字段创建记录

创建记录并为指定字段赋值。

```go
db.Select("Name", "Age", "CreatedAt").Create(&user)
// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
```

创建记录并忽略传递给 ‘Omit’ 的字段值

```go
db.Omit("Name", "Age", "CreatedAt").Create(&user)
// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")
```

此处的 Select 不是 SQL 的 SELECT 查询，而是 GORM 的字段选择器。

- 1. 在 Create 操作中 - 选择要插入的字段
```go
db.Select("Name", "Age", "CreatedAt").Create(&user)
// 只插入 Name、Age、CreatedAt 这三个字段
// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
```
- 2. 在 Find 操作中 - 选择要查询的字段
```go
db.Select("Name", "Age").Find(&users)
// 只查询 Name 和 Age 字段
// SELECT `name`,`age` FROM `users`
```
- 3.其他 GORM 字段选择器
```go
// Omit - 排除指定字段
db.Omit("Email", "Password").Create(&user)

// Select - 选择指定字段
db.Select("Name", "Age").Create(&user)

// 组合使用
db.Select("Name", "Age").Omit("ID").Create(&user)
```

## 批量插入

要高效地插入大量记录，请将切片传递给Create方法。 GORM 将生成一条 SQL 来插入所有数据，以返回所有主键值，并触发 Hook 方法。 当这些记录可以被分割成多个批次时，GORM会开启一个**事务**来处理它们。

```go
var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
db.Create(&users)

for _, user := range users {
  user.ID // 1,2,3
}
```

还可以通过`db.CreateInBatches`方法来指定批量插入的批次大小

```go
var users = []User{{Name: "jinzhu_1"}, ...., {Name: "jinzhu_10000"}}

// batch size 100
db.CreateInBatches(users, 100)
```

## 创建钩子

GORM允许用户通过实现这些接口 BeforeSave, BeforeCreate, AfterSave, AfterCreate来自定义钩子。 这些钩子方法会在创建一条记录时被调用。

```go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.UUID = uuid.New()

    if u.Role == "admin" {
        return errors.New("invalid role")
    }
    return
}
```

如果你想跳过Hooks方法，可以使用SkipHooks会话模式，例子如下

```go
DB.Session(&gorm.Session{SkipHooks: true}).Create(&user)

DB.Session(&gorm.Session{SkipHooks: true}).Create(&users)

DB.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(users, 100)

```

## 根据Map创建

GORM支持通过 `map[string]interface{}` 与 `[]map[string]interface{}{}`来创建记录。

```go
db.Model(&User{}).Create(map[string]interface{}{
  "Name": "jinzhu", "Age": 18,
})

// batch insert from `[]map[string]interface{}{}`
db.Model(&User{}).Create([]map[string]interface{}{
  {"Name": "jinzhu_1", "Age": 18},
  {"Name": "jinzhu_2", "Age": 20},
})
```

## 使用SQL表达式、Context Valuer创建记录

GORM允许使用SQL表达式来插入数据，有两种方法可以达成该目的，使用`map[string]interface{}` 或者 `Customized Data Types`， 示例如下：

```go
// Create from map
db.Model(User{}).Create(map[string]interface{}{
  "Name": "jinzhu",
  "Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 100)"}},
})
// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"));

// Create from customized data type
type Location struct {
    X, Y int
}

// Scan implements the sql.Scanner interface
func (loc *Location) Scan(v interface{}) error {
  // Scan a value into struct from database driver
}

func (loc Location) GormDataType() string {
  return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
  return clause.Expr{
    SQL:  "ST_PointFromText(?)",
    Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
  }
}

type User struct {
  Name     string
  Location Location
}

db.Create(&User{
  Name:     "jinzhu",
  Location: Location{X: 100, Y: 100},
})
// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"))
```

## 高级选项

### 关联创建

创建关联数据时，如果关联值非零，这些关联会被upsert，并且它们的Hooks方法也会被调用。

```go
type CreditCard struct {
  gorm.Model
  Number   string
  UserID   uint
}

type User struct {
  gorm.Model
  Name       string
  CreditCard CreditCard
}

db.Create(&User{
  Name: "jinzhu",
  CreditCard: CreditCard{Number: "411111111111"},
})
// INSERT INTO `users` ...
// INSERT INTO `credit_cards` ...
```

可以通过Select, Omit方法来跳过关联更新，示例如下：

```go
db.Omit("CreditCard").Create(&user)

// skip all associations
db.Omit(clause.Associations).Create(&user)

```

### 默认值

可以通过结构体Tag default来定义字段的默认值，示例如下：

```go
type User struct {
  ID   int64
  Name string `gorm:"default:galeone"`
  Age  int64  `gorm:"default:18"`
}
```

这些默认值会被当作结构体字段的零值插入到数据库中

注意，当结构体的字段默认值是零值的时候比如 0, '', false，这些字段值将不会被保存到数据库中，你可以使用指针类型或者Scanner/Valuer来避免这种情况。

```go
type User struct {
  gorm.Model
  Name string
  Age  *int           `gorm:"default:18"`
  Active sql.NullBool `gorm:"default:true"`
}
```

注意，若要让字段在数据库中拥有默认值则必须使用defaultTag来为结构体字段设置默认值。如果想要在数据库迁移的时候跳过默认值，可以使用 default:(-)，示例如下：

```go
type User struct {
  ID        string `gorm:"default:uuid_generate_v3()"` // db func
  FirstName string
  LastName  string
  Age       uint8
  FullName  string `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`
}
```
