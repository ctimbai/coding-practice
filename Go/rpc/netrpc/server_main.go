package main

import (
	"coding-practice/Go/rpc/netrpc/server"
	"net"
	"net/rpc"
)

func main() {
	rpc.RegisterName("MathService", new(server.MathService))
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		logs.Fatal("listen error", e)
	}
	rpc.Accept(l)
}
