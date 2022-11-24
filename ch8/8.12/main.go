package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"go-book/ch8/8.12/broadcast"
	"log"
	"net"
)

var (
	Hostname = "localhost"
	Port     = "8090"
	Address  = fmt.Sprintf("%s:%s", Hostname, Port)
	b        *broadcast.Broadcaster
)

func main() {
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("listen on %s", Address)

	go b.Run()

	go func() {
		var input string
		for {
			fmt.Scanln(&input)
			b.Messages <- fmt.Sprintf("admin: %s", input)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	who := fmt.Sprintf("%s-%s", conn.RemoteAddr().String(), uuid.New().String())
	client := broadcast.NewClient(who)
	go func() {
		// read message from channel
		for msg := range client.Messages {
			fmt.Fprintln(conn, msg)
		}
	}()

	// when the client has been connected with server
	b.Messages <- fmt.Sprintf("%s has arrived", who)
	client.Messages <- fmt.Sprintf("%s, welcome!", who)
	b.Registers <- client

	// get message from os.Stdin
	input := bufio.NewScanner(conn)
	for input.Scan() {
		b.Messages <- fmt.Sprintf("%s: %s", who, input.Text())
	}

	// when the client has been disconnected with server
	b.Unregisters <- client
	b.Messages <- fmt.Sprintf("%s has left", who)
	conn.Close()
}

func init() {
	b = broadcast.NewBroadcaster()
}
