package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net"
	"time"
)

var (
	Hostname = "localhost"
	Port     = "8090"
	Address  = fmt.Sprintf("%s:%s", Hostname, Port)
	b        *Broadcaster
	d        *Database
	duration time.Duration
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
	client := NewClient(who)
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
	ticker := time.NewTicker(duration)
	go func() {
		// when the client has no message over the duration, it will be exited
		<-ticker.C
		b.Messages <- fmt.Sprintf("%s has left, timeout", who)
		conn.Close()
	}()
	for input.Scan() {
		b.Messages <- fmt.Sprintf("%s: %s", who, input.Text())
		ticker.Reset(duration) // when the client has sent any message, the duration will be refreshed
	}

	// when the client has been disconnected with server
	b.Unregisters <- client
	b.Messages <- fmt.Sprintf("%s has left", who)
	conn.Close()
}

func init() {
	b = NewBroadcaster()
	d = NewDatabase()
	duration, _ = time.ParseDuration("5s")
}
