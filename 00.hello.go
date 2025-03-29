package main // 创建文件默认的值是go_begin,改成main之后才能run

func main() {
	hello("疯狂的世界")
}

func hello(str string) {
	println("hello! ", str)
}
