package main

import (
    "fmt"
    "reflect"
)

func test(b interface{}) {
    t := reflect.TypeOf(b)

    // 如果指针
    t.Elem().SetInt(100)

    v := reflect.ValueOf(b)
    x := t.Kind()
    iv := v.Interface()

    fmt.Println(t)
    fmt.Println(v)
    fmt.Println(x)
    fmt.Println(iv)
}

func main() {

    a := 200

    test(a)
}
