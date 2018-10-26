package main

import (
    "encoding/json"
    "fmt"
)

type Link struct {
    Name string
    Next *Link
}

type Student struct {
    Name string
    Age  int
}

type ST struct {
    Name string `json:"stname"`
    Age  int    `json:"stage"`
}

func trans(p *Link) {

    for p != nil {
        fmt.Println(*p)
        p = p.Next

    }

}

func main() {

    var st1 ST = ST{
        Name: "ssssssssss",
        Age:  1,
    }
    data, err := json.Marshal(st1)
    if err != nil {

        fmt.Println("errrr", err)
        return
    }

    fmt.Println(string(data))

    var li2 *Link = &Link{

        Name: "L222222",
    }

    var li Link = Link{
        Name: "L1111",
        Next: li2,
    }

    trans(&li)

    fmt.Println(li)
    var stu Student
    stu.Age = 10

    var stu1 *Student = &Student{
        Name: "stu111111111",
    }

    fmt.Println(stu1)
    fmt.Printf("%p", &stu.Name)
    fmt.Println()
    fmt.Printf("%p", &stu.Age)

}
