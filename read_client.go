package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
	}
    defer conn.Close()

	buf := make([]byte, 1<<19) // 512 KiB
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break

		}
		fmt.Printf("read %d bytes\n", n)
	}
<<<<<<< HEAD
}
=======
	conn.Close()
}
>>>>>>> ed9f67c4ef1da77c3ca40beb79346cf40d61530e
