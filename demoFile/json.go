package main

import (
    "encoding/json"
    "fmt"
)

type Stu struct {
    Name string
    Age  int
}

func main() {

    s1 := Stu{

        Name: "katherine",
        Age:  10,
    }

    data, _ := json.Marshal(s1)

    fmt.Println(string(data))

    js := string(data)
    var s2 Stu
    json.Unmarshal([]byte(js), &s2)

    fmt.Println(s2)
}
