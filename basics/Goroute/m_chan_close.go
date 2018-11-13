package main

// 多个携程通过 chan 通信 通过 close 来管理结束主进程
import (
    "fmt"
)

func w(ch chan int, exitchan chan bool) {

    for i := 0; i < 10; i++ {

        ch <- i
    }

    close(ch)
    exitchan <- true

}

func r(ch chan int, exitchan chan bool) {

    for {

        v, ok := <-ch
        if !ok {
            break
        }
        fmt.Println(v)

    }

    exitchan <- true
}

func main() {

    ch := make(chan int, 10)
    exitchan := make(chan bool, 2)

    go w(ch, exitchan)
    go r(ch, exitchan)

    for i := 0; i < 2; i++ {
        _ = <-exitchan

    }

}
