package main

import (
    "fmt"
)

func d(n int) int {

    if n == 1 {
        return 1
    }

    return d(n-1) * n
}

func main() {

    c := d(5)

    fmt.Println(c)

}
