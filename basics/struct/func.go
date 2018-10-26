package main

import (
    "fmt"
)

type ST struct {
    Name string `json:"stname"`
    Age  int    `json:"stage"`
}

func (s *ST) init(name string, age int) {

    s.Name = name
    s.Age = age

}

type MST struct {
    ST
    Sex int
}

func (s *ST) pr() {

    fmt.Println("STTTTTTTTTTTTTTTTT")
}

func main() {

    var ms MST

    ms.Name = "KKK"
    ms.Sex = 2

    var s ST = ST{
        Name: "s11111",
    }

    ms.pr()

    fmt.Println(s)

    var s2 ST

    s2.init("B", 19)

    fmt.Println(s2)

}
