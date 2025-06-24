## 检索单个对象

GORM 提供了 First、Take、Last 方法，以便从数据库中检索单个对象。当查询数据库时它添加了 LIMIT 1 条件，且没有找到记录时，它会返回 ErrRecordNotFound 错误

```go
// 获取第一条记录（主键升序）
db.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1;

// 获取一条记录，没有指定排序字段
db.Take(&user)
// SELECT * FROM users LIMIT 1;

// 获取最后一条记录（主键降序）
db.Last(&user)
// SELECT * FROM users ORDER BY id DESC LIMIT 1;

result := db.First(&user)
result.RowsAffected // 返回找到的记录数
result.Error        // returns error or nil

// 检查 ErrRecordNotFound 错误
errors.Is(result.Error, gorm.ErrRecordNotFound)
```

如果你想避免ErrRecordNotFound错误，你可以使用Find，比如db.Limit(1).Find(&user)，Find方法可以接受struct和slice的数据。

对单个对象使用Find而不带limit，db.Find(&user)将会查询整个表并且只返回第一个对象，只是性能不高并且不确定的。

First and Last 方法会按主键排序找到第一条记录和最后一条记录 (分别)。 只有在目标 struct 是指针或者通过 db.Model() 指定 model 时，该方法才有效。 此外，如果相关 model 没有定义主键，那么将按 model 的第一个字段进行排序。

```go
var user User
var users []User

// works because destination struct is passed in
db.First(&user)
// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

// works because model is specified using `db.Model()`
result := map[string]interface{}{}
db.Model(&User{}).First(&result)
// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

// doesn't work
result := map[string]interface{}{}
db.Table("users").First(&result)

// works with Take
// map[string]interface{}类型部分，后面一个{}是初始化部分
result := map[string]interface{}{}
db.Table("users").Take(&result)

// no primary key defined, results will be ordered by first field (i.e., `Code`)
type Language struct {
  Code string
  Name string
}
db.First(&Language{})
// SELECT * FROM `languages` ORDER BY `languages`.`code` LIMIT 1
```

### 结构体方式 vs Map 方式的选择

在Go项目中，**结构体方式**是**绝对的主流**，使用频率大概是 **95% vs 5%**。

#### 主流方式：使用结构体（推荐）

```go
var user User
db.First(&user)
```

**为什么这种方式更常用：**

**优势1：类型安全**
```go
var user User
db.First(&user)

// 编译时就能检查字段类型
fmt.Printf("用户名: %s", user.Name)  // ✅ 类型安全
fmt.Printf("年龄: %d", user.Age)     // ✅ 类型安全
```

**优势2：IDE 支持好**
```go
var user User
db.First(&user)

// IDE 有完整的自动补全和类型提示
user.Name     // ✅ IDE 知道这是 string 类型
user.Age      // ✅ IDE 知道这是 int 类型
user.Email    // ✅ IDE 有自动补全
```

**优势3：代码可读性强**
```go
// 一眼就知道要查询什么类型的数据
var user User
db.First(&user)

// 结构体字段一目了然
log.Printf("用户: %+v", user)
```

**优势4：支持所有 GORM 功能**
```go
var user User
db.First(&user)

// 支持钩子函数
func (u *User) AfterFind(tx *gorm.DB) error {
    // 查询后的处理逻辑
    return nil
}

// 支持关联查询
db.Preload("Orders").First(&user)
```

#### Map 方式：特殊场景使用

```go
result := map[string]interface{}{}
db.Model(&User{}).First(&result)
```

**什么时候使用 Map 方式：**

**场景1：动态查询字段**
```go
// 根据用户权限动态决定查询哪些字段
func getUserData(db *gorm.DB, userID uint, isAdmin bool) map[string]interface{} {
    result := map[string]interface{}{}
    
    query := db.Model(&User{})
    if isAdmin {
        // 管理员可以看到所有字段
        query.First(&result)
    } else {
        // 普通用户只能看到部分字段
        query.Select("name", "email").Where("id = ?", userID).First(&result)
    }
    
    return result
}
```

**场景2：API 接口返回**
```go
// 直接返回 JSON 格式的数据
func getUserAPI(c *gin.Context) {
    result := map[string]interface{}{}
    db.Model(&User{}).First(&result)
    
    c.JSON(200, result)  // 直接返回，不需要结构体转换
}
```

**场景3：跨表查询复杂结果**
```go
// 复杂的联表查询，结果不对应单一结构体
result := map[string]interface{}{}
db.Table("users").
   Select("users.name, profiles.avatar, COUNT(orders.id) as order_count").
   Joins("LEFT JOIN profiles ON users.id = profiles.user_id").
   Joins("LEFT JOIN orders ON users.id = orders.user_id").
   Group("users.id").
   First(&result)
```

#### 实际项目中的使用统计

基于实际Go项目观察：

```go
// 95% 的查询使用结构体方式
var user User
db.First(&user)

var users []User
db.Find(&users)

var product Product
db.Where("id = ?", 1).First(&product)

// 只有 5% 的特殊场景使用 map 方式
result := map[string]interface{}{}
db.Model(&User{}).Select("name", "email").First(&result)
```

#### 推荐的最佳实践

**日常开发（推荐结构体）**
```go
// ✅ 推荐：日常 CRUD 操作
func getUserByID(db *gorm.DB, id uint) (*User, error) {
    var user User
    result := db.First(&user, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

// ✅ 推荐：列表查询
func getUsers(db *gorm.DB) ([]User, error) {
    var users []User
    result := db.Find(&users)
    return users, result.Error
}
```

**特殊场景（使用 Map）**
```go
// ✅ 特殊场景：动态字段查询
func getDynamicUserData(db *gorm.DB, fields []string) map[string]interface{} {
    result := map[string]interface{}{}
    db.Model(&User{}).Select(fields).First(&result)
    return result
}
```

**总结：**
- **结构体方式是Go项目的绝对主流**，具有类型安全、IDE支持好、代码可读性强等优势
- **Map 方式只在特殊场景使用**：动态字段查询、API直接返回、复杂跨表查询
- **建议：** 在学习和日常开发中，优先使用结构体方式，只有在确实需要动态性的场景下才考虑 Map 方式

### 根据主键检索

如果主键是数字类型，您可以使用 内联条件 来检索对象。 当使用字符串时，需要额外的注意来避免SQL注入；

```go
db.First(&user, 10)
// SELECT * FROM users WHERE id = 10;

db.First(&user, "10")
// SELECT * FROM users WHERE id = 10;

db.Find(&users, []int{1,2,3})
// SELECT * FROM users WHERE id IN (1,2,3);
```

如果主键是字符串(例如像uuid)，查询将被写成如下：

```go
db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";
```

当目标对象有一个主键值时，将使用主键构建查询条件，例如：

```go
var user = User{ID: 10}
db.First(&user)
// SELECT * FROM users WHERE id = 10;

var result User
db.Model(User{ID: 10}).First(&result)
// SELECT * FROM users WHERE id = 10;
```

如果您使用 gorm 的特定字段类型（例如 gorm.DeletedAt），它将运行不同的查询来检索对象。

```go
type User struct {
  ID           string `gorm:"primarykey;size:16"`
  Name         string `gorm:"size:24"`
  DeletedAt    gorm.DeletedAt `gorm:"index"`
}

var user = User{ID: 15}
db.First(&user)
//  SELECT * FROM `users` WHERE `users`.`id` = '15' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
```

## 检索全部对象

```go
// Get all records
result := db.Find(&users)
// SELECT * FROM users;

result.RowsAffected // returns found records count, equals `len(users)`
result.Error        // returns error
```

## 条件

### string条件


```go
// Get first matched record
db.Where("name = ?", "jinzhu").First(&user)
// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

// Get all matched records
db.Where("name <> ?", "jinzhu").Find(&users)
// SELECT * FROM users WHERE name <> 'jinzhu';

// IN
db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&users)
// SELECT * FROM users WHERE name LIKE '%jin%';

// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

// Time
db.Where("updated_at > ?", lastWeek).Find(&users)
// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

// BETWEEN
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
```

如果对象设置了主键，条件查询将不会覆盖主键的值，而是用 And 连接条件。 例如：

```go
var user = User{ID: 10}
db.Where("id = ?", 20).First(&user)
// SELECT * FROM users WHERE id = 10 and id = 20 ORDER BY id ASC LIMIT 1
```

这个查询将会给出record not found错误 所以，在你想要使用例如 user 这样的变量从数据库中获取新值前，需要将例如 id 这样的主键设置为nil。

### struct & map条件

```go
// Struct
db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

// Map
db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

// Slice of primary keys
db.Where([]int64{20, 21, 22}).Find(&users)
// SELECT * FROM users WHERE id IN (20, 21, 22);
```

零值会被忽略：

```go
db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
// SELECT * FROM users WHERE name = "jinzhu";
```

如果想要零值参与，不能使用结构体的形式，需要使用map的形式：

```go
db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
```

### 指定结构体查询字段

使用struct搜索时，您可以通过将相关字段名或dbname传递给Where（）来指定在查询条件中使用结构中的哪些特定值，例如：

```go
db.Where(&User{Name: "jinzhu"}, "name", "Age").Find(&users)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

db.Where(&User{Name: "jinzhu"}, "Age").Find(&users)
// SELECT * FROM users WHERE age = 0;
```

### 内联条件

查询条件可以以类似于Where的方式内联到First和Find等方法中。(而不是写在where中)

```go
// Get by primary key if it were a non-integer type
db.First(&user, "id = ?", "string_primary_key")
// SELECT * FROM users WHERE id = 'string_primary_key';

// Plain SQL
db.Find(&user, "name = ?", "jinzhu")
// SELECT * FROM users WHERE name = "jinzhu";

db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

// Struct
db.Find(&users, User{Age: 20})
// SELECT * FROM users WHERE age = 20;

// Map
db.Find(&users, map[string]interface{}{"age": 20})
// SELECT * FROM users WHERE age = 20;
```

### Not条件

```go
db.Not("name = ?", "jinzhu").First(&user)
// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;

// Not In
db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

// Struct
db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

// Not In slice of primary keys
db.Not([]int64{1,2,3}).First(&user)
// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
```

### Or条件

```go

db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

// Struct
db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

// Map
db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);
```

## 选择特定字段

选择允许您指定要从数据库中检索的字段。否则，GORM将默认选择所有字段。

```go
db.Select("name", "age").Find(&users)
// SELECT name, age FROM users;

db.Select([]string{"name", "age"}).Find(&users)
// SELECT name, age FROM users;

db.Table("users").Select("COALESCE(age,?)", 42).Rows()
// SELECT COALESCE(age,'42') FROM users;
```

## 排序

```go
db.Order("age desc, name").Find(&users)
// SELECT * FROM users ORDER BY age desc, name;

// Multiple orders
db.Order("age desc").Order("name").Find(&users)
// SELECT * FROM users ORDER BY age desc, name;

db.Clauses(clause.OrderBy{
  Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
}).Find(&User{})
// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
```

## limit & offset

```go
db.Limit(3).Find(&users)
// SELECT * FROM users LIMIT 3;

// Cancel limit condition with -1
db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
// SELECT * FROM users LIMIT 10; (users1)
// SELECT * FROM users; (users2)

db.Offset(3).Find(&users)
// SELECT * FROM users OFFSET 3;

db.Limit(10).Offset(5).Find(&users)
// SELECT * FROM users OFFSET 5 LIMIT 10;

// Cancel offset condition with -1
db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
// SELECT * FROM users OFFSET 10; (users1)
// SELECT * FROM users; (users2)
```

## group by & having

```go
type result struct {
  Date  time.Time
  Total int
}

db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&result)
// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1


db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"

rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
defer rows.Close()
for rows.Next() {
  ...
}

rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
defer rows.Close()
for rows.Next() {
  ...
}

type Result struct {
  Date  time.Time
  Total int64
}
db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)
```

## distinct

```go
db.Distinct("name", "age").Order("name, age desc").Find(&results)
```

## joins

```go
type result struct {
  Name  string
  Email string
}

db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
for rows.Next() {
  ...
}

db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

// multiple joins with parameter
db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)
```

### joins预加载

可以将Joins与单个SQL结合使用，例如：

```go
db.Joins("Company").Find(&users)
// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;

// inner join
db.InnerJoins("Company").Find(&users)
// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` INNER JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;
```

**join & where**

```go
db.Joins("Company", db.Where(&Company{Alive: true})).Find(&users)
// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id` AND `Company`.`alive` = true;
```

### joins一个衍生表

```go
type User struct {
    Id  int
    Age int
}

type Order struct {
    UserId     int
    FinishedAt *time.Time
}

query := db.Table("order").Select("MAX(order.finished_at) as latest").Joins("left join user user on order.user_id = user.id").Where("user.age > ?", 18).Group("order.user_id")
db.Model(&Order{}).Joins("join (?) q on order.finished_at = q.latest", query).Scan(&results)
// SELECT `order`.`user_id`,`order`.`finished_at` FROM `order` join (SELECT MAX(order.finished_at) as latest FROM `order` left join user user on order.user_id = user.id WHERE user.age > 18 GROUP BY `order`.`user_id`) q on order.finished_at = q.latest
```

## scan

将结果扫描到结构体中的工作方式与我们使用Find的方式类似

```go
type Result struct {
  Name string
  Age  int
}

var result Result
db.Table("users").Select("name", "age").Where("name = ?", "Antonio").Scan(&result)

// Raw SQL
db.Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)
```