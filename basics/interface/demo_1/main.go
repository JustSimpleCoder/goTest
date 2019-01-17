package main

import (
    "fmt"
    "reflect"
)

type Paper interface {
    GetSize() string
}

type Book interface {
    GetName() string
    Paper
}

type Math struct {
}

func (m *Math) GetName() string {

    return "Math"
}

func (m *Math) GetSize() string {

    return "A4"
}

func main() {

    // _t1()
    _t2()

}

func _t2() {

    // 基础的接口
    var bookinterface Book
    var mathbook Math

    bookinterface = &mathbook

    name := bookinterface.GetName()
    size := bookinterface.GetSize()

    v := reflect.TypeOf(bookinterface)

    fmt.Println(v)
    fmt.Println(name, size)
}

func _t1() {
    //空接口  谁都可以实现
    var a interface{}

    var b int
    b = 1

    a = b

    fmt.Printf("%T", a)

}
