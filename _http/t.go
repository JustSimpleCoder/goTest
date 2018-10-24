package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
)

func main() {


	// http.Handle("/y", func(){ fmt.Println("123")  })

	err := http.ListenAndServe("10.202.78.52:2918", nil)

	if err != nil {

		log.Fatal("ListenAndServe:",err)
	}

}
