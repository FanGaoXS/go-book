package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var (
	addresses string
	connPool  map[string]net.Conn
)

func main() {
	addrs := parse(addresses)
	connPool = make(map[string]net.Conn, len(addrs))

	for _, addr := range addrs {
		go readFrom(addr)
	}

	defer func() {
		for addr, conn := range connPool {
			conn.Close()
			log.Printf("conn with %s has been closed", addr)
		}
	}()

	select {}
}

func init() {
	flag.StringVar(&addresses, "addrs", "", "addresses")
	flag.Parse()
}

func parse(address string) []string {
	return strings.Split(address, ",")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func readFrom(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	connPool[addr] = conn
	mustCopy(os.Stdout, conn)
}
