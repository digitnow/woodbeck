package main

import (
	"net"
	"log"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8200")
	if err != nil {
		log.Fatal(err)
	}
    
	log.Println("os.Args[1] = ", os.Args[1])
	_, err = conn.Write([]byte(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	} 
	response := string(buf[:n])
	log.Printf("reply from proxy: %s", response)
}