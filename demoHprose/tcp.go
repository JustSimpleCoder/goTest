package main

import "github.com/hprose/hprose-golang/rpc"

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	server := rpc.NewTCPServer("tcp4://0.0.0.0:1314/")
	server.AddFunction("hello", hello)
	server.Start()
}
