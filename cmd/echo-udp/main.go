package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alcortesm/echo-udp"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	log.SetFlags(0)
	log.SetOutput(new(logger))
	log.Println("echo-udp started")

	addr, err := echoudp.Addr()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("listening: ", err)
	}
	log.Printf("listening on %v\n", addr)

	conn.SetReadBuffer(1048576) // use a 1MB buffer

	go func() {
		s := <-c
		log.Printf("signal (%v) caught, terminating...", s)
		os.Exit(0)
	}()

	var buf [1024]byte
	for {
		rlen, _, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			log.Fatal("reading: ", err)
		}
		log.Printf("received: %q\n", string(buf[:rlen]))
	}
	os.Exit(1)
}

type logger struct{}

func (_ logger) Write(b []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format(format) + " " + string(b))
}

// RFC3339 with milliseconds and right 0 padding.
const format = "2006-01-02T15:04:05.000Z07:00"
