package main

import (
    "flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"woodbeck/ch09/handlers"
)

var (
	addr = flag.String("listen", "127.0.0.1:8080", "listen address")
)

func main() {
	flag.Parse()
	err := run(*addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server stoppet")
}

func run(addr string) error {
    srv := &http.Server {
		Addr: addr,
		Handler: http.TimeoutHandler(handlers.DefaultHandler(), 2*time.Minute,""),
	    IdleTimeout: 5*time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		for {
			if <-c == os.Interrupt {
				_ = srv.Close()
				return
			}
		}
	}()

	var err error
	err = srv.ListenAndServe()
	if err == http.ErrServerClosed {
		err = nil
	}
	return err
}
