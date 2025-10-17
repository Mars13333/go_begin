## 保存所有字段

save会保存所有字段，即使字段是零值

```go
db.First(&user)

user.Name = "jinzhu 2"
user.Age = 100
db.Save(&user)
// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
```

## 更新单个列

当使用 `Update` 更新单列时，需要有一些条件，否则将会引起`ErrMissingWhereClause` 错误，查看 阻止全局更新 了解详情。

```go
// Generics API
ctx := context.Background()

// Update with conditions
err := gorm.G[User](db).Where("active = ?", true).Update(ctx, "name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

// Update with ID condition
err := gorm.G[User](db).Where("id = ?", 111).Update(ctx, "name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

// Update with multiple conditions
err := gorm.G[User](db).Where("id = ? AND active = ?", 111, true).Update(ctx, "name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
```

当使用 Model 方法，并且它有主键值时，主键将会被用于构建条件，例如：

```go
// Traditional API
// 根据条件更新
db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

// User 的 ID 是 `111`
db.Model(&user).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

// 根据条件和 model 的值进行更新
db.Model(&user).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

```


## 更新多列

Updates 方法支持 `struct` 和 `map[string]interface{}` 参数。当使用 struct 更新时，默认情况下GORM 只会更新非零值的字段

```go
// Generics API
ctx := context.Background()

// Update attributes with `struct`, will only update non-zero fields
rows, err := gorm.G[User](db).Where("id = ?", 111).Updates(ctx, User{Name: "hello", Age: 18, Active: false})
// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

// Update attributes with `map`
rows, err := gorm.G[User](db).Where("id = ?", 111).Updates(ctx, map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
```

```go
// Traditional API
// 根据 `struct` 更新属性，只会更新非零值的字段
db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

// 根据 `map` 更新属性
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

```

## 更新选定字

如果您想要在更新时选择、忽略某些字段，您可以使用 `Select`、`Omit`

```go
// Generics API
ctx := context.Background()

// Select with Map
rows, err := gorm.G[User](db).Where("id = ?", 111).Select("name").Updates(ctx, map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello' WHERE id=111;

rows, err := gorm.G[User](db).Where("id = ?", 111).Omit("name").Updates(ctx, map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

// Select with Struct (select zero value fields)
rows, err := gorm.G[User](db).Where("id = ?", 111).Select("Name", "Age").Updates(ctx, User{Name: "new_name", Age: 0})
// UPDATE users SET name='new_name', age=0 WHERE id=111;

// Select all fields (select all fields include zero value fields)
rows, err := gorm.G[User](db).Where("id = ?", 111).Select("*").Updates(ctx, User{Name: "jinzhu", Role: "admin", Age: 0})

// Select all fields but omit Role (select all fields include zero value fields)
rows, err := gorm.G[User](db).Where("id = ?", 111).Select("*").Omit("Role").Updates(ctx, User{Name: "jinzhu", Role: "admin", Age: 0})
```

```go
// Traditional API
// 选择 Map 的字段
// User 的 ID 是 `111`:
db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello' WHERE id=111;

db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

// 选择 Struct 的字段（会选中零值的字段）
db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
// UPDATE users SET name='new_name', age=0 WHERE id=111;

// 选择所有字段（选择包括零值字段的所有字段）
db.Model(&user).Select("*").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})

// 选择除 Role 外的所有字段（包括零值字段的所有字段）
db.Model(&user).Select("*").Omit("Role").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})
```

## 更新 Hook

GORM 支持的 hook 包括：BeforeSave, BeforeUpdate, AfterSave, AfterUpdate. 更新记录时将调用这些方法，查看 Hooks 获取详细信息

```go
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
    if u.Role == "admin" {
        return errors.New("admin user not allowed to update")
    }
    return
}
```


## 批量更新

如果没有通过 `Model` 指定一个含有主键的记录，GORM 会执行批量更新

```go
// Update with struct
db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE role = 'admin';

// Update with map
db.Table("users").Where("id IN ?", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);
```

