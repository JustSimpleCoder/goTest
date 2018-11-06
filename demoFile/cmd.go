package main

import (
    "fmt"
    "os"
)

type s struct {
    Name string
    Age  int
}

func main() {

    var str = "y 10"

    var stu s
    fmt.Sscanf(str, "%s %d", &stu.Name, &stu.Age)
    fmt.Println(stu)

    var a = "abc"

    // 打印到终端||文件
    fmt.Fprintf(os.Stdout, a)

    file, err := os.OpenFile("c:/Users/yaobin/go/test.log", os.O_CREATE|os.O_APPEND, 0777)
    if err != nil {

        fmt.Printf("Open File error ", err)
        return
    }

    fmt.Fprintf(file, "abc\r\n")

}
