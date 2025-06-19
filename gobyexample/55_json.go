package main

import (
	"encoding/json" // 导入JSON处理包
	"fmt"
	"os"
	"strings"
)

// response1结构体：没有JSON标签，字段名直接作为JSON键名
type response1 struct {
	Page   int      // 序列化为 "Page"
	Fruits []string // 序列化为 "Fruits"
}

// response2结构体：使用JSON标签自定义JSON键名
type response2 struct {
	Page   int      `json:"page"`   // 序列化为 "page"
	Fruits []string `json:"fruits"` // 序列化为 "fruits"
}

func main() {

	// 序列化基本数据类型为JSON
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// 序列化切片为JSON数组
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	// 序列化map为JSON对象
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// 序列化结构体（无JSON标签）
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 序列化结构体（有JSON标签）
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// 反序列化JSON到map[string]interface{}
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 从map中提取值，需要类型断言
	num := dat["num"].(float64) // JSON数字默认解析为float64
	fmt.Println(num)

	strs := dat["strs"].([]interface{}) // JSON数组解析为[]interface{}
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 反序列化JSON到结构体
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 使用JSON编码器直接输出到io.Writer
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// 使用JSON解码器从io.Reader读取JSON
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}

/*
Go语言JSON处理示例 - encoding/json包使用

主旨：
1. 演示encoding/json包中JSON序列化和反序列化的方法
2. 展示结构体JSON标签的使用和效果
3. 理解JSON编码器和解码器的用法
4. 学习如何处理不同类型的JSON数据

关键特性：
- encoding/json：Go标准库的JSON处理包
- 支持结构体、map、切片、基本类型的序列化/反序列化
- 提供JSON标签自定义字段名
- 支持流式JSON处理（编码器/解码器）

JSON标签语法：
- `json:"field_name"`：自定义JSON字段名
- `json:"field_name,omitempty"`：空值时忽略该字段
- `json:"-"`：忽略该字段，不参与JSON处理
- `json:",string"`：将数字类型序列化为字符串

序列化函数：
- json.Marshal(v)：将Go值序列化为JSON字节切片
- json.MarshalIndent(v, prefix, indent)：格式化输出JSON

反序列化函数：
- json.Unmarshal(data, v)：将JSON字节切片反序列化为Go值

流式处理：
- json.NewEncoder(w)：创建JSON编码器，输出到io.Writer
- json.NewDecoder(r)：创建JSON解码器，从io.Reader读取

输出示例：
true
1
2.34
"gopher"
["apple","peach","pear"]
{"apple":5,"lettuce":7}
{"Page":1,"Fruits":["apple","peach","pear"]}
{"page":1,"fruits":["apple","peach","pear"]}
map[num:6.13 strs:[a b]]
6.13
a
{1 [apple peach]}
apple
{"apple":5,"lettuce":7}
{1 [apple peach]}

数据类型映射：
- bool → JSON boolean
- int/float → JSON number
- string → JSON string
- slice → JSON array
- map/struct → JSON object
- nil → JSON null

JSON标签高级用法：
```go
type Person struct {
    Name     string `json:"name"`
    Age      int    `json:"age,omitempty"`
    Password string `json:"-"`                    // 忽略字段
    Score    int    `json:"score,string"`         // 数字转字符串
}
```

实际应用场景：
- Web API数据交换
- 配置文件读写
- 数据持久化
- 跨语言数据交换
- 日志记录

注意事项!!!!!!：
- JSON数字默认解析为float64，需要类型断言
- 结构体字段必须大写才能被序列化
- 反序列化时JSON字段名不区分大小写
- 使用omitempty标签可以忽略零值字段
- 编码器/解码器适合处理大文件或流式数据
*/
