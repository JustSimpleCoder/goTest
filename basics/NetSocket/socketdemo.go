package main

import (
    "fmt"
    "net"
)

func main() {

    fmt.Println("start tcp server....")

    listen, err := net.Listen("tcp", "127.0.0.1:2918")

    if err != nil {
        fmt.Println(err)
        return
    }

    for {

        conn, err := listen.Accept()

        if err != nil {
            fmt.Println(err)
            continue
        }
        go process(conn)
    }

}

func process(conn net.Conn) {

    defer conn.Close()

    for {

        buf := make([]byte, 512)
        _, err := conn.Read(buf)

        if err != nil {
            fmt.Println(err)
            return
        }

        fmt.Println("Read:", string(buf))
    }

}
