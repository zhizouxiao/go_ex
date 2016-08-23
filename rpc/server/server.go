package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Results struct {
	X int
}

type RPCMethods struct {
	// just an empty struct
}

// this is an RPC method
func (m *RPCMethods) Multiply(args Args, reply *int) error {
	*reply = args.A * args.B
	fmt.Println("Multiply==========", *reply)
	return nil // no error
}

func main() {
	// expose methods of RPCMethods instance
	rpc.Register(&RPCMethods{})

	// start server
	rpc.HandleHTTP()
	http.ListenAndServe(":12345", nil)
}
