package main

import (
	"fmt"
	"net"

	"github.com/Nikita3549/httpling/internal/server"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	fmt.Println("Listening on port 8080")
	if err != nil {
		panic(err)
	}

	server.Serve(listener)
}
