package server

import (
	"errors"
	"io"
	"log"
	"net"
)

func Serve(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	io.Copy(conn, conn)
}
