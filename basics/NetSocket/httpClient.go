package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {

    res, err := http.Get("http://www.baidu.com")
    if err != nil {
        fmt.Println(err)
        return
    }

    data, err := ioutil.ReadAll(res.Body)

    if err != nil {

        fmt.Println(err)
        return
    }

    fmt.Println(string(data))

}
