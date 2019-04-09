package main

import "fmt"

func a(x int) {

	//设置recover

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var a = [10]int{}

	a[x] = 111
}

func main() {
	fmt.Println(1)
	a(2)
	fmt.Println(2)
}
