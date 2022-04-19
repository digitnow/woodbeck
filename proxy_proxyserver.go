package main

import (
    "sync"
	"net"
	"log"
	"io"
)

func main() {
	var wg sync.WaitGroup
	proxyServer, err := net.Listen("tcp", "127.0.0.1:8200")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bundet til %s", proxyServer.Addr().String())
    wg.Add(1)
    go func() {
		defer wg.Done()
		for {
			log.Println("før proxyServer.Accept() kallet")
			conn, err := proxyServer.Accept()
			if err != nil {
				return
			}
			go func(client net.Conn) {
				defer client.Close()

				server, err := net.Dial("tcp", "127.0.0.1:8300")
                if err != nil {
					log.Println(err)
					return
				}
				defer server.Close()
				err = proxy(client, server)
				if err != nil && err != io.EOF {
					log.Println(err)
				}
			}(conn)
		}
	}()
	wg.Wait()
}

func proxy(client io.Reader, server io.Writer) error {
	clientWriter, clientIsWriter := client.(io.Writer)
	serverReader, serverIsReader := server.(io.Reader)

	if serverIsReader && clientIsWriter {
		go func() {
			_, _ = io.Copy(clientWriter, serverReader)
		}()
	}
	_, err := io.Copy(server, client)
	return err
}