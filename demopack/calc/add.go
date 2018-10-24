package calc

func Add(a int, b int, c chan int) int {

    c <- (a + b)

    return a + b
}

func init() {

    Name = " B"
    Age = 33

}

var Name string = "K"
var Age int = 19
