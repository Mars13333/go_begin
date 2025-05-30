package main

func main() {

}

/*
在日常开发中我们少不了对文件读取，今天我们就从三部分来讲解：全量读，带缓冲区读，任意位置读。
*/

/*
全量读

package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    dat, err := ioutil.ReadFile("./main.go")
    fmt.Println(err)
    fmt.Println(string(dat))
}

运行程序可以打印整个 main.go 文件内容，如果我们将 ./main.go 修改为 ./main.go1，
程序将出现 no such file or directory 的错误， 所以在文件读取的时候一定要注意检查 err。
*/

/*
带缓冲区读

package main

import (
    "fmt"
    "os"
)

func main() {
    f, _ := os.Open("./main.go")
    defer f.Close()

    buf := make([]byte, 16)
    f.Read(buf)

    fmt.Println(string(buf))
}

运行程序会输出 main.go 的前 16 个字节内容，具体为：
package main

im
*/

/*
任意位置读

有些时候我们想在一个文件特定地方读取特定长度的内容，那我们有什么方法可以使用呢？

第一种： f.Seek + f.Read
package main

import (
    "fmt"
    "os"
)

func main() {
    f, _ := os.Open("./main.go")
    defer f.Close()

    b1 := make([]byte, 2)

    f.Seek(5, 0)
    f.Read(b1)
    fmt.Println(string(b1))
}

运行代码输出结果为：

ge


---

第二种：使用 f.ReadAt

package main

import (
    "fmt"
    "os"
)

func main() {
    f, _ := os.Open("./main.go")
    defer f.Close()

    b1 := make([]byte, 2)

    f.ReadAt(b1, 5)
    fmt.Println(string(b1))
}

运行结果同样为：

ge

---


但注意：

第一种方式是非并发安全的，例如：

package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    f, _ := os.Open("./main.go")
    defer f.Close()

    for i := 0; i < 5; i++ {
        go func() {
            b1 := make([]byte, 2)

            f.Seek(5, 0)
            f.Read(b1)
            fmt.Println(string(b1))

            f.Seek(2, 0)
            f.Read(b1)
            fmt.Println(string(b1))
        }()
    }

    time.Sleep(time.Second)
}

输出结果为：
ge
ge
ge
ck
ck
ai
ck
 m
ck
ck


---

第二种 f.ReadAt 是并发安全的，例如：

package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    f, _ := os.Open("./main.go")
    defer f.Close()

    for i := 0; i < 5; i++ {
        go func() {
            b1 := make([]byte, 2)
            f.ReadAt(b1, 5)
            fmt.Println(string(b1))

            f.ReadAt(b1, 2)
            fmt.Println(string(b1))
        }()
    }

    time.Sleep(time.Second)
}

输出结果为：

ge
ge
ck
ck
ge
ge
ck
ck
ge
ck
*/
