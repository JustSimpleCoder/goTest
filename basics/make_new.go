package main

import (
    "fmt"
)

func test() {

    s1 := new([]int)

    fmt.Println(s1)

    s2 := make([]int, 10)
    fmt.Println(s2)

    *s1 = make([]int, 6)
    (*s1)[0] = 100
    fmt.Println(s1)

    *s1 = make([]int, 5)

    var a map[string]string

    a = make(map[string]string, 10)
    a["a"] = "aaaaaaaaaaa"

    var b map[string]string = map[string]string{
        "b": "bbbb",
    }

    c := make(map[string]int, 10)
    c["c"] = 1000
    c["cc"] = 2000

    for _, k := range c {
        fmt.Println(k)
    }

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

}

func main() {
    test()
}
