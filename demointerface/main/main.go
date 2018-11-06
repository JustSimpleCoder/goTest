package main

import (
    "fmt"
    "goTest/demointerface/homework"
    "math/rand"
    "time"
)

func main() {

    // hs := make([]*homework.Hosts)
    var hs []*homework.Hosts

    for i := 0; i < 10; i++ {
        one := &homework.Hosts{
            IP:   fmt.Sprintf("10.200.%d.%d", i, rand.Intn(255)),
            PORT: "8080",
        }
        hs = append(hs, one)
    }

    n := &homework.OBO{}
    // n := &homework.R{}
    for j := 0; j < 20; j++ {

        h := n.DoSend(hs)
        fmt.Println(h)
        time.Sleep(time.Second)
    }

}
