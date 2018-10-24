package calc

func Sub(a int, b int, c chan int) int {

    c <- (a - b)
    return a - b
}
