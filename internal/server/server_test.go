package server_test

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"

	"github.com/Nikita3549/httpling/internal/server"
)

func TestServe_Echo(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	t.Cleanup(func() { l.Close() })

	go server.Serve(l)
	addr := l.Addr().String()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(func() { conn.Close() })
	conn.SetDeadline(time.Now().Add(2 * time.Second))

	const testData = "Test!"
	_, err = fmt.Fprint(conn, testData)
	if err != nil {
		t.Fatalf("send data: %v", err)
	}

	buf := make([]byte, len(testData))
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		t.Fatalf("read data: %v", err)
	}

	if testData != string(buf) {
		t.Fatalf("expected: %v, received: %v", testData, string(buf))
	}
}
