package main

import (
    "fmt"
)

func main() {

    //定义  长度是数组的一部分
    var a [10]int

    a[0] = 1

    //遍历
    //
    for k, v := range a {

        fmt.Println(k, v)

    }
}
