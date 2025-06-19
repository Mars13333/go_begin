package main

import (
	"os"
	"text/template"
)

func main() {

	// 创建新模板，名称为"t1"
	t1 := template.New("t1")
	// 解析模板字符串，{{.}}表示当前值
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// 使用template.Must简化错误处理，如果解析失败会panic
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// 执行模板，输出字符串
	t1.Execute(os.Stdout, "some text")
	// 执行模板，输出整数
	t1.Execute(os.Stdout, 5)
	// 执行模板，输出字符串切片
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// 创建辅助函数，简化模板创建过程
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// 创建模板t2，使用结构体字段名
	t2 := Create("t2", "Name: {{.Name}}\n")

	// 执行模板，传入结构体
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// 执行模板，传入map
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// 创建模板t3，使用条件语句
	// {{if . -}} 如果当前值不为空则输出"yes"，否则输出"no"
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty") // 输出：yes
	t3.Execute(os.Stdout, "")          // 输出：no

	// 创建模板t4，使用range循环
	// {{range .}} 遍历切片或数组，{{.}}表示当前元素
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}

/*
Go语言文本模板示例 - text/template包使用

主旨：
1. 演示text/template包的基本使用方法
2. 展示模板的创建、解析和执行过程
3. 理解模板语法：变量、条件、循环等
4. 学习如何将数据与模板结合生成文本

关键特性：
- text/template：Go标准库的文本模板引擎
- 支持变量替换、条件判断、循环等语法
- 可以处理各种数据类型：字符串、数字、切片、结构体、map等
- 提供安全的模板执行机制

模板语法详解：

1. 变量输出：
   - {{.}} : 输出当前值
   - {{.FieldName}} : 输出结构体字段或map键值

2. 条件语句：
   - {{if .}} ... {{end}} : 如果当前值不为空则执行
   - {{if .}} ... {{else}} ... {{end}} : 带else的条件语句

3. 循环语句：
   - {{range .}} ... {{end}} : 遍历切片、数组或map
   - {{range $key, $value := .}} ... {{end}} : 带键值的遍历

4. 模板函数：
   - template.New(name) : 创建新模板
   - template.Must() : 简化错误处理，失败时panic
   - Execute(writer, data) : 执行模板

数据绑定：
- 结构体：通过字段名访问 {{.Name}}
- Map：通过键名访问 {{.Key}}
- 切片/数组：通过索引访问 {{index . 0}}
- 基本类型：直接使用 {{.}}

输出示例：
Value: some text
Value: 5
Value: [Go Rust C++ C#]
Name: Jane Doe
Name: Mickey Mouse
yes
no
Range: Go Rust C++ C#

实际应用场景：
- 生成HTML页面
- 创建配置文件
- 生成邮件内容
- 代码生成器
- 报告生成

注意事项：
- 模板语法使用双大括号 {{}}
- 模板执行是安全的，不会执行任意代码
- 可以使用template.Must简化错误处理
- 支持模板嵌套和复用
*/
