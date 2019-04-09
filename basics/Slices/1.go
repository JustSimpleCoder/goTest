package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3}

	fmt.Println("s1=", s1)

	s2 := make([]int, 5, 10)

	fmt.Println("s2=", s2)

}
