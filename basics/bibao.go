package main

import (
	"fmt"
)

func Addr() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func main() {

	f := Addr()

	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))

}
