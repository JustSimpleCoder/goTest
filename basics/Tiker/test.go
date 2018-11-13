package main

import (
    "fmt"
    "time"
)

func test() {

    defer func() {

        if err := recover(); err != nil {
            fmt.Println("panic  test ", err)
        }
    }()

    var a map[int]string

    a[1] = "a"
}

func calc() {

    for {
        fmt.Println("doooooo")
        time.Sleep(time.Second)
    }

}

func main() {

    go test()

    for i := 0; i < 2; i++ {

        go calc()
    }

    time.Sleep(time.Second * 2)

}
