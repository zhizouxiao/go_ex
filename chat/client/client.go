package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Results struct {
	X int
}

func main() {

	client, _ := rpc.DialHTTP("tcp", "localhost:12345")

	var results Results
	client.Call("RPCMethods.Multiply", Args{1314, 2}, &results)

	fmt.Println("results: ", results)
}
