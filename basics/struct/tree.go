package main

import (
    "fmt"
)

type integer int

func main() {

    var i integer = 1000
    var j int = 10

    i = integer(j)
    fmt.Println(i)
    fmt.Println(j)
}
