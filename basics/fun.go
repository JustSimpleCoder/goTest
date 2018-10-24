package main

import (
    "errors"
    "fmt"
    // "time"
)

func initConfig() (err error) {

    return errors.New(" init config falied")
}

func test() {

    defer func() {

        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()

    err := initConfig()

    if err != nil {
        panic(err)
    }

    // a := 0
    // c := 100 / a
    // return c
}

func main() {

    test()
    // for {
    //     test()
    //     time.Sleep(time.Second)
    // }
}
