package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

// 自定义表名
// 因为 GORM 默认使用结构体名称的复数形式作为表名，但你的表名是 02_user。 学习所需，每个案例不一致，所以需要自定义表名方便学习掌握
// 值接收者 - 不需要实例数据，所以不需要使用(u User) 是一种语法糖，等同于 func (u User) TableName() string { return "02_user" }
func (User) TableName() string {
	return "02_user"
}

func main() {
	// 数据库连接信息
	// mysql 8.0 版本需要添加 allowNativePasswords=true 参数
	dsn := "root:123457@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true"

	// 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自动迁移（如果表结构已存在，可以跳过）
	// db.AutoMigrate(&User{})

	// 示例，查看所有用户
	var users []User
	db.Find(&users)
	for _, user := range users {
		log.Printf("User: %+v", user)
	}
}

/*
为什么 db.Find 中要用指针？

1. GORM 需要修改传入的切片来填充查询结果
   - var users []User 初始化为 nil 切片
   - db.Find(&users) GORM 需要向这个切片添加数据

2. Go 语言的值传递特性
   - db.Find(users) 传递的是切片的副本，GORM 无法修改原始切片
   - db.Find(&users) 传递切片的指针，GORM 可以修改原始切片

3. GORM 内部实现过程：
   - 执行 SQL 查询
   - 创建结构体实例
   - 填充数据
   - 将实例追加到 dest 指向的切片中
   - 这需要 dest 是指针，否则无法修改原始切片

4. 类比理解：
   func appendToSlice(slice []int) {
       slice = append(slice, 1, 2, 3)  // 不会修改原始切片
   }

   func appendToSlicePtr(slice *[]int) {
       *slice = append(*slice, 1, 2, 3)  // 会修改原始切片
   }

总结：因为 GORM 需要向切片中添加查询结果，所以必须传递指针，否则无法修改原始切片。
*/
