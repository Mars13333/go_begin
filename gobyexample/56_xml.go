package main

import (
	"encoding/xml"
	"fmt"
)

// Plant结构体：定义XML标签和属性
type Plant struct {
	XMLName xml.Name `xml:"plant"`   // XML根元素名称
	Id      int      `xml:"id,attr"` // XML属性，attr表示这是属性而不是子元素
	Name    string   `xml:"name"`    // XML子元素
	Origin  []string `xml:"origin"`  // XML子元素，切片会生成多个同名元素
}

// String方法：自定义Plant结构体的字符串表示
func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	// 创建Plant实例
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// 序列化Plant为XML，使用MarshalIndent格式化输出
	// 参数：要序列化的对象、前缀、缩进
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// 添加XML声明头并输出
	fmt.Println(xml.Header + string(out))

	// 反序列化XML到Plant结构体
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	// 创建另一个Plant实例
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// 定义嵌套结构体，演示XML嵌套
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`            // 根元素
		Plants  []*Plant `xml:"parent>child>plant"` // 嵌套路径：parent/child/plant
	}

	// 创建嵌套结构实例
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	// 序列化嵌套结构为XML
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}

/*
Go语言XML处理示例 - encoding/xml包使用

主旨：
1. 演示encoding/xml包中XML序列化和反序列化的方法
2. 展示XML标签和属性的定义方式
3. 理解XML嵌套结构的处理
4. 学习如何自定义结构体的字符串表示

关键特性：
- encoding/xml：Go标准库的XML处理包
- 支持结构体、切片、基本类型的序列化/反序列化
- 提供XML标签自定义元素名和属性
- 支持XML嵌套和复杂结构

XML标签语法：
- `xml:"element_name"`：自定义XML元素名
- `xml:"attr_name,attr"`：定义XML属性
- `xml:"parent>child"`：定义嵌套路径
- `xml:",chardata"`：元素内容为字符数据
- `xml:",innerxml"`：元素内容为原始XML
- `xml:",comment"`：元素内容为注释

序列化函数：
- xml.Marshal(v)：将Go值序列化为XML字节切片
- xml.MarshalIndent(v, prefix, indent)：格式化输出XML

反序列化函数：
- xml.Unmarshal(data, v)：将XML字节切片反序列化为Go值

输出示例：
<plant id="27">
  <name>Coffee</name>
  <origin>Ethiopia</origin>
  <origin>Brazil</origin>
</plant>

<?xml version="1.0" encoding="UTF-8"?>
<plant id="27">
  <name>Coffee</name>
  <origin>Ethiopia</origin>
  <origin>Brazil</origin>
</plant>

Plant id=27, name=Coffee, origin=[Ethiopia Brazil]

<nesting>
  <parent>
    <child>
      <plant id="27">
        <name>Coffee</name>
        <origin>Ethiopia</origin>
        <origin>Brazil</origin>
      </plant>
      <plant id="81">
        <name>Tomato</name>
        <origin>Mexico</origin>
        <origin>California</origin>
      </plant>
    </child>
  </parent>
</nesting>

XML标签高级用法：
```go
type Person struct {
    XMLName xml.Name `xml:"person"`
    ID      int      `xml:"id,attr"`
    Name    string   `xml:"name"`
    Age     int      `xml:"age,omitempty"`
    Bio     string   `xml:"bio,chardata"`
    Comment string   `xml:"comment,comment"`
}
```

数据类型映射：
- bool → XML boolean
- int/float → XML number
- string → XML text
- slice → 多个同名XML元素
- struct → XML元素
- xml.Name → XML元素名

实际应用场景：
- Web服务SOAP协议
- 配置文件读写
- RSS/Atom订阅源
- 数据交换格式
- 文档处理

注意事项!!!!：
- XMLName字段用于指定根元素名
- 属性使用attr标签
- 嵌套路径使用>分隔
- 切片会生成多个同名元素
- 结构体字段必须大写才能被序列化
- XML声明头需要手动添加
*/
