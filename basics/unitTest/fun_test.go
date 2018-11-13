package main

import (
    "testing"
)

func TestFunAdd(t *testing.T) {

    sum := add(2, 2)

    if sum != 4 {
        t.Fatal("add err")
    }
}
