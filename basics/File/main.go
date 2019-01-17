package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

    test()
    testSscanf()
    testBufioRead()
    testReadFromFile()

}

func test() {
    fmt.Fprintf(os.Stdout, "~~zhongduan\n")

}

func testSscanf() {

    var str = "ssssssssssssssssssstring"
    var a string
    fmt.Sscanf(str, "%s", &a)

    fmt.Println(a)
}

func testBufioRead() {

    reader := bufio.NewReader(os.Stdin)

    str, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("read failed err:", err)
        return
    }

    fmt.Printf("str = %s\n", str)
}

func testReadFromFile() {

    file, err := os.Open("./a.log")
    defer file.Close()
    if err != nil {

        fmt.Println("read file failed err:", err)

    }

    reader := bufio.NewReader(file)

    str, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("read failed err:", err)
        return
    }

    fmt.Printf("str = %s\n", str)

}
