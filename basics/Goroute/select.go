package main

import (
    "fmt"
    "time"
)

func main() {

    ch := make(chan int, 10)
    ch2 := make(chan int, 10)

    go func() {
        var i int
        for {

            ch <- i
            time.Sleep(time.Second)
            ch <- i * i
            i++
        }
    }()

    for {

        select {
        case v := <-ch:
            fmt.Println(v)

        case v := <-ch2:
            fmt.Println(v)
        default:
            fmt.Println("no data")
            time.Sleep(2 * time.Second)
        }
    }

}
