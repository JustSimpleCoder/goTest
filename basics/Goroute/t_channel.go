package main

import (
    "fmt"
)

var m1 = make(map[string]string)

func main() {

    var mapChan chan map[string]string

    mapChan = make(chan map[string]string, 10)

    m1["a"] = "bbbbb"
    m1["b"] = "ccccc"

    mapChan <- m1

    var m2 = make(map[string]string)

    m2 = <-mapChan

    fmt.Println(m2)
}

// chan 的定义
