package main

import "fmt"
import "encoding/json"

type P struct {
	Name     string `json:"name"`
	IsOnline bool   `json:"isOnline,string"`
}

func main() {

	p1 := P{"Y", true}

	str, err := json.MarshalIndent(p1, "", " ")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(str))

}
