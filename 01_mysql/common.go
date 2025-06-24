package main

import (
	"fmt"
	"math/rand"
)

func randomName() string {
	firstName := []string{"张", "李", "王", "赵", "孙", "周", "吴", "郑", "冯", "陈"}
	lastName := []string{"三", "四", "五", "六", "七", "八", "九", "十"}

	// %s%s是格式化模板，表示两个字符串占位符
	name := fmt.Sprintf("%s%s", firstName[rand.Intn(len(firstName))], lastName[rand.Intn(len(lastName))])
	return name
}

func randomAge() int {
	age := rand.Intn(100)
	return age
}
