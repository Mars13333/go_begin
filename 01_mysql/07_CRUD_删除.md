## 删除一条记录

删除一条记录时，删除对象需要指定主键，否则会触发 批量删除

```go
// Generics API
ctx := context.Background()

// Delete by ID
err := gorm.G[Email](db).Where("id = ?", 10).Delete(ctx)
// DELETE from emails where id = 10;

// Delete with additional conditions
err := gorm.G[Email](db).Where("id = ? AND name = ?", 10, "jinzhu").Delete(ctx)
// DELETE from emails where id = 10 AND name = "jinzhu";
```

```go
// Traditional API
// Email 的 ID 是 `10`
db.Delete(&email)
// DELETE from emails where id = 10;

// 带额外条件的删除
db.Where("name = ?", "jinzhu").Delete(&email)
// DELETE from emails where id = 10 AND name = "jinzhu";

```

## 根据主键删除

GORM 允许通过主键(可以是复合主键)和内联条件来删除对象，它可以使用数字（如以下例子。也可以使用字符串——译者注）。

```go
db.Delete(&User{}, 10)
// DELETE FROM users WHERE id = 10;

db.Delete(&User{}, "10")
// DELETE FROM users WHERE id = 10;

db.Delete(&users, []int{1,2,3})
// DELETE FROM users WHERE id IN (1,2,3);

```

## 钩子函数

对于删除操作，GORM 支持 BeforeDelete、AfterDelete Hook，在删除记录时会调用这些方法，查看 Hook 获取详情

```go
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
    if u.Role == "admin" {
        return errors.New("admin user not allowed to delete")
    }
    return
}
```

## 批量删除

如果指定的值不包括主属性，那么 GORM 会执行批量删除，它将删除所有匹配的记录

```go
db.Where("email LIKE ?", "%jinzhu%").Delete(&Email{})
// DELETE from emails where email LIKE "%jinzhu%";

db.Delete(&Email{}, "email LIKE ?", "%jinzhu%")
// DELETE from emails where email LIKE "%jinzhu%";

```

可以将一个主键切片传递给Delete 方法，以便更高效的删除数据量大的记录

```go
var users = []User{{ID: 1}, {ID: 2}, {ID: 3}}
db.Delete(&users)
// DELETE FROM users WHERE id IN (1,2,3);

db.Delete(&users, "name LIKE ?", "%jinzhu%")
// DELETE FROM users WHERE name LIKE "%jinzhu%" AND id IN (1,2,3); 
```

### 阻止全局删除

当你试图执行不带任何条件的批量删除时，GORM将不会运行并返回ErrMissingWhereClause 错误

如果一定要这么做，你必须添加一些条件，或者使用原生SQL，或者开启AllowGlobalUpdate 模式


## 软删除

如果你的模型包含了 gorm.DeletedAt字段（该字段也被包含在gorm.Model中），那么该模型将会自动获得软删除的能力！

当调用Delete时，GORM并不会从数据库中删除该记录，而是将该记录的DeleteAt设置为当前时间，而后的一般查询方法将无法查找到此条记录。

```go
// user's ID is `111`
db.Delete(&user)
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

// Batch Delete
db.Where("age = ?", 20).Delete(&User{})
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

// Soft deleted records will be ignored when querying
db.Where("age = 20").Find(&user)
// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
```
如果你并不想嵌套gorm.Model，你也可以像下方例子那样开启软删除特性：

```go
type User struct {
  ID      int
  Deleted gorm.DeletedAt
  Name    string
}
```
