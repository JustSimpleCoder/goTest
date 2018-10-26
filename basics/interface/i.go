package main

import (
    "fmt"
)

type test interface {
    i()
}

type B struct {
}

func (b B) i() {
    fmt.Println("bbbb")
}

func main() {

    var a interface{}

    b := B{}
    a = b

    a.i()
}
