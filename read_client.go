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
	conn.Close()
}
