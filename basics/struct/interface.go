package main

import (
    "fmt"
)

type Stu struct {
    Name string
}

type Test interface {
    Pr()
}

func (p Stu) Pr() {

    fmt.Println("aaaaaaaa")
}
func main() {

    var t Test
    var s Stu

    fmt.Println(t)
    // t.Pr()

    t = &s
    fmt.Println(t)
    t.Pr()

    //t 是指用多态
}

//String()

// "%s" 会自动查抄 String() 的接口
