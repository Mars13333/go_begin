package main

import (
	"errors"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User3 struct {
	gorm.Model // 包含 ID, CreatedAt, UpdatedAt, DeletedAt 字段
	Name       string
	Age        int
	// Birthday   time.Time
	Birthday *time.Time
}

// 自定义表名
// 因为 GORM 默认使用结构体名称的复数形式作为表名，但你的表名是 02_user。 学习所需，每个案例不一致，所以需要自定义表名方便学习掌握
// 值接收者 - 不需要实例数据，所以不需要使用(u User) 是一种语法糖，等同于 func (u User) TableName() string { return "02_user" }
func (User3) TableName() string {
	return "user3"
}

// 勾子
func (u *User3) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Age < 18 {
		return errors.New("age must be greater than 18")
	}
	return nil
}

func main3() {
	dsn := "root:123457@tcp(172.17.0.2:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自动迁移（如果表结构已存在，可以跳过，如果不存在则在数据库中生成）
	db.AutoMigrate(&User3{})

	// createSingleUser(db)
	// createMultiUsers(db)
	// createWithSpecified(db)
	// ignoreWithSpecified(db)
	// createInBatches(db)
	// createWithMap(db)
	createMapBatches(db)
}

func createSingleUser(db *gorm.DB) {
	// random name & random age
	// go run 03_CRUD_创建.go common.go 同时指定多个文件。common里面有工具函数。
	// go run . 不合适。这是运行整个包，还在练习阶段。推荐上面的方式。
	name := randomName()
	age := randomAge()

	// 创建
	now := time.Now()
	user := User3{Name: name, Age: age, Birthday: &now}

	result := db.Create(&user) // 通过数据的指针来创建

	if result.Error != nil {
		log.Fatalf("failed to create user: %v", result.Error)
	}

	// log.Printf("user created: %v", user)
	// 输出：{1 2024-01-15... 张三 20 2024-01-15...}
	// log.Println("--------------------------------")
	log.Printf("user created: %+v", user)
	// 输出：{ID:1 CreatedAt:2024-01-15... Name:张三 Age:20 Birthday:2024-01-15...}
	// log.Println("--------------------------------")
	// log.Printf("user created: %#v", user)
	// 输出：main.User3{Model:gorm.Model{ID:0x1, CreatedAt:time.Time{...}, Name:"张三", Age:20, Birthday:time.Time{...}}
}

func createMultiUsers(db *gorm.DB) {
	now := time.Now()
	users := []*User3{
		{Name: randomName(), Age: randomAge(), Birthday: &now},
		{Name: randomName(), Age: randomAge(), Birthday: &now},
		{Name: randomName(), Age: randomAge(), Birthday: &now},
	}
	result := db.Create(users)
	if result.Error != nil {
		log.Fatalf("failed to create users: %v", result.Error)
	}
	log.Printf("users created: %+v", users)
}

// 用指定字段创建记录
func createWithSpecified(db *gorm.DB) {
	now := time.Now()
	// select 非SQL的select ，而是选择器
	db.Select("Name", "Age").Create(&User3{Name: "张三", Age: 20, Birthday: &now})
}

// 用指定字段创建记录 注：select可以和omit组合使用
func ignoreWithSpecified(db *gorm.DB) {
	now := time.Now()
	// omit 忽略字段
	db.Omit("Age").Create(&User3{Name: "张三", Age: 20, Birthday: &now})
}

// 批量插入 大于1000条的情况再使用。其他可以直接使用Create即可。
// 这些记录可以被分割成多个批次时，GORM会开启一个事务来处理它们。
func createInBatches(db *gorm.DB) {
	users := []User3{
		{Name: randomName()},
		{Name: randomName()},
		{Name: randomName()},
	}
	db.CreateInBatches(users, 2)
}

// 使用map
// 勾子函数不会被调用，而使用结构体的db.Create()会调用。
// 更灵活，如果是API接口处理，推荐使用map形式，因为使用结构体需要先创建结构体，再赋值。
func createWithMap(db *gorm.DB) {
	now := time.Now()
	user := map[string]interface{}{
		"Name":      "张三啊啊",
		"Age":       28,
		"Birthday":  &now,
		"CreatedAt": now, // 显式添加 当使用 map[string]interface{} 创建记录时，GORM 无法识别 map 中的字段对应关系
		"UpdatedAt": now, // 显式添加
	}
	db.Model(&User3{}).Create(&user)
}

// map[string]interface{}  这是映射
// []map[string]interface{}  这是(映射)切片
func createMapBatches(db *gorm.DB) {
	now := time.Now()
	users := []map[string]interface{}{
		{"Name": "jinzhu_11", "Age": 18, "Birthday": &now},
		{"Name": "jinzhu_22", "Age": 20, "Birthday": &now},
	}
	db.Model(&User3{}).Create(users)
}

// 关联创建
/*
何时跳过关联：
关联对象已存在
需要分步创建
有复杂的业务逻辑
需要手动控制事务
最佳实践：
简单场景：让 GORM 自动处理关联
复杂场景：手动控制事务和关联创建
明确你的业务需求，选择合适的方式
*/

// 默认值
/*
优点
简化代码：不需要每次都手动设置常用值
数据一致性：确保字段有合理的初始值
减少错误：避免忘记设置重要字段
数据库约束：在数据库层面保证数据完整性
缺点
隐藏逻辑：默认值可能不够明显
维护成本：修改默认值需要考虑历史数据
调试困难：有时不清楚值是手动设置还是默认值
// ✅ 推荐：状态、标志、计数器类字段
type User struct {
    Status    string `gorm:"default:active"`
    IsDeleted bool   `gorm:"default:false"`
    LoginCount int   `gorm:"default:0"`
}

// ✅ 推荐：业务规则明确的字段
type Order struct {
    Status string `gorm:"default:pending"`
    Total  decimal.Decimal `gorm:"default:0"`
}

// ✅ 高频场景
type User struct {
    gorm.Model
    Username  string
    Email     string
    // 这些字段经常使用默认值
    Status    string    `gorm:"default:active"`      // 95% 的项目都会用
    IsDeleted bool      `gorm:"default:false"`       // 软删除标记
    CreatedBy uint      `gorm:"default:0"`           // 创建者ID
    Role      string    `gorm:"default:user"`        // 用户角色
    Points    int       `gorm:"default:0"`           // 积分系统
}
*/
