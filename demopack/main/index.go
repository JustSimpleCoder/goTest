package main

import (
    "fmt"
    "goTest/demopack/calc"
)

var pipe chan int

func main() {

    pipe = make(chan int, 100)

    go calc.Add(1, 2, pipe)

    go calc.Sub(1, 2, pipe)

    s1 := <-pipe
    s2 := <-pipe

    fmt.Println(s1, s2, "\r\n")

    fmt.Printf("%d \r\n", 100)
    fmt.Printf("%x \r\n", 100)
    fmt.Printf("%o \r\n", 100)

    Name := "Katherine"
    fmt.Println("main Name:", Name)
    fmt.Println("other Name:", calc.Name)

    // fmt.Println(sum, cha)
}
