package main

func main() {

}

/*
Go 写文件的相关操作
	创建文件
	在文件指定位置写入内容
	通过 Buffered Writter 写文件

*/

/*
创建文件
使用 os.Create(name string) 方法创建文件
使用 os.Stat(name string) 方法获取文件信息

package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    path := "test.txt"

    newFile, err := os.Create(path)
    checkErr(err)
    defer newFile.Close()

    fileInfo, err := os.Stat(path)
    if err != nil {
        if os.IsNotExist(err) {
            fmt.Println("file doesn't exist!")
            return
        }
    }
    fmt.Println("file does exist, file name : ", fileInfo.Name())
}
*/

/*
在文件指定位置写入内容
使用 writeAt 可以在文件指定位置写入内容
package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func readFile(path string) {
    data, err := ioutil.ReadFile(path)
    checkErr(err)
    fmt.Println("file content: ", string(data))
}

func main() {
    path := "test.txt"
    str := "hello"
    newStr := "world"

    newFile, err := os.Create(path)
    checkErr(err)

    n1, err := newFile.WriteString(str)
    checkErr(err)
    fmt.Println("n1: ", n1)
    readFile(path)

    n2, err := newFile.WriteAt([]byte(newStr), 6)
    checkErr(err)
    fmt.Println("n2: ", n2)
    readFile(path) // the file content should be "helloworld"

    n3, err := newFile.WriteAt([]byte(newStr), 0)
    checkErr(err)
    fmt.Println("n3: ", n3)
    readFile(path) // the file content should be "worldworld"
}


*/

/*
通过 Buffered Writer 写文件

使用 Buffered Writer 可以避免太多次的磁盘 IO 操作。写入的内容首先是存在内存中，当调用 Flush() 方法后才会写入磁盘。

package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func readFile(path string) {
    data, err := ioutil.ReadFile(path)
    checkErr(err)
    fmt.Println("file content: ", string(data))
}

func main() {
    path := "test.txt"
    str := "hello"

    newFile, err := os.Create(path)
    checkErr(err)
    defer newFile.Close()

    bufferWriter := bufio.NewWriter(newFile)

    for _, v := range str {
        written, err := bufferWriter.WriteString(string(v))
        checkErr(err)
        fmt.Println("written: ", written)
    }

    readFile(path) // NOTE: you'll read nothing here because without Flush() operation

    // let's check how much is stored in buffer
    unflushSize := bufferWriter.Buffered()
    fmt.Println("unflushSize: ", unflushSize)

    // write memory buffer to disk
    bufferWriter.Flush()

    readFile(path) // now you can get content from file
}
*/
