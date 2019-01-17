package main

import (
    "context"
    "fmt"
    "time"
)

func main() {

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

    defer cancel()

    intChan := make(chan int)
    go func() {

        time.Sleep(1 * time.Second)
        intChan <- 1
    }()

    select {
    case <-ctx.Done():
        fmt.Println("time out")

    case x := <-intChan:

        fmt.Println("xxxxxxxx", x)
    }

}
