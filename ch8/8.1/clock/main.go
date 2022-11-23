package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var (
	port     string
	timezone string
	address  string
)

func main() {
	address = fmt.Sprintf("localhost:%s", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("listen on %s timezone is %s", port, timezone)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		content := fmt.Sprintf("time=%s timezone=%s address=%s\n", time.Now().Format("15:04:05"), timezone, address)
		_, err := io.WriteString(conn, content)
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func init() {
	flag.StringVar(&port, "p", "8000", "port number")
	flag.StringVar(&timezone, "tz", "beijing", "timezone")
	flag.Parse()
}
