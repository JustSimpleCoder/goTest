package main

import (
    "fmt"
    "sync"
    "time"
)

var a = make(map[int]int)

var lock sync.Mutex

type task struct {
    n int
}

func calc(t *task) {

    var sum = 1
    for i := 1; i < t.n; i++ {
        sum *= i
    }
    fmt.Println(t.n, sum)
    lock.Lock()
    a[t.n] = sum
    lock.Unlock()
}

func main() {

    for i := 1; i < 11; i++ {

        t := &task{n: i}
        go calc(t)
    }
    time.Sleep(time.Second)

    lock.Lock()
    for k, v := range a {
        fmt.Printf("%d! = %d \n", k, v)
    }
    lock.Unlock()

}
