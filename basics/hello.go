package main

import (
    "fmt"
    "time"
)

var pipe chan int

func t_go_add(a int, b int, pipe chan int) {

    sum := a + b
    pipe <- sum
}

func test_pipe(pipe chan int) (int, int) {

    t1 := <-pipe

    fmt.Println(t1)
    fmt.Println(len(pipe))

    return t1, len(pipe)
}

func main() {

    str := " hello word"

    str1 := str[0:3]

    fmt.Println(str1)

    pipe := make(chan int, 3)

    go t_go_add(1, 2, pipe)
    go t_go_add(2, 3, pipe)
    c1, p_l := test_pipe(pipe)
    go t_go_add(3, 4, pipe)

    fmt.Println(c1, p_l)

    time.Sleep(time.Second)
}
