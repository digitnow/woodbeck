package main

import (
	"net"
	"crypto/rand"
	"fmt"
)

func main() {	
	// Genererer mock data
	payload := make([]byte, 1<<24) // alloker 16777216 bytes
	_, err := rand.Read(payload) // leser inn tilfeldige byte i payload
	if err != nil {
		fmt.Println(err)
	}

	listener, err := net.Listen("tcp","127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
	}

	//go func() {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		} 
		defer conn.Close()
        fmt.Println("hei, jeg venter")
		_, err = conn.Write(payload)
		if err != nil {
			fmt.Println(err)
		}
	}
	//}()

}