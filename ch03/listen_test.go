package ch03

import (
	"io"
	"net"
	"testing"
)

func TestListener(t *testing.T) {

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}
		// Anonym funksjon, kan brukes som parameter/variabel
		go func() {
			io.Copy(conn, conn)
			conn.Close()
		}()
	}

	t.Logf("Bundet til %q", listener.Addr())

}

/*

 */
