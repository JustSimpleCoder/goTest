package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

    file, err := os.Open("c:/Users/yaobin/go/test.log")
    if err != nil {

        fmt.Println(" read error")

        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    str, err := reader.ReadString('\n')

    fmt.Printf("read : %s", str)
}
