package main

import (
    "fmt"
)

func main() {

    intChan := make(chan int, 10)

    for i := 0; i < 10; i++ {

        intChan <- i
    }

    close(intChan)

    for {

        var b int

        b, ok := <-intChan

        if ok == false {
            fmt.Println("chan is close")
            break
        }

        fmt.Println(b)
    }

}
