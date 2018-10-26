package main

import (
    "encoding/json"
    "fmt"
    "reflect"
)

type Stu struct {
    Name string `json:"stu_name" xml:"ffff"`
}

func testRe(x interface{}) {

    v := reflect.ValueOf(x)
    t := reflect.TypeOf(x)

    kd := v.Kind()

    if kd != reflect.Ptr {

        fmt.Println("Not a stuct")
        return
    }

    num := v.Elem().NumField()

    fmt.Println(num)

    v.Elem().Field(0).SetString("YYYYYYY")
    tag := t.Elem().Field(0).Tag.Get("xml")

    fmt.Println("taggggggggggggggg:", tag)

    for i := 0; i < num; i++ {
        fmt.Println(i, v.Elem().Field(i))
        fmt.Println(i, v.Elem().Field(i).Kind())
    }

    nf := v.Elem().NumMethod()

    fmt.Println(nf)
    var p []reflect.Value
    v.Method(1).Call(p)

}

func (s Stu) Echo() {

    fmt.Println(s)
}

func (s Stu) Echo2() {

    fmt.Println("222222", s)
}

func main() {

    var s Stu = Stu{
        Name: "ssssssssssssssssss",
    }
    data, _ := json.Marshal(s)

    fmt.Println(string(data))

    testRe(&s)
    testRe(1)

}
