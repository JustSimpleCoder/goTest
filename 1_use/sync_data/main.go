package main

import (
    "fmt"
    "goPush/K12User/service"
    // "os"
    "runtime"
)

func main() {

    fmt.Println("Start........")
    // os.Exit(1)
    runtime.GOMAXPROCS(2)
    go service.ProcessUser()
    service.WebBegin()

    //web
    //listen Db
}
