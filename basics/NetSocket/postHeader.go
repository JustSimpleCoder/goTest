package main

import (
    "fmt"
    "net"
    "net/http"
    "time"
)

var url = []string{

    "http://www.baidu.com",
    "http://www.google.com",
}

func main() {

    for _, v := range url {

        c := http.Client{
            Transport: &http.Transport{
                Dial: func(network, addr string) (net.Conn, error) {
                    return net.DialTimeout(network, addr, time.Second)
                },
            },
        }
        resheader, err := c.Head(v)
        if err != nil {

            fmt.Println(err)

            return
        }

        fmt.Println(resheader.Status)
    }
}
