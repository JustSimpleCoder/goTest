package main

import (
    "fmt"
)

func testRange() {
    var x [10]int

    for i := 0; i < 10; i++ {

        if i == 0 {
            x[i] = 1
        } else {
            x[i] = getSum(&x)
        }

    }

    for _, v := range x {
        fmt.Println(v)
    }
}

func getSum(m *[10]int) int {

    sum := 0
    for _, v := range m {
        sum += v
    }
    return sum

}

func main() {

    testRange()

}
